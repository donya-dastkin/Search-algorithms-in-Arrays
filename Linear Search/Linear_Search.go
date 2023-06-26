package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 5, 7, 6, 8, 10, 9}
	x := 6
	Result := linearSearch(arr, x)
	fmt.Println("Result :", Result)
}

func linearSearch(arr [5]int, x int) int {

	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}
