// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package coldata

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
)

func (m *memColumn) Append(args AppendArgs) {
	switch args.ColType {
	case types.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]bool, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([][]byte, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]apd.Decimal, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Int8:
		fromCol := args.Src.Int8()
		toCol := m.Int8()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]int8, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]int16, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]int32, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]int64, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Float32:
		fromCol := args.Src.Float32()
		toCol := m.Float32()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]float32, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	case types.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		numToAppend := args.SrcEndIdx - args.SrcStartIdx
		if args.Sel == nil {
			toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			m.nulls.Extend(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend)
		} else {
			sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
			appendVals := make([]float64, len(sel))
			for i, selIdx := range sel {
				appendVals[i] = fromCol[selIdx]
			}
			toCol = append(toCol[:args.DestIdx], appendVals...)
			// TODO(asubiotto): Change Extend* signatures to allow callers to pass in
			// SrcEndIdx instead of numToAppend.
			m.nulls.ExtendWithSel(args.Src.Nulls(), args.DestIdx, args.SrcStartIdx, numToAppend, args.Sel)
		}
		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Copy(args CopyArgs) {
	if args.Nils != nil && args.Sel64 == nil {
		panic("Nils set without Sel64")
	}

	m.Nulls().UnsetNullRange(args.DestIdx, args.DestIdx+(args.SrcEndIdx-args.SrcStartIdx))

	switch args.ColType {
	case types.Bool:
		fromCol := args.Src.Bool()
		toCol := m.Bool()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Bytes:
		fromCol := args.Src.Bytes()
		toCol := m.Bytes()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Decimal:
		fromCol := args.Src.Decimal()
		toCol := m.Decimal()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Int8:
		fromCol := args.Src.Int8()
		toCol := m.Int8()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Int16:
		fromCol := args.Src.Int16()
		toCol := m.Int16()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Int32:
		fromCol := args.Src.Int32()
		toCol := m.Int32()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Int64:
		fromCol := args.Src.Int64()
		toCol := m.Int64()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Float32:
		fromCol := args.Src.Float32()
		toCol := m.Float32()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	case types.Float64:
		fromCol := args.Src.Float64()
		toCol := m.Float64()
		if args.Sel64 != nil {
			sel := args.Sel64
			// TODO(asubiotto): Template this and the uint16 case below.
			if args.Nils != nil {
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					toColSliced := toCol[args.DestIdx:]
					for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
						if args.Nils[i] || nulls.NullAt64(selIdx) {
							m.nulls.SetNull64(uint64(i) + args.DestIdx)
						} else {
							toColSliced[i] = fromCol[selIdx]
						}
					}
					return
				}
				// Nils but no Nulls.
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if args.Nils[i] {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils.
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(selIdx) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nils or Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		} else if args.Sel != nil {
			sel := args.Sel
			if args.Src.MaybeHasNulls() {
				nulls := args.Src.Nulls()
				toColSliced := toCol[args.DestIdx:]
				for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
					if nulls.NullAt64(uint64(selIdx)) {
						m.nulls.SetNull64(uint64(i) + args.DestIdx)
					} else {
						toColSliced[i] = fromCol[selIdx]
					}
				}
				return
			}
			// No Nulls.
			toColSliced := toCol[args.DestIdx:]
			for i, selIdx := range sel[args.SrcStartIdx:args.SrcEndIdx] {
				toColSliced[i] = fromCol[selIdx]
			}
			return
		}
		// No Sel or Sel64.
		copy(toCol[args.DestIdx:], fromCol[args.SrcStartIdx:args.SrcEndIdx])
		if args.Src.MaybeHasNulls() {
			// TODO(asubiotto): This should use Extend but Extend only takes uint16
			// arguments.
			srcNulls := args.Src.Nulls()
			for curDestIdx, curSrcIdx := args.DestIdx, args.SrcStartIdx; curSrcIdx < args.SrcEndIdx; curDestIdx, curSrcIdx = curDestIdx+1, curSrcIdx+1 {
				if srcNulls.NullAt64(curSrcIdx) {
					m.nulls.SetNull64(curDestIdx)
				}
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", args.ColType))
	}
}

func (m *memColumn) Slice(colType types.T, start uint64, end uint64) Vec {
	switch colType {
	case types.Bool:
		col := m.Bool()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Bytes:
		col := m.Bytes()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Decimal:
		col := m.Decimal()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Int8:
		col := m.Int8()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Int16:
		col := m.Int16()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Int32:
		col := m.Int32()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Int64:
		col := m.Int64()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Float32:
		col := m.Float32()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	case types.Float64:
		col := m.Float64()
		return &memColumn{
			col:   col[start:end],
			nulls: m.nulls.Slice(start, end),
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) PrettyValueAt(colIdx uint16, colType types.T) string {
	if m.nulls.NullAt(colIdx) {
		return "NULL"
	}
	switch colType {
	case types.Bool:
		return fmt.Sprintf("%v", m.Bool()[colIdx])
	case types.Bytes:
		return fmt.Sprintf("%v", m.Bytes()[colIdx])
	case types.Decimal:
		return fmt.Sprintf("%v", m.Decimal()[colIdx])
	case types.Int8:
		return fmt.Sprintf("%v", m.Int8()[colIdx])
	case types.Int16:
		return fmt.Sprintf("%v", m.Int16()[colIdx])
	case types.Int32:
		return fmt.Sprintf("%v", m.Int32()[colIdx])
	case types.Int64:
		return fmt.Sprintf("%v", m.Int64()[colIdx])
	case types.Float32:
		return fmt.Sprintf("%v", m.Float32()[colIdx])
	case types.Float64:
		return fmt.Sprintf("%v", m.Float64()[colIdx])
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}
