package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	nFlag := flag.Int("n", 10000, "an integer")
	flag.Parse()
	n := *nFlag
	if n < 3 || n > 100000 {
		fmt.Println("Please provide a number between 3 and 100000")
		os.Exit(1)
	}
	nthValue := returnNthValue(n)
	fmt.Printf("The %dth value in the sequence is: %s\n", n, nthValue.String())
	os.Exit(0)
}

func returnNthValue(n int) *big.Int {
	curr := new(big.Int)
	minus1 := big.NewInt(2)
	minus2 := big.NewInt(1)
	for i := 2; i < n; i++ {
		if (i+1)%2 == 0 {
			curr.Add(minus1, minus2)
		} else {
			curr.Mul(minus1, big.NewInt(2))
			curr.Sub(curr, minus2)
		}
		minus2.Set(minus1)
		minus1.Set(curr)
	}
	return curr
}
