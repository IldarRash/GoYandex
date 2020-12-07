package main

import (
	"fmt"
	"os"
)

func main() {
	result := 0
	var n = 0
	fmt.Fscan(os.Stdin, &n)
	var arr [100000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &arr[i])
	}

	for i, v := range arr {
		var flag bool
		if i == 0 {
			flag = checkHaos(0, v, arr[i+1], true, false)
		} else if i == n-1 {
			flag = checkHaos(arr[i-1], v, 0, false, true)
		} else if i != 0 && i < n {
			flag = checkHaos(arr[i-1], v, arr[i+1], false, false)
		}
		if flag {
			result++
		}
	}
	fmt.Print(result)
}

func checkHaos(prev int, curr int, next int, first bool, last bool) bool {
	if first {
		return curr > next
	}
	if last {
		return prev < curr
	}
	return prev < curr && curr > next
}
