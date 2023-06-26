package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	Result := TernarySearch(arr, target)
	fmt.Println("Result :", Result)
}

func TernarySearch(arr [10]int, target int) int {
	left := 0
	right:= len(arr) - 1
	for left <= right {
		mid1 := left + ((right - left) / 3)
		mid2 := right - ((right - left) / 3)
		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}
		if target < arr[mid1] {
			right = mid1 - 1
		} else if target > arr[mid2] {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}
	}
	return -1
}
