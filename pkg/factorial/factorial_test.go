package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		numbrer   int64
		exptected string
	}{
		{
			numbrer:   0,
			exptected: "1",
		},
		{
			numbrer:   1,
			exptected: "1",
		},
		{
			numbrer:   6,
			exptected: "720",
		},
		{
			numbrer:   7,
			exptected: "5040",
		},
	}

	for _, testCase := range testCases {
		actual := Calculate(testCase.numbrer)
		if actual.String() != testCase.exptected {
			t.Errorf("%d! should be %s, got %s.", testCase.numbrer, testCase.exptected, actual)
		}
	}
}
