package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"
)

func is_prime(n uint64) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {

	}

	sn := math.Sqrt(float64(n))

	for i := uint64(3); float64(i) <= sn; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true

}

func gcd(m uint64, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(n)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func pollards_rho(n uint64) uint64 {
	if n == 1 {
		return 1
	}

	x := uint64(2)
	y := uint64(2)
	d := uint64(1)

	f := func(x uint64) uint64 {
		return (x*x + 1) % n
	}

	for d == 1 {
		x = f(x)
		y = f(f(y))
		z := uint64(math.Abs(float64(x - y)))
		d = gcd(z, n)
	}

	if d == n {
		return uint64(0)
	} else {
		return d
	}

}

func factoriz(x uint64) string {
	d := uint64(0)
	q := uint64(0)

	text := ""

	for x >= 4 && x%2 == 0 {
		text += "2 * "
		x = x / 2
	}

	d = 3
	q = x / d

	for q >= d {
		if x%d == 0 {
			text += strconv.Itoa(int(d)) + " * "
			x = q
		} else {
			d += 2
		}

		q = x / d
	}

	text += strconv.Itoa(int(x)) + " * "
	return text
}

func main() {
	start := time.Now()
	args := os.Args

	if len(args) < 2 {
		fmt.Println("need args[1].")
		os.Exit(1)
	}

	n, err := strconv.ParseUint(args[1], 10, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if n < 1 {
		fmt.Println("should be args[1] > 0.")
		os.Exit(1)
	}

	text := "Factorization: " + strconv.Itoa(int(n)) + " = 1 * "

	for {
		d := pollards_rho(n)

		if d == 0 || d == 1 {
			if is_prime(n) == true {
				text += strconv.Itoa(int(n)) + " * "
			} else {
				text += factoriz(n)
			}
			break

		} else {
			if is_prime(d) == true {
				text += strconv.Itoa(int(d)) + " * "
			} else {
				text += factoriz(d)
			}
		}

		n = n / d
	}
	textLength := len(text) - 3
	fmt.Println(text[:textLength])
	end := time.Now()
	fmt.Printf("Time: %f(sec)\n", (end.Sub(start)).Seconds())
}
