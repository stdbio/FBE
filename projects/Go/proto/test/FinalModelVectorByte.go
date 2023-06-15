//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
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

// Fast Binary Encoding Byte vector final model
type FinalModelVectorByte struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int

    // Vector item final model
    model *fbe.FinalModelByte
}

// Create a new Byte vector final model
func NewFinalModelVectorByte(buffer *fbe.Buffer, offset int) *FinalModelVectorByte {
    fbeResult := FinalModelVectorByte{buffer: buffer, offset: offset}
    fbeResult.model = fbe.NewFinalModelByte(buffer, offset)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelVectorByte) FBEAllocationSize(values []byte) int {
    size := 4
    for _, value := range values {
        size += fm.model.FBEAllocationSize(value)
    }
    return size
}

// Get the final offset
func (fm *FinalModelVectorByte) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelVectorByte) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelVectorByte) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelVectorByte) FBEUnshift(size int) { fm.offset -= size }

// Check if the vector is valid
func (fm *FinalModelVectorByte) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + 4) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    fbeVectorSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))

    size := 4
    fm.model.SetFBEOffset(fm.FBEOffset() + 4)
    for i := fbeVectorSize; i > 0; i-- {
        offset := fm.model.Verify()
        if offset == fbe.MaxInt {
            return fbe.MaxInt
        }
        fm.model.FBEShift(offset)
        size += offset
    }
    return size
}

// Get the vector
func (fm *FinalModelVectorByte) Get() ([]byte, int, error) {
    values := make([]byte, 0)

    if (fm.buffer.Offset() + fm.FBEOffset() + 4) > fm.buffer.Size() {
        return values, 0, errors.New("model is broken")
    }

    fbeVectorSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if fbeVectorSize == 0 {
        return values, 4, nil
    }

    values = make([]byte, 0, fbeVectorSize)

    size := 4
    fm.model.SetFBEOffset(fm.FBEOffset() + 4)
    for i := 0; i < fbeVectorSize; i++ {
        value, offset, err := fm.model.Get()
        if err != nil {
            return values, size, err
        }
        values = append(values, value)
        fm.model.FBEShift(offset)
        size += offset
    }
    return values, size, nil
}

// Set the vector
func (fm *FinalModelVectorByte) Set(values []byte) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + 4) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(len(values)))

    size := 4
    fm.model.SetFBEOffset(fm.FBEOffset() + 4)
    for _, value := range values {
        offset, err := fm.model.Set(value)
        if err != nil {
            return size, err
        }
        fm.model.FBEShift(offset)
        size += offset
    }
    return size, nil
}
