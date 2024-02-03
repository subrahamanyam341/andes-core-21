//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/dharitri-org/protobuf/protobuf  --gogoslick_out=. dct.proto
package dct

import (

    "math/big"
dct  "github.com/subrahamanyam341/andes-core-21/data/dct/proto"

)
// New returns a new batch from given buffers
func New() *dct.DCToken {
	return &dct.DCToken{
		Value: big.NewInt(0),
	}
}
