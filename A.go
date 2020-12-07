package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a = 0
		x = 0
		b = 0
		c = 0
	)
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &x)
	fmt.Fscan(os.Stdin, &b)
	fmt.Fscan(os.Stdin, &c)

	var result = a*(x*x) + b*x + c
	fmt.Println(result)
}
