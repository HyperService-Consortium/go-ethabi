package ethabi

import (
	abi "github.com/Myriad-Dreamin/go-ethabi/ethlib"
)

type Decoder struct {
}

func NewDecoder() *Decoder {
	return new(Decoder)
}

func (enc *Decoder) Decodes(descs []string, rawInput []byte) ([]interface{}, error) {
	var args = make([]abi.Argument, 0, len(descs))
	for _, desc := range descs {
		// todo: add component support
		typ, err := abi.NewType(desc, nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{Type: typ})
	}
	return abi.Arguments(args).UnpackValues(rawInput)
}
