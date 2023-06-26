package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 5, 7, 6, 8, 10, 9}
	target := 6
	Result := linearSearch(arr, target)
	fmt.Println("Result :", Result)
}

func linearSearch(arr [9]int, target int) int {

	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
