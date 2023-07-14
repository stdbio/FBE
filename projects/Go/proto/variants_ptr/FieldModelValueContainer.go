//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants_ptr.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package variants_ptr

import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding ValueContainer field model
type FieldModelValueContainer struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    Vv *FieldModelVectorV
    Vm *FieldModelMapInt32V
}

// Create a new ValueContainer field model
func NewFieldModelValueContainer(buffer *fbe.Buffer, offset int) *FieldModelValueContainer {
    fbeResult := FieldModelValueContainer{buffer: buffer, offset: offset}
    fbeResult.Vv = NewFieldModelVectorV(buffer, 4 + 4)
    fbeResult.Vm = NewFieldModelMapInt32V(buffer, fbeResult.Vv.FBEOffset() + fbeResult.Vv.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelValueContainer) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelValueContainer) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.Vv.FBESize() +
        fm.Vm.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelValueContainer) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.Vv.FBEExtra() +
        fm.Vm.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelValueContainer) FBEType() int { return 3 }

// Get the field offset
func (fm *FieldModelValueContainer) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelValueContainer) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelValueContainer) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelValueContainer) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelValueContainer) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelValueContainer) VerifyType(fbeVerifyType bool) bool {
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
func (fm *FieldModelValueContainer) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Vv.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Vv.Verify() {
        return false
    }
    fbeCurrentSize += fm.Vv.FBESize()

    if (fbeCurrentSize + fm.Vm.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Vm.Verify() {
        return false
    }
    fbeCurrentSize += fm.Vm.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelValueContainer) GetBegin() (int, error) {
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
func (fm *FieldModelValueContainer) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelValueContainer) Get() (*ValueContainer, error) {
    fbeResult := NewValueContainer()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelValueContainer) GetValue(fbeValue *ValueContainer) error {
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
func (fm *FieldModelValueContainer) GetFields(fbeValue *ValueContainer, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Vv.FBESize()) <= fbeStructSize {
        fbeValue.Vv, _ = fm.Vv.Get()
    } else {
        fbeValue.Vv = make([]V, 0)
    }
    fbeCurrentSize += fm.Vv.FBESize()

    if (fbeCurrentSize + fm.Vm.FBESize()) <= fbeStructSize {
        fbeValue.Vm, _ = fm.Vm.Get()
    } else {
        fbeValue.Vm = make(map[int32]V)
    }
    fbeCurrentSize += fm.Vm.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelValueContainer) SetBegin() (int, error) {
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
func (fm *FieldModelValueContainer) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelValueContainer) Set(fbeValue *ValueContainer) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelValueContainer) SetFields(fbeValue *ValueContainer) error {
    var err error = nil

    if err = fm.Vv.Set(fbeValue.Vv); err != nil {
        return err
    }
    if err = fm.Vm.Set(fbeValue.Vm); err != nil {
        return err
    }
    return err
}
