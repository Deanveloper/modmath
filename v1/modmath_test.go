// Each of these tests just come out of either intuition or from my notes
// in cryptography class... don't judge me
package modmath_test


import (
	"testing"
	"github.com/deanveloper/nikola"
	. "github.com/deanveloper/modmath/v1"
)

func TestLpr(t *testing.T) {
	nikola.SuggestEqual(t, 0, Mod(0, 5))
	nikola.SuggestEqual(t, 1, Mod(1, 10))
	nikola.SuggestEqual(t, 2, Mod(202, 10))
	nikola.SuggestEqual(t, 3, Mod(47291873, 4729187))
}

func TestSolve(t *testing.T) {
	var r int
	var e error

	r, e = Solve(3, 5, 10)
	nikola.SuggestEqual(t, 5, r)
	nikola.SuggestEqual(t, nil, e)

	r, e = Solve(20, 5, 25)
	nikola.SuggestEqual(t, 4, r)
	nikola.SuggestEqual(t, nil, e)

	r, e = Solve(20, 5, 30)
	nikola.SuggestEqual(t, 0, r)
	nikola.SuggestEqual(t, NoSolution, e)
}

func TestSolveExp(t *testing.T) {
	nikola.SuggestEqual(t, 4, Power(7, 365, 9))
	nikola.SuggestEqual(t, 5, Power(5, 291, 11))
}

func TestSolveCrt(t *testing.T) {
	nikola.SuggestEqual(t, 7, ChineseRemainder(2, 5, 1, 3))
	nikola.SuggestEqual(t, 5871, ChineseRemainder(12, 93, 29, 127))
	nikola.SuggestEqual(t, 1316, ChineseRemainder(5, 23, 70, 89))
}

func TestSolveCrtMany(t *testing.T) {
	// I only have one example in my notes of this section
	nikola.SuggestEqual(t, 49, ChineseRemainderMany([]CrtEntry{{1, 3}, {4, 5}, {0, 7}}))
}