//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructBytes final model
type FinalModelStructBytes struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    F1 *fbe.FinalModelBytes
    F2 *FinalModelOptionalBytes
    F3 *FinalModelOptionalBytes
}

// Create a new StructBytes final model
func NewFinalModelStructBytes(buffer *fbe.Buffer, offset int) *FinalModelStructBytes {
    fbeResult := FinalModelStructBytes{buffer: buffer, offset: offset}
    fbeResult.F1 = fbe.NewFinalModelBytes(buffer, 0)
    fbeResult.F2 = NewFinalModelOptionalBytes(buffer, 0)
    fbeResult.F3 = NewFinalModelOptionalBytes(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelStructBytes) FBEAllocationSize(fbeValue *StructBytes) int {
    fbeResult := 0 +
        fm.F1.FBEAllocationSize(fbeValue.F1) +
        fm.F2.FBEAllocationSize(fbeValue.F2) +
        fm.F3.FBEAllocationSize(fbeValue.F3) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelStructBytes) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelStructBytes) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelStructBytes) FBEType() int { return 120 }

// Get the final offset
func (fm *FinalModelStructBytes) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelStructBytes) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelStructBytes) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelStructBytes) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelStructBytes) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelStructBytes) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F2.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F3.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelStructBytes) Get() (*StructBytes, int, error) {
    fbeResult := NewStructBytes()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelStructBytes) GetValue(fbeValue *StructBytes) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelStructBytes) GetFields(fbeValue *StructBytes) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1, fbeFieldSize, err = fm.F1.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F2, fbeFieldSize, err = fm.F2.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F3, fbeFieldSize, err = fm.F3.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

// Set the struct value
func (fm *FinalModelStructBytes) Set(fbeValue *StructBytes) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelStructBytes) SetFields(fbeValue *StructBytes) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1.Set(fbeValue.F1); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F2.Set(fbeValue.F2); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F3.Set(fbeValue.F3); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}
