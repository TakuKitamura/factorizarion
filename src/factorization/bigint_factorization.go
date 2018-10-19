package factorization

import (
	"errors"
	"math/big"
)

func BigintIsPrime(n big.Int) bool {
	tempBigInt := big.NewInt(0)

	if n.Cmp(big.NewInt(2)) == -1 { // n < 2
		return false
	} else if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 { //n == 2
		return true
	} else if big.NewInt(0).Cmp(tempBigInt.Mod(&n, big.NewInt(2))) == 0 || big.NewInt(0).Cmp(tempBigInt.Mod(&n, big.NewInt(3))) == 0 { // n % 2 == 0
		return false
	}

	i := big.NewInt(5)
	w := big.NewInt(2)

	for i.Mul(i, i).Cmp(&n) != 1 {
		if big.NewInt(0).Cmp(tempBigInt.Mod(&n, i)) == 0 {
			return false
		} else {
			i.Add(i, w)
			w.Sub(big.NewInt(6), w)
		}
	}

	// sn := new(big.Int).Sqrt(&n)

	// for i := big.NewInt(3); i.Cmp(sn) != 1; i = i.Add(i, big.NewInt(2)) { // i = 3; i <= sn; i+=2
	// 	if big.NewInt(0).Cmp(tempBigInt.Mod(&n, i)) == 0 {
	// 		return false
	// 	}
	// }
	return true
}

func BigintGcd(m big.Int, n big.Int) *big.Int {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	return z.GCD(x, y, &m, &n)
}

func BigintPollardsRho(n big.Int) *big.Int {
	if n.Cmp(&n) == 1 { //n == 1
		return big.NewInt(1)
	}

	x := big.NewInt(2)
	y := big.NewInt(2)
	d := big.NewInt(1)

	f := func(x *big.Int) *big.Int {
		xx := x.Mul(x, x)
		xxPlusOne := xx.Add(xx, big.NewInt(1))
		return xxPlusOne.Mod(xxPlusOne, &n)
	}

	i := big.NewInt(0)
	sqrt := new(big.Int).Sqrt(&n)
	limit := sqrt.Sqrt(sqrt)

	for d.Cmp(big.NewInt(1)) == 0 && i.Cmp(limit) == -1 {
		x = f(x)
		y = f(f(y))
		z := big.NewInt(0)

		if x.Cmp(y) == 1 { // x > y
			z.Sub(x, y)
		} else {
			z.Sub(y, x)
		}

		d = BigintGcd(*z, n)
		i.Add(i, big.NewInt(1))
	}

	if d == &n {
		return big.NewInt(0)
	}

	return d

}

func BigintTrialDivision(n big.Int) *big.Int {

	x := *new(big.Int).Set(&n)

	tempBigInt := big.NewInt(0)
	d := big.NewInt(0)
	q := big.NewInt(0)

	for x.Cmp(big.NewInt(3)) == 1 && big.NewInt(0).Cmp(tempBigInt.Mod(&x, big.NewInt(2))) == 0 { // x > 3 && x%2 == 0
		return big.NewInt(2)
	}

	d = big.NewInt(3)

	q.Div(&x, d)

	for q.Cmp(d) != -1 { //q >= d
		if big.NewInt(0).Cmp(tempBigInt.Mod(&x, d)) == 0 { //x%d == 0
			return d
		} else {
			d.Add(d, big.NewInt(2))
		}
		q.Div(&x, d)
	}

	return &x
}

func BigintMain(x big.Int) (string, error) {

	n := *new(big.Int).Set(&x)

	if n.Cmp(big.NewInt(1)) == -1 { //n < 1
		err := errors.New("should be args[1] > 0.")
		return "", err
	}
	buf := make([]byte, 0)
	buf = append(buf, "Factorization: "+n.String()+" = 1 * "...)

	for {
		d := BigintPollardsRho(n)
		if d.Cmp(big.NewInt(2)) == -1 { // d < 2
			if BigintIsPrime(n) == true || n.Cmp(big.NewInt(1)) == 0 {
				buf = append(buf, n.String()+" * "...)
				break
			} else {
				d = BigintTrialDivision(n)
			}
		} else {
			if BigintIsPrime(*d) == false {
				d = BigintTrialDivision(*d)
			}
		}
		buf = append(buf, d.String()+" * "...)
		n.Div(&n, d)
	}
	text := string(buf)
	textLength := len(text) - 3
	formula := text[:textLength]

	return formula, nil
}
