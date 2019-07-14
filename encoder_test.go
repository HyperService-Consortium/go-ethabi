package ethabi

import (
	"fmt"
	"math/big"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println(NewEncoder().Encode("uint64", uint64(2)))
	fmt.Println(NewEncoder().Encodes([]string{"uint256", "uint64"}, []interface{}{big.NewInt(666), uint64(2)}))
}

func TestDecode(t *testing.T) {
	x, _ := NewEncoder().Encodes([]string{"uint256", "uint64", "bytes"}, []interface{}{big.NewInt(666), uint64(2), []byte{1, 2, 3, 4, 55, 66, 73}})

	fmt.Println(NewDecoder().Decodes([]string{"uint256", "uint64", "bytes"}, x))
}
