//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants_ptr.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package variants_ptr

import "fmt"
import "strconv"
import "strings"
import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// Value key
type ValueKey struct {
}

// Convert Value flags key to string
func (k *ValueKey) String() string {
    var sb strings.Builder
    sb.WriteString("ValueKey(")
    sb.WriteString(")")
    return sb.String()
}

// Value struct
type Value struct {
    V V `json:"v"`
    Vo *V `json:"vo"`
    Vo2 *V `json:"vo2"`
}

// Create a new Value struct
func NewValue() *Value {
    return &Value{
        V: *NewV(),
        Vo: nil,
        Vo2: nil,
    }
}

// Create a new Value struct from the given field values
func NewValueFromFieldValues(vV V, voV *V, vo2V *V) *Value {
    return &Value{vV, voV, vo2V}
}

// Create a new Value struct from JSON
func NewValueFromJSON(buffer []byte) (*Value, error) {
    result := *NewValue()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *Value) Copy() *Value {
    var result = *s
    return &result
}

// Struct deep clone
func (s *Value) Clone() *Value {
    // Serialize the struct to the FBE stream
    writer := NewValueModel(fbe.NewEmptyBuffer())
    _, _ = writer.Serialize(s)

    // Deserialize the struct from the FBE stream
    reader := NewValueModel(writer.Buffer())
    result, _, _ := reader.Deserialize()
    return result
}

// Get the struct key
func (s *Value) Key() ValueKey {
    return ValueKey{
    }
}

// Convert struct to optional
func (s *Value) Optional() *Value {
    return s
}

// Convert struct to optional
func (s *Value) Ptr() *Value {
    return s
}

// Get the FBE type
func (s *Value) FBEType() int { return 2 }

// Convert struct to string
func (s *Value) String() string {
    var sb strings.Builder
    sb.WriteString("Value(")
    sb.WriteString("v=")
    sb.WriteString(s.V.String())
    sb.WriteString(",vo=")
    if s.Vo != nil { 
        sb.WriteString(s.Vo.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",vo2=")
    if s.Vo2 != nil { 
        sb.WriteString(s.Vo2.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *Value) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}
