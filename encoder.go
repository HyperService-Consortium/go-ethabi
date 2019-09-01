package ethabi

import (
	"fmt"
	"reflect"

	abi "github.com/HyperService-Consortium/go-ethabi/ethlib"
)

type Encoder struct {
}

func NewEncoder() *Encoder {
	return new(Encoder)
}

func (enc *Encoder) Encode(desc string, val interface{}) ([]byte, error) {

	// todo: add component support
	typ, err := abi.NewType(desc, nil)
	if err != nil {
		return nil, err
	}
	return typ.Pack(reflect.ValueOf(val))
}

func (enc *Encoder) Encodes(descs []string, vals []interface{}) ([]byte, error) {
	if len(descs) != len(vals) {
		return nil, fmt.Errorf("argument count mismatch: %d for %d", len(descs), len(vals))
	}
	var args = make([]abi.Argument, 0, len(descs))
	for _, desc := range descs {
		// todo: add component support
		typ, err := abi.NewType(desc, nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{Type: typ})
	}
	return abi.Arguments(args).Pack(vals...)
}
