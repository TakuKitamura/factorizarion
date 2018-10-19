package factorization

import (
	"errors"
	"math"
	"math/big"
	"strconv"
)

func UintIsPrime(n uint64) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		// pass
	}

	sn := math.Sqrt(float64(n))

	for i := uint64(3); float64(i) <= sn; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true

}

func UintGcd(m uint64, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(m)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func UintPollardsRho(n uint64) uint64 {
	if n == 1 {
		return 1
	}

	x := uint64(2)
	y := uint64(2)
	d := uint64(1)

	f := func(x uint64) uint64 {
		return (x*x + 1) % n
	}
	// i := uint64(0)
	// limit := uint64(math.Pow(float64(n), 0.25))
	// for d == 1 && i < limit {
	x = f(x)
	y = f(f(y))

	z := uint64(0)
	if x > y {
		z = x - y
	} else {
		z = y - x
	}

	d = UintGcd(z, n)
	// 	i += 1
	// }

	if d == n {
		return uint64(0)
	}

	return d

}

func UintTrialDivision(x uint64) uint64 {
	for x >= 4 && x%2 == 0 {
		return 2
	}

	d := uint64(3)
	q := x / d

	for q >= d {
		if x%d == 0 {
			return d
		}
		d += 2
		q = x / d
	}

	return x
}

func UintMain(n uint64) (string, error) {

	if n < 1 {
		err := errors.New("should be args[1] > 0.")
		return "", err
	}

	text := "Factorization: " + strconv.FormatUint(n, 10) + " = 1 * "

	for {
		d := UintPollardsRho(n)
		if d < 2 {
			if UintIsPrime(n) == true || n == 1 {
				text += strconv.FormatUint(n, 10) + " * "
				break
			} else {
				d = UintTrialDivision(n)
			}
		} else {
			if UintIsPrime(d) == false {
				d = UintTrialDivision(d)
			}
		}
		text += strconv.FormatUint(d, 10) + " * "
		n = n / d

	}
	textLength := len(text) - 3
	formula := text[:textLength]

	return formula, nil
}
