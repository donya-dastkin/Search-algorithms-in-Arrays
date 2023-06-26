package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	Result := linearSearch(arr, 3)
	fmt.Println(Result)
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
