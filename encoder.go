package ethabi

import (
	"reflect"

	libabi "github.com/Myriad-Dreamin/go-ethabi/ethlib"
)

type Encoder struct {
}

func NewEncoder() *Encoder {
	return new(Encoder)
}

func (enc *Encoder) Encode(desc string, val interface{}) (output []byte, err error) {
	typ, err := libabi.NewType(test.typ, test.components)
	if err != nil {
		return
	}
	output, err = typ.pack(reflect.ValueOf(test.input))
	if err != nil {
		return
	}
	return
}
