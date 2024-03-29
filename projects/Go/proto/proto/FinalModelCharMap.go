//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package proto

import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding CharMap final model
type FinalModelCharMap struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    Abbr *FinalModelMapCharString
}

// Create a new CharMap final model
func NewFinalModelCharMap(buffer *fbe.Buffer, offset int) *FinalModelCharMap {
    fbeResult := FinalModelCharMap{buffer: buffer, offset: offset}
    fbeResult.Abbr = NewFinalModelMapCharString(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelCharMap) FBEAllocationSize(fbeValue *CharMap) int {
    fbeResult := 0 +
        fm.Abbr.FBEAllocationSize(fbeValue.Abbr) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelCharMap) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelCharMap) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelCharMap) FBEType() int { return 1 }

// Get the final offset
func (fm *FinalModelCharMap) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelCharMap) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelCharMap) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelCharMap) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelCharMap) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelCharMap) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.Abbr.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.Abbr.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelCharMap) Get() (*CharMap, int, error) {
    fbeResult := NewCharMap()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelCharMap) GetValue(fbeValue *CharMap) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelCharMap) GetFields(fbeValue *CharMap) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.Abbr.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.Abbr, fbeFieldSize, err = fm.Abbr.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

// Set the struct value
func (fm *FinalModelCharMap) Set(fbeValue *CharMap) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelCharMap) SetFields(fbeValue *CharMap) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.Abbr.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.Abbr.Set(fbeValue.Abbr); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}
