//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: ptrpkg.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package ptrpkg

import "errors"
import "fbeproj/proto/fbe"
import "fbeproj/proto/variants_ptr"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = variants_ptr.Version

// Fast Binary Encoding Line field model
type FieldModelLine struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    Value *variants_ptr.FieldModelValue
    Value_ptr *FieldModelPtrvariants_ptrValue
}

// Create a new Line field model
func NewFieldModelLine(buffer *fbe.Buffer, offset int) *FieldModelLine {
    fbeResult := FieldModelLine{buffer: buffer, offset: offset}
    fbeResult.Value = variants_ptr.NewFieldModelValue(buffer, 4 + 4)
    fbeResult.Value_ptr = NewFieldModelPtrvariants_ptrValue(buffer, fbeResult.Value.FBEOffset() + fbeResult.Value.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelLine) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelLine) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.Value.FBESize() +
        fm.Value_ptr.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelLine) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.Value.FBEExtra() +
        fm.Value_ptr.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelLine) FBEType() int { return 1 }

// Get the field offset
func (fm *FieldModelLine) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelLine) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelLine) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelLine) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelLine) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelLine) VerifyType(fbeVerifyType bool) bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return false
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return false
    }

    fbeStructType := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4))
    if fbeVerifyType && (fbeStructType != fm.FBEType()) {
        return false
    }

    fm.buffer.Shift(fbeStructOffset)
    fbeResult := fm.VerifyFields(fbeStructSize)
    fm.buffer.Unshift(fbeStructOffset)
    return fbeResult
}

// // Check if the struct value fields are valid
func (fm *FieldModelLine) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Value.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Value.Verify() {
        return false
    }
    fbeCurrentSize += fm.Value.FBESize()

    if (fbeCurrentSize + fm.Value_ptr.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Value_ptr.Verify() {
        return false
    }
    fbeCurrentSize += fm.Value_ptr.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelLine) GetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, nil
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return 0, errors.New("model is broken")
    }

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Get the struct value (end phase)
func (fm *FieldModelLine) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelLine) Get() (*Line, error) {
    fbeResult := NewLine()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelLine) GetValue(fbeValue *Line) error {
    fbeBegin, err := fm.GetBegin()
    if fbeBegin == 0 {
        return err
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset()))
    fm.GetFields(fbeValue, fbeStructSize)
    fm.GetEnd(fbeBegin)
    return nil
}

// Get the struct fields values
func (fm *FieldModelLine) GetFields(fbeValue *Line, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Value.FBESize()) <= fbeStructSize {
        _ = fm.Value.GetValue(&fbeValue.Value)
    } else {
        fbeValue.Value = *variants_ptr.NewValue()
    }
    fbeCurrentSize += fm.Value.FBESize()

    if (fbeCurrentSize + fm.Value_ptr.FBESize()) <= fbeStructSize {
        fbeValue.Value_ptr, _ = fm.Value_ptr.Get()
    } else {
        fbeValue.Value_ptr = nil
    }
    fbeCurrentSize += fm.Value_ptr.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelLine) SetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := fm.FBEBody()
    fbeStructOffset := fm.buffer.Allocate(fbeStructSize) - fm.buffer.Offset()
    if (fbeStructOffset <= 0) || ((fm.buffer.Offset() + fbeStructOffset + fbeStructSize) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeStructOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset, uint32(fbeStructSize))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4, uint32(fm.FBEType()))

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Set the struct value (end phase)
func (fm *FieldModelLine) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelLine) Set(fbeValue *Line) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelLine) SetFields(fbeValue *Line) error {
    var err error = nil

    if err = fm.Value.Set(&fbeValue.Value); err != nil {
        return err
    }
    if err = fm.Value_ptr.Set(fbeValue.Value_ptr); err != nil {
        return err
    }
    return err
}
