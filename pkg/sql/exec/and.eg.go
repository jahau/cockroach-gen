// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package exec

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
)

type andOp struct {
	OneInputNode

	leftIdx   int
	rightIdx  int
	outputIdx int
}

// NewAndOp returns a new operator that logical-ANDs the boolean columns at
// leftIdx and rightIdx, returning the result in outputIdx.
func NewAndOp(input Operator, leftIdx, rightIdx, outputIdx int) Operator {
	return &andOp{
		OneInputNode: NewOneInputNode(input),
		leftIdx:      leftIdx,
		rightIdx:     rightIdx,
		outputIdx:    outputIdx,
	}
}

func (a *andOp) Init() {
	a.input.Init()
}

func (a *andOp) Next(ctx context.Context) coldata.Batch {
	batch := a.input.Next(ctx)
	if a.outputIdx == batch.Width() {
		batch.AppendCol(coltypes.Bool)
	}
	n := batch.Length()
	if n == 0 {
		return batch
	}
	leftCol := batch.ColVec(a.leftIdx)
	rightCol := batch.ColVec(a.rightIdx)
	outputCol := batch.ColVec(a.outputIdx)

	leftColVals := leftCol.Bool()
	rightColVals := rightCol.Bool()
	outputColVals := outputCol.Bool()
	outputNulls := outputCol.Nulls()
	if leftCol.MaybeHasNulls() {
		leftNulls := leftCol.Nulls()
		if rightCol.MaybeHasNulls() {
			rightNulls := rightCol.Nulls()
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					idx := i
					isLeftNull := leftNulls.NullAt(idx)
					isRightNull := rightNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			} else {
				_ = rightColVals[n-1]
				_ = outputColVals[n-1]
				for i := range leftColVals[:n] {
					idx := uint16(i)
					isLeftNull := leftNulls.NullAt(idx)
					isRightNull := rightNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					idx := i
					isLeftNull := leftNulls.NullAt(idx)
					isRightNull := false
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			} else {
				_ = rightColVals[n-1]
				_ = outputColVals[n-1]
				for i := range leftColVals[:n] {
					idx := uint16(i)
					isLeftNull := leftNulls.NullAt(idx)
					isRightNull := false
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			}
		}
	} else {
		if rightCol.MaybeHasNulls() {
			rightNulls := rightCol.Nulls()
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					idx := i
					isLeftNull := false
					isRightNull := rightNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			} else {
				_ = rightColVals[n-1]
				_ = outputColVals[n-1]
				for i := range leftColVals[:n] {
					idx := uint16(i)
					isLeftNull := false
					isRightNull := rightNulls.NullAt(idx)
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					idx := i
					isLeftNull := false
					isRightNull := false
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			} else {
				_ = rightColVals[n-1]
				_ = outputColVals[n-1]
				for i := range leftColVals[:n] {
					idx := uint16(i)
					isLeftNull := false
					isRightNull := false
					leftVal := leftColVals[idx]
					rightVal := rightColVals[idx]
					if (!leftVal && !isLeftNull) || (!rightVal && !isRightNull) {
						// Rule 1: at least one boolean is FALSE.
						outputColVals[idx] = false
					} else if (leftVal && !isLeftNull) && (rightVal && !isRightNull) {
						// Rule 2: both booleans are TRUE.
						outputColVals[idx] = true
					} else {
						// Rule 3.
						outputNulls.SetNull(idx)
					}
				}
			}
		}
	}

	return batch
}
