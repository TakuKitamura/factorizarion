package main

import (
	"factorizarion/src/factorization"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args

	if len(args) < 2 {
		log.Fatal("need args[1].")
		os.Exit(1)
	}

	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(args[1], 10)
	if !ok || big.NewInt(0).Cmp(bigInt) != -1 {
		log.Fatal("invalid integer.")
		os.Exit(1)
	}

	maxUint64 := new(big.Int).SetUint64(math.MaxUint64)

	if bigInt.Cmp(maxUint64) == -1 {
		fmt.Println("Type: Uint64")
		uInt64, err := strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		result, err := factorization.UintMain(uInt64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(result)
	} else {
		fmt.Println("Type: BigInt")
		result, err := factorization.BigintMain(*bigInt)
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
