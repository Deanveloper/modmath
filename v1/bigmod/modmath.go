// A package that does modular arithmetic with math/big integers.
package bigmod

import (
	"errors"
	"math/big"
)

var NoSolution = errors.New("no solution")

// Solves the equation ax=b mod n. Note that
// if there are multiple LPR solutions that the
// lowest one is returned. If there are no solutions,
// then (nil, NoSolution) is returned
func Solve(a, b, m *big.Int) (*big.Int, error) {
	gcd := new(big.Int).GCD(nil, nil, a, m)

	// If a and m are coprime, just multiply by the inverse
	if gcd.IsUint64() && gcd.Uint64() == 1 {
		aInv := new(big.Int).ModInverse(a, m)
		aInvB := new(big.Int).Mul(aInv, b)

		return aInvB.Mod(aInvB, m), nil
	}

	// If gcd divides b evenly, then solve a/d x = b/d mod m/d (d = gcd)
	if new(big.Int).Mod(b, gcd).Sign() == 0 {
		ad := new(big.Int).Div(a, gcd)
		bd := new(big.Int).Div(b, gcd)
		nd := new(big.Int).Div(m, gcd)
		return Solve(ad, bd, nd)
	}

	// else, no solution
	return nil, NoSolution
}
