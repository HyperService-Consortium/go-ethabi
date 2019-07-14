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
