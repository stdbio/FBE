//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: extra.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package extra

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

// Info key
type InfoKey struct {
}

// Convert Info flags key to string
func (k *InfoKey) String() string {
    var sb strings.Builder
    sb.WriteString("InfoKey(")
    sb.WriteString(")")
    return sb.String()
}

// Info struct
type Info struct {
    Info string `json:"info"`
    Extra *Extra `json:"extra"`
    Extras []*Extra `json:"extras"`
    Extras1 []*Extra `json:"extras1"`
}

// Create a new Info struct
func NewInfo() *Info {
    return &Info{
        Info: "",
        Extra: nil,
        Extras: make([]*Extra, 0),
        Extras1: make([]*Extra, 0),
    }
}

// Create a new Info struct from the given field values
func NewInfoFromFieldValues(infoV string, extraV *Extra, extrasV []*Extra, extras1V []*Extra) *Info {
    return &Info{infoV, extraV, extrasV, extras1V}
}

// Create a new Info struct from JSON
func NewInfoFromJSON(buffer []byte) (*Info, error) {
    result := *NewInfo()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *Info) Copy() *Info {
    var result = *s
    return &result
}

// Struct deep clone
func (s *Info) Clone() *Info {
    // Serialize the struct to the FBE stream
    writer := NewInfoModel(fbe.NewEmptyBuffer())
    _, _ = writer.Serialize(s)

    // Deserialize the struct from the FBE stream
    reader := NewInfoModel(writer.Buffer())
    result, _, _ := reader.Deserialize()
    return result
}

// Get the struct key
func (s *Info) Key() InfoKey {
    return InfoKey{
    }
}

// Convert struct to optional
func (s *Info) Optional() *Info {
    return s
}

// Convert struct to optional
func (s *Info) Ptr() *Info {
    return s
}

// Get the FBE type
func (s *Info) FBEType() int { return 1 }

// Convert struct to string
func (s *Info) String() string {
    var sb strings.Builder
    sb.WriteString("Info(")
    sb.WriteString("info=")
    sb.WriteString("\"" + s.Info + "\"")
    sb.WriteString(",extra=")
    if s.Extra != nil { 
        sb.WriteString(s.Extra.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",extras=")
    if s.Extras != nil {
        first := true
        sb.WriteString("[" + strconv.FormatInt(int64(len(s.Extras)), 10) + "][")
        for _, v := range s.Extras {
            if v != nil { 
                if first { sb.WriteString("") } else { sb.WriteString(",") }
                sb.WriteString(v.String())
            } else {
                if first { sb.WriteString("") } else { sb.WriteString(",") }
                sb.WriteString("null")
            }
            first = false
        }
        sb.WriteString("]")
    } else {
        sb.WriteString(",extras=[0][]")
    }
    sb.WriteString(",extras1=")
    if s.Extras1 != nil {
        first := true
        sb.WriteString("[" + strconv.FormatInt(int64(len(s.Extras1)), 10) + "][")
        for _, v := range s.Extras1 {
            if v != nil { 
                if first { sb.WriteString("") } else { sb.WriteString(",") }
                sb.WriteString(v.String())
            } else {
                if first { sb.WriteString("") } else { sb.WriteString(",") }
                sb.WriteString("null")
            }
            first = false
        }
        sb.WriteString("]")
    } else {
        sb.WriteString(",extras1=[0][]")
    }
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *Info) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}
