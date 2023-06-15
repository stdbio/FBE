//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package variants

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding Value final model
type FinalModelValue struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    V *FinalModelV
}

// Create a new Value final model
func NewFinalModelValue(buffer *fbe.Buffer, offset int) *FinalModelValue {
    fbeResult := FinalModelValue{buffer: buffer, offset: offset}
    fbeResult.V = NewFinalModelV(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelValue) FBEAllocationSize(fbeValue *Value) int {
    fbeResult := 0 +
        fm.V.FBEAllocationSize(&fbeValue.V) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelValue) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelValue) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelValue) FBEType() int { return 2 }

// Get the final offset
func (fm *FinalModelValue) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelValue) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelValue) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelValue) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelValue) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelValue) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.V.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.V.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelValue) Get() (*Value, int, error) {
    fbeResult := NewValue()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelValue) GetValue(fbeValue *Value) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelValue) GetFields(fbeValue *Value) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.V.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.V.GetValue(&fbeValue.V); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

// Set the struct value
func (fm *FinalModelValue) Set(fbeValue *Value) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelValue) SetFields(fbeValue *Value) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.V.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.V.Set(&fbeValue.V); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}
