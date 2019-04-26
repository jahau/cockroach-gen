// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package exec

import (
	"context"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
	"github.com/pkg/errors"
)

// NewConstOp creates a new operator that produces a constant value constVal of
// type t at index outputIdx.
func NewConstOp(input Operator, t types.T, constVal interface{}, outputIdx int) (Operator, error) {
	switch t {
	case types.Bool:
		return &constBoolOp{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(bool),
		}, nil
	case types.Bytes:
		return &constBytesOp{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.([]byte),
		}, nil
	case types.Decimal:
		return &constDecimalOp{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(apd.Decimal),
		}, nil
	case types.Int8:
		return &constInt8Op{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(int8),
		}, nil
	case types.Int16:
		return &constInt16Op{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(int16),
		}, nil
	case types.Int32:
		return &constInt32Op{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(int32),
		}, nil
	case types.Int64:
		return &constInt64Op{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(int64),
		}, nil
	case types.Float32:
		return &constFloat32Op{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(float32),
		}, nil
	case types.Float64:
		return &constFloat64Op{
			input:     input,
			outputIdx: outputIdx,
			typ:       t,
			constVal:  constVal.(float64),
		}, nil
	default:
		return nil, errors.Errorf("unsupported const type %s", t)
	}
}

type constBoolOp struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  bool
}

func (c constBoolOp) Init() {
	c.input.Init()
}

func (c constBoolOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Bool()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constBytesOp struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  []byte
}

func (c constBytesOp) Init() {
	c.input.Init()
}

func (c constBytesOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Bytes()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constDecimalOp struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  apd.Decimal
}

func (c constDecimalOp) Init() {
	c.input.Init()
}

func (c constDecimalOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Decimal()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt8Op struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  int8
}

func (c constInt8Op) Init() {
	c.input.Init()
}

func (c constInt8Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Int8()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt16Op struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  int16
}

func (c constInt16Op) Init() {
	c.input.Init()
}

func (c constInt16Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Int16()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt32Op struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  int32
}

func (c constInt32Op) Init() {
	c.input.Init()
}

func (c constInt32Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Int32()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constInt64Op struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  int64
}

func (c constInt64Op) Init() {
	c.input.Init()
}

func (c constInt64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Int64()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constFloat32Op struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  float32
}

func (c constFloat32Op) Init() {
	c.input.Init()
}

func (c constFloat32Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Float32()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

type constFloat64Op struct {
	input Operator

	typ       types.T
	outputIdx int
	constVal  float64
}

func (c constFloat64Op) Init() {
	c.input.Init()
}

func (c constFloat64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(c.typ)
		col := batch.ColVec(c.outputIdx).Float64()
		for i := range col {
			col[i] = c.constVal
		}
	}
	return batch
}

// NewConstNullOp creates a new operator that produces a constant (untyped) NULL
// value at index outputIdx.
func NewConstNullOp(input Operator, outputIdx int) Operator {
	return &constNullOp{
		input:     input,
		outputIdx: outputIdx,
	}
}

type constNullOp struct {
	input     Operator
	outputIdx int
}

func (c constNullOp) Init() {
	c.input.Init()
}

func (c constNullOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	if batch.Length() == 0 {
		return batch
	}

	if batch.Width() == c.outputIdx {
		batch.AppendCol(types.Int8)
		col := batch.ColVec(c.outputIdx)
		col.Nulls().SetNulls()
	}
	return batch
}
