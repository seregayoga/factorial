package factorial

import (
	"math/big"
)

// Calculate calculates factorial
func Calculate(n int64) *big.Int {
	return big.NewInt(0).MulRange(1, n)
}
