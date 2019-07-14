package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func reverseBigInt(x *big.Int, y *big.Int) {
	runes := []rune(x.Text(10))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	y.SetString(string(runes), 10)

}

func testCrankyNumber(x *big.Int, y *big.Int) bool {
	runes := []rune(x.Text(10))
	j := new(big.Int)
	k := new(big.Int)
	z := new(big.Int)
	for i := 1; i < len(runes); i++ {
		j.SetString(string(runes[0:i]), 10)
		k.SetString(string(runes[i:]), 10)
		if y.Cmp(z.Mul(j, k)) == 0 {
			return true
		}
	}
	return false
}

func setBigNumber(x *big.Int, y int64) {
	x.Exp(big.NewInt(10), big.NewInt(y), nil)
}

func crankyNumber(x *big.Int, y *big.Int) bool {
	reverseBigInt(x, y)
	if testCrankyNumber(x, y) {
		fmt.Printf("--- Num (%s), RNum (%s)\n", x.Text(10), y.Text(10))
		return true
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	//args := os.Args[1:]

	st := os.Args[1]
	end := os.Args[2]

	//setBigNumber(limit, 14)
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

	one := big.NewInt(1)
	//	for x := int64(0); x < 338752614; x++ {
	//x := big.NewInt(11)
	for ; start.Cmp(limit) <= 0; start.Add(start, one) {
		//	num := big.NewInt(x)
		//	y := new(big.Int)
		crankyNumber(start, new(big.Int))
	}

}
