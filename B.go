package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a = 0
		b = 0
		c = 0
	)
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
	fmt.Fscan(os.Stdin, &c)

	result := isChet(a)
	if result {
		result = result && isChet(b) && isChet(c)
		printAnswer(result)
	} else {
		result = result || isChet(b) || isChet(c)
		printAnswer(!result)
	}
}

func printAnswer(result bool) {
	if result {
		fmt.Print("WIN")
	} else {
		fmt.Print("FAIL")
	}
}

func isChet(elem int) bool {
	return elem%2 == 0
}
