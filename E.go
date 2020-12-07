package main

import (
	"fmt"
	"os"
)

func main() {
	var number = 0
	var result = 0
	var n = 0
	fmt.Fscan(os.Stdin, &n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(os.Stdin, &str)
		len := len(str)
		number = number + len + 1
		if result < len {
			result = len
		}
		if number >= n {
			break
		}
	}

	fmt.Print(result)
}
