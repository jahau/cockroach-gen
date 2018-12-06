// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sql

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/util"
)

// planPhysicalProps describes known ordering information for the rows generated by
// this node. The ordering information includes columns the output is ordered
// by and columns for which we know all rows have the same value. See
// physicalProps for more details.
//
// Stable after optimizePlan() (or makePlan).
// Available after newPlan(), but may change on intermediate plan
// nodes during optimizePlan() due to index selection.
func planPhysicalProps(plan planNode) physicalProps {
	switch n := plan.(type) {
	case *explainPlanNode:
		return planPhysicalProps(n.run.results)
	case *limitNode:
		return planPhysicalProps(n.plan)
	case *max1RowNode:
		return planPhysicalProps(n.plan)
	case *spoolNode:
		return planPhysicalProps(n.source)
	case *indexJoinNode:
		return n.props
	case *serializeNode:
		return planPhysicalProps(n.source)
	case *deleteNode:
		if n.run.rowsNeeded {
			return planPhysicalProps(n.source)
		}
	case *projectSetNode:
		return n.props

	case *filterNode:
		return n.props

	case *groupNode:
		return n.props

	case *windowNode:
		// TODO: window partitions can be ordered if the source is ordered
		// appropriately.
	case *joinNode:
		return n.props
	case *unionNode:
		// TODO(knz): this can be ordered if the source is ordered already.
	case *insertNode:
		// TODO(knz): RETURNING is ordered by the PK.
	case *updateNode, *upsertNode:
		// After an update, the original order may have been destroyed.
		// For example, if the PK is updated by a SET expression.
		// So we can't assume any ordering.
		//
		// TODO(knz/radu): this can be refined by an analysis which
		// determines whether the columns that participate in the ordering
		// of the source are being updated. If they are not, the source
		// ordering can be propagated.

	case *scanNode:
		return n.props
	case *ordinalityNode:
		return n.props
	case *renderNode:
		return n.props
	// TODO(richardwu): memoize the props for sort and distinct nodes.
	case *sortNode:
		return sortPhysicalProps(n)
	case *distinctNode:
		return distinctPhysicalProps(n)
	case *lookupJoinNode:
		return n.props
	case *zigzagJoinNode:
		return n.props

	// Every other node simply has no guarantees on its output rows.
	case *CreateUserNode:
	case *DropUserNode:
	case *alterIndexNode:
	case *alterSequenceNode:
	case *alterTableNode:
	case *alterUserSetPasswordNode:
	case *cancelQueriesNode:
	case *cancelSessionsNode:
	case *controlJobsNode:
	case *createDatabaseNode:
	case *createIndexNode:
	case *createSequenceNode:
	case *createStatsNode:
	case *createTableNode:
	case *createViewNode:
	case *delayedNode:
	case *dropDatabaseNode:
	case *dropIndexNode:
	case *dropSequenceNode:
	case *dropTableNode:
	case *dropViewNode:
	case *explainDistSQLNode:
	case *hookFnNode:
	case *iterativeSortStrategy:
	case *relocateNode:
	case *renameColumnNode:
	case *renameDatabaseNode:
	case *renameIndexNode:
	case *renameTableNode:
	case *rowCountNode:
	case *rowSourceToPlanNode:
	case *scatterNode:
	case *scrubNode:
	case *sequenceSelectNode:
	case *setClusterSettingNode:
	case *setVarNode:
	case *setZoneConfigNode:
	case *showFingerprintsNode:
	case *showRangesNode:
	case *showTraceNode:
	case *showTraceReplicaNode:
	case *showZoneConfigNode:
	case *sortAllStrategy:
	case *sortTopKStrategy:
	case *sortValues:
	case *splitNode:
	case *truncateNode:
	case *unaryNode:
	case *valuesNode:
	case *virtualTableNode:
	case *zeroNode:
	default:
		panic(fmt.Sprintf("unhandled node type: %T", plan))
	}

	return physicalProps{}
}

func sortPhysicalProps(n *sortNode) physicalProps {
	underlying := planPhysicalProps(n.plan)
	props := underlying.copy()

	// If we aren't sorting, the underlying plan's ordering can be more specific
	// than the sortNode's ordering, so we want to use that. E.g:
	//   CREATE INDEX foo ON t (a, b);
	//   SELECT a, b, c FROM t@foo ORDER BY a;
	// We want to use (a, b) instead of just (a) (we return below).
	// If we do need sorting, we must use the sortNode's ordering.
	if n.needSort {
		// We will sort and can guarantee the desired ordering.
		props.ordering = make(sqlbase.ColumnOrdering, 0, len(n.ordering))
		for _, o := range n.ordering {
			props.addOrderColumn(o.ColIdx, o.Direction)
		}
	}

	if numPlanColumns := len(planColumns(n.plan)); len(n.columns) < numPlanColumns {
		// The sortNode is projecting away columns, e.g:
		//   SELECT k FROM kv ORDER BY v.
		colMap := make([]int, numPlanColumns)
		for i := range colMap {
			if i < len(n.columns) {
				colMap[i] = i
			} else {
				colMap[i] = -1
			}
		}
		return props.project(colMap)
	}
	return props
}

func distinctPhysicalProps(n *distinctNode) physicalProps {
	underlying := planPhysicalProps(n.plan)
	props := underlying.copy()

	if !n.distinctOnColIdxs.Empty() {
		props.addWeakKey(n.distinctOnColIdxs)
	} else {
		// No specific DISTINCT ON columns. All child plan columns
		// form a key after distinct.
		var allCols util.FastIntSet
		allCols.AddRange(0, len(planMutableColumns(n.plan))-1)
		props.addWeakKey(allCols)
	}

	return props
}
