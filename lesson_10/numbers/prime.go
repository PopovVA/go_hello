package numbers

import (
	"fmt"
	"math/big"
)

func CountPrimes(n int64) (res int64) {

	var i int64
	for i = 0; i <= n; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			fmt.Println(i, " - простое")
			res++
		}
	}
	return res
}
