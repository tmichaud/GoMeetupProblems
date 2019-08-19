package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func reverseBigInt(x *big.Int, y *big.Int) rune {
	runes := []rune(x.Text(10))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	y.SetString(string(runes), 10)
	return runes[len(runes)-1]
}

var Ft = map[rune]map[rune]rune{}

func setupFastTest() {
	s := "0123456789"
	for _, x := range s {
		Ft[x] = map[rune]rune{}
		for _, y := range s {
			Ft[x][y] = rune(((x-'0')*(y-'0'))%10 + '0')
		}
	}
}

// Check the last digit - if it matches - do the longer test
func fastTest(r1, r2, r3 rune) bool {
	// do array lookup of r1,r2 and check if it matchs r3
	return Ft[r1][r2] == r3
}

func testCrankyNumber(x *big.Int, y *big.Int, r rune) bool {
	runes := []rune(x.Text(10))
	j := new(big.Int)
	k := new(big.Int)
	z := new(big.Int)
	for i := 1; i < len(runes); i++ {
		if !fastTest(runes[i], runes[len(runes)-1], r) {
			return false
		}

		j.SetString(string(runes[0:i]), 10)
		k.SetString(string(runes[i:]), 10)
		if y.Cmp(z.Mul(j, k)) == 0 {
			return true
		}
	}
	return false
}

func crankyNumber(x *big.Int, y *big.Int) bool {
	r := reverseBigInt(x, y)
	if testCrankyNumber(x, y, r) {
		fmt.Printf("--- Num (%s), RNum (%s)\n", x.Text(10), y.Text(10))
		return true
	}
	return false
}

func testNumbers(start, limit *big.Int) {
	one := big.NewInt(1)
	for ; start.Cmp(limit) <= 0; start.Add(start, one) {
		crankyNumber(start, new(big.Int))
	}
}

func setBigNumber(x *big.Int, y int64) {
	x.Exp(big.NewInt(10), big.NewInt(y), nil)
}

func main() {
	fmt.Println("vim-go")
	setupFastTest()

	st := os.Args[1]
	end := os.Args[2]

	start := new(big.Int)
	limit := new(big.Int)

	sti, err := strconv.Atoi(st)
	if err != nil {
		log.Fatal(err)
	}

	endi, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}

	setBigNumber(start, int64(sti))
	setBigNumber(limit, int64(endi))

	fmt.Printf("Start is (%s) Limit is (%s)\n", start.Text(10), limit.Text(10))

	testNumbers(start, limit)

}
