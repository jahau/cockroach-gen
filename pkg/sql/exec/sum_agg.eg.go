// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License included
// in the file licenses/BSL.txt and at www.mariadb.com/bsl11.
// Change Date: 2022-10-01
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt and at
// https://www.apache.org/licenses/LICENSE-2.0

package exec

import (
	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/pkg/errors"
)

func newSumAgg(t types.T) (aggregateFunc, error) {
	switch t {
	case types.Decimal:
		return &sumDecimalAgg{}, nil
	case types.Int8:
		return &sumInt8Agg{}, nil
	case types.Int16:
		return &sumInt16Agg{}, nil
	case types.Int32:
		return &sumInt32Agg{}, nil
	case types.Int64:
		return &sumInt64Agg{}, nil
	case types.Float32:
		return &sumFloat32Agg{}, nil
	case types.Float64:
		return &sumFloat64Agg{}, nil
	default:
		return nil, errors.Errorf("unsupported sum agg type %s", t)
	}
}

type sumDecimalAgg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []apd.Decimal
	}
}

var _ aggregateFunc = &sumDecimalAgg{}

func (a *sumDecimalAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.Reset()
}

func (a *sumDecimalAgg) Reset() {
	copy(a.scratch.vec, zeroDecimalColumn)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumDecimalAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumDecimalAgg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroDecimalColumn)
	}
}

func (a *sumDecimalAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Decimal(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			if _, err := tree.DecimalCtx.Add(&a.scratch.vec[a.scratch.curIdx], &a.scratch.vec[a.scratch.curIdx], &col[i]); err != nil {
				panic(err)
			}
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			if _, err := tree.DecimalCtx.Add(&a.scratch.vec[a.scratch.curIdx], &a.scratch.vec[a.scratch.curIdx], &col[i]); err != nil {
				panic(err)
			}
		}
	}
}

type sumInt8Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []int8
	}
}

var _ aggregateFunc = &sumInt8Agg{}

func (a *sumInt8Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int8()
	a.Reset()
}

func (a *sumInt8Agg) Reset() {
	copy(a.scratch.vec, zeroInt8Column)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumInt8Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt8Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroInt8Column)
	}
}

func (a *sumInt8Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Int8(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	}
}

type sumInt16Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []int16
	}
}

var _ aggregateFunc = &sumInt16Agg{}

func (a *sumInt16Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int16()
	a.Reset()
}

func (a *sumInt16Agg) Reset() {
	copy(a.scratch.vec, zeroInt16Column)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumInt16Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt16Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroInt16Column)
	}
}

func (a *sumInt16Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Int16(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	}
}

type sumInt32Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []int32
	}
}

var _ aggregateFunc = &sumInt32Agg{}

func (a *sumInt32Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int32()
	a.Reset()
}

func (a *sumInt32Agg) Reset() {
	copy(a.scratch.vec, zeroInt32Column)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumInt32Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt32Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroInt32Column)
	}
}

func (a *sumInt32Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Int32(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	}
}

type sumInt64Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []int64
	}
}

var _ aggregateFunc = &sumInt64Agg{}

func (a *sumInt64Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Int64()
	a.Reset()
}

func (a *sumInt64Agg) Reset() {
	copy(a.scratch.vec, zeroInt64Column)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumInt64Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt64Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroInt64Column)
	}
}

func (a *sumInt64Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Int64(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	}
}

type sumFloat32Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []float32
	}
}

var _ aggregateFunc = &sumFloat32Agg{}

func (a *sumFloat32Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Float32()
	a.Reset()
}

func (a *sumFloat32Agg) Reset() {
	copy(a.scratch.vec, zeroFloat32Column)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumFloat32Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumFloat32Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroFloat32Column)
	}
}

func (a *sumFloat32Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Float32(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	}
}

type sumFloat64Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// vec points to the output vector we are updating.
		vec []float64
	}
}

var _ aggregateFunc = &sumFloat64Agg{}

func (a *sumFloat64Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Float64()
	a.Reset()
}

func (a *sumFloat64Agg) Reset() {
	copy(a.scratch.vec, zeroFloat64Column)
	a.scratch.curIdx = -1
	a.done = false
}

func (a *sumFloat64Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumFloat64Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.vec[idx+1:], zeroFloat64Column)
	}
}

func (a *sumFloat64Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Float64(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.vec[a.scratch.curIdx] = a.scratch.vec[a.scratch.curIdx] + col[i]
		}
	}
}
