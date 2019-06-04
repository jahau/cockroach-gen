// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License included
// in the file licenses/BSL.txt and at www.mariadb.com/bsl11.
//
// Change Date: 2022-10-01
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt and at
// https://www.apache.org/licenses/LICENSE-2.0

package sqlsmith

import (
	gosql "database/sql"
	"fmt"
	"math/rand"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/util/syncutil"
)

// sqlsmith-go
//
// sqlsmith-go is a random SQL query generator, based off of sqlsmith:
//
//   https://github.com/anse1/sqlsmith
//
// You can think of it as walking a randomly generated AST and materializing
// that AST as it goes, which it then feeds into Cockroach with the hopes of
// finding panics.
//
// However, naively generating such an AST will only find certain kinds of
// panics: they're almost guaranteed not to pass semantic analysis, and so
// any components of the system beyond that will probably not be tested.
// To get around this, sqlsmith tracks scopes and types, very similar to
// how the optbuilder works, to create ASTs which will likely pass
// semantic analysis.
//
// It does this by building the tree top-down. Every level of the tree
// requests input of a certain form. For instance, a SELECT will request
// a list of projections which respect the scope that the SELECT introduces,
// and a function call will request an input value of a particular type,
// subject to the same scope it has. This raises a question: what if we
// are unable to construct an expression meeting the restrictions requested
// by the parent expression? Rather than do some fancy constraint solving
// (which could be an interesting direction for this tool to go in the
// future, but I've found to be difficult when I've tried in the past)
// sqlsmith will simply try randomly to generate an expression, and once
// it fails a certain number of times, it will retreat up the tree and
// retry at a higher level.

const retryCount = 20

// Smither is a sqlsmith generator.
type Smither struct {
	rnd            *rand.Rand
	db             *gosql.DB
	lock           syncutil.RWMutex
	tables         []*tableRef
	indexes        map[tree.TableName]map[tree.Name]*tree.CreateIndex
	nameCounts     map[string]int
	alters         *WeightedSampler
	scalars, bools *WeightedSampler
	selectStmts    *WeightedSampler

	stmtSampler, tableExprSampler *WeightedSampler
	statements                    statementWeights
	tableExprs                    tableExprWeights

	disableWith bool
}

// NewSmither creates a new Smither. db is used to populate existing tables
// for use as column references. It can be nil to skip table population.
func NewSmither(db *gosql.DB, rnd *rand.Rand, opts ...SmitherOption) (*Smither, error) {
	s := &Smither{
		rnd:         rnd,
		db:          db,
		nameCounts:  map[string]int{},
		scalars:     NewWeightedSampler(scalarWeights, rnd.Int63()),
		bools:       NewWeightedSampler(boolWeights, rnd.Int63()),
		selectStmts: NewWeightedSampler(selectStmtWeights, rnd.Int63()),
		alters:      NewWeightedSampler(alterWeights, rnd.Int63()),

		statements: allStatements,
		tableExprs: allTableExprs,
	}
	for _, opt := range opts {
		opt.Apply(s)
	}
	s.stmtSampler = NewWeightedSampler(s.statements.Weights(), rnd.Int63())
	s.tableExprSampler = NewWeightedSampler(s.tableExprs.Weights(), rnd.Int63())
	return s, s.ReloadSchemas()
}

var prettyCfg = func() tree.PrettyCfg {
	cfg := tree.DefaultPrettyCfg()
	cfg.LineWidth = 120
	cfg.Simplify = false
	return cfg
}()

// Generate returns a random SQL string.
func (s *Smither) Generate() string {
	i := 0
	for {
		scope := s.makeScope()
		stmt, ok := scope.makeStmt()
		if !ok {
			i++
			if i > 1000 {
				panic("exhausted generation attempts")
			}
			continue
		}
		i = 0
		return prettyCfg.Pretty(stmt)
	}
}

func (s *Smither) name(prefix string) tree.Name {
	s.lock.Lock()
	s.nameCounts[prefix]++
	count := s.nameCounts[prefix]
	s.lock.Unlock()
	return tree.Name(fmt.Sprintf("%s_%d", prefix, count))
}

// SmitherOption is an option for the Smither client.
type SmitherOption interface {
	Apply(*Smither)
}

// DisableMutations causes the Smither to not emit statements that could
// mutate any on-disk data.
func DisableMutations() SmitherOption {
	return disableMutations{}
}

type disableMutations struct{}

func (d disableMutations) Apply(s *Smither) {
	s.statements = nonMutatingStatements
	s.tableExprs = nonMutatingTableExprs
}

// DisableWith causes the Smither to not emit WITH clauses.
func DisableWith() SmitherOption {
	return disableWith{}
}

type disableWith struct{}

func (d disableWith) Apply(s *Smither) {
	s.disableWith = true
}
