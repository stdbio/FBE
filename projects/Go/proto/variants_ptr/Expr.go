//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants_ptr.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package variants_ptr

import "fmt"
import "strconv"
// import "strings"
import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// type Expr interface{}
type Expr interface{}
// List of Expr types
// bool
// string
// int32

// Create a new Expr variant
func NewExpr() *Expr {
    return nil
}

// Get the variant index
func GetExprIndex() int {
    return 0
}
