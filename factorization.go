package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"
)

func bigintIsPrime(n big.Int) bool {
	tempBigInt := big.NewInt(0)

	if n.Cmp(big.NewInt(2)) == -1 { // n < 2
		return false
	} else if n.Cmp(big.NewInt(2)) == 0 { //n == 2
		return true
	} else if big.NewInt(0).Cmp(tempBigInt.Mod(&n, big.NewInt(2))) == 0 { // n % 2 == 0
		return false
	}

	sn := new(big.Int).Sqrt(&n)

	for i := big.NewInt(3); i.Cmp(sn) != 1; i = i.Add(i, big.NewInt(2)) { // i = 3; i <= sn; i+=2
		if big.NewInt(0).Cmp(tempBigInt.Mod(&n, i)) == 0 {
			return false
		}
	}
	return true
}

func bigintGcd(m big.Int, n big.Int) *big.Int {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	return z.GCD(x, y, &m, &n)
}

func bigintPollardsRho(n big.Int) *big.Int {
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

	for d.Cmp(big.NewInt(1)) == 0 {
		x = f(x)
		y = f(f(y))
		z := big.NewInt(0)

		if x.Cmp(y) == 1 { // x > y
			z.Sub(x, y)
		} else {
			z.Sub(y, x)
		}
		d = bigintGcd(*z, n)
	}

	if d == &n {
		return big.NewInt(0)
	}

	return d

}

func bigintTrialDivision(n big.Int) *big.Int {

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

func uintIsPrime(n uint64) bool {
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

func uintGcd(m uint64, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(m)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func uintPollardsRho(n uint64) uint64 {
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
	// limit := uint64(math.Pow(float64(n), 0.05))
	// for d == 1 && i < limit {
	x = f(x)
	y = f(f(y))

	z := uint64(0)
	if x > y {
		z = x - y
	} else {
		z = y - x
	}

	d = uintGcd(z, n)
	// 	i += 1
	// }
	// fmt.Println(limit, i, 123)

	if d == n {
		return uint64(0)
	}

	return d

}

func uintTrialDivision(x uint64) uint64 {
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

func uintMain(n uint64) (string, error) {

	if n < 1 {
		err := errors.New("should be args[1] > 0.")
		return "", err
	}

	text := "Factorization: " + strconv.FormatUint(n, 10) + " = 1 * "

	for {
		d := uintPollardsRho(n)
		if d < 2 {
			if uintIsPrime(n) == true || n == 1 {
				text += strconv.FormatUint(n, 10) + " * "
				break
			} else {
				d = uintTrialDivision(n)
			}
		} else {
			if uintIsPrime(d) == false {
				d = uintTrialDivision(d)
			}
		}
		text += strconv.FormatUint(d, 10) + " * "
		n = n / d

	}
	textLength := len(text) - 3
	formula := text[:textLength]

	return formula, nil
}

func bigintMain(n big.Int) (string, error) {

	if n.Cmp(big.NewInt(1)) == -1 { //n < 1
		err := errors.New("should be args[1] > 0.")
		return "", err
	}

	text := "Factorization: " + n.String() + " = 1 * "

	for {
		d := bigintPollardsRho(n)
		if d.Cmp(big.NewInt(2)) == -1 { // d < 2
			if bigintIsPrime(n) == true || n.Cmp(big.NewInt(1)) == 0 {
				text += n.String() + " * "
				break
			} else {
				d = bigintTrialDivision(n)
			}
		} else {
			if bigintIsPrime(*d) == false {
				d = bigintTrialDivision(*d)
			}
		}
		text += d.String() + " * "
		n.Div(&n, d)
	}

	textLength := len(text) - 3
	formula := text[:textLength]

	return formula, nil
}

func main() {
	start := time.Now()
	args := os.Args

	if len(args) < 2 {
		log.Fatal("need args[1].")
		os.Exit(1)
	}

	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(args[1], 10)
	if !ok {
		log.Fatal("invalid integer.")
		os.Exit(1)
	}

	maxUint64 := new(big.Int).SetUint64(math.MaxUint64)

	if bigInt.Cmp(maxUint64) == -1 {
		uInt64, err := strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		result, err := uintMain(uInt64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(result)
	} else {
		result, err := bigintMain(*bigInt)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(result)
	}

	end := time.Now()
	fmt.Printf("Time: %f(sec)\n", (end.Sub(start)).Seconds())
	os.Exit(0)
}
