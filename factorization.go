package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

func isPrime(n big.Int) bool {
	tempBigInt := big.NewInt(0)

	if n.Cmp(big.NewInt(2)) == -1 { // n < 2
		return false
	} else if n.Cmp(big.NewInt(2)) == 0 { //n == 2
		return true
	} else if big.NewInt(0).Cmp(tempBigInt.Mod(&n, big.NewInt(2))) == 0 { // n % 2 == 0
		return false
	}

	sn := n.Sqrt(&n)
	for i := big.NewInt(3); i.Cmp(sn) != 1; i = i.Add(i, big.NewInt(2)) { // i = 3; i <= sn; i+=2
		if big.NewInt(0).Cmp(tempBigInt.Mod(&n, i)) == 0 {
			return false
		}
	}
	return true
}

func gcd(m big.Int, n big.Int) *big.Int {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	return z.GCD(x, y, &m, &n)
}

func pollards_rho(n big.Int) *big.Int {
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
		d = gcd(*z, n)
	}

	if d == &n {
		return big.NewInt(0)
	}

	return d
}

func factoriz(n big.Int) string {

	x := *new(big.Int).Set(&n)

	tempBigInt := big.NewInt(0)
	d := big.NewInt(0)
	q := big.NewInt(0)

	text := ""

	for x.Cmp(big.NewInt(3)) == 1 && big.NewInt(0).Cmp(tempBigInt.Mod(&x, big.NewInt(2))) == 0 { // x > 3 && x%2 == 0
		text += "2 * "
		x.Div(&x, big.NewInt(2))
	}

	d = big.NewInt(3)

	q.Div(&x, d)

	for q.Cmp(d) != -1 { //q >= d
		if big.NewInt(0).Cmp(tempBigInt.Mod(&x, d)) == 0 { //x%d == 0
			text += d.String() + " * "
			x.Set(q)
		} else {
			d.Add(d, big.NewInt(2))
		}
		q.Div(&x, d)
	}

	text += x.String() + " * "

	return text
}

func main() {
	start := time.Now()
	args := os.Args

	if len(args) < 2 {
		fmt.Println("need args[1].")
		os.Exit(1)
	}

	n := new(big.Int)
	n, ok := n.SetString(args[1], 10)
	if !ok {
		fmt.Println("args[1] nedd integer.")
		os.Exit(1)
	}

	if n.Cmp(big.NewInt(1)) == -1 { //n < 1
		fmt.Println("should be args[1] > 0.")
		os.Exit(1)
	}

	text := "Factorization: " + n.String() + " = 1 * "

	for {
		d := pollards_rho(*n)
		if d.Cmp(big.NewInt(2)) == -1 { // d < 2
			if isPrime(*n) == true {
				text += n.String() + " * "
			} else {
				text += factoriz(*n)
			}
			break
		} else {
			if isPrime(*d) == true {
				text += d.String() + " * "
			} else {
				text += factoriz(*d)
			}
			n.Div(n, d)
		}
	}
	textLength := len(text) - 3
	fmt.Println(text[:textLength])
	end := time.Now()
	fmt.Printf("Time: %f(sec)\n", (end.Sub(start)).Seconds())
}
