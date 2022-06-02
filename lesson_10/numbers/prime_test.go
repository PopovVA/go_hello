package numbers

import (
	"testing"
)

type testCasePrimesCount struct {
	input  int64
	output int64
}

func TestPrimesCount(t *testing.T) {
	testCases := []testCasePrimesCount{
		{
			10,
			4,
		},
		{
			15,
			6,
		},
	}

	for _, cse := range testCases {
		res := CountPrimes(cse.input)
		if res != cse.output {
			t.Fatalf("Result incorrect, expected: %d, actual: %d", cse.output, res)
		}
	}
}
