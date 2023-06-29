package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	Result := BinarySearch(arr, target)
	fmt.Println("Result :", Result)
	Result2 := binarySearch(arr, target)
	fmt.Println("Result :", Result2)
}

func BinarySearch(arr [10]int, target int) int {
	var left int = 0
	var right int = len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}


//func binarySearch(arr [10]int, target int) int {
//	var left int = 0
//	var right int = len(arr) - 1
//	if left > right {
//		return -1
//	}
//	mid := (left + right) / 2
//	if arr[mid] == target {
//		return mid
//	}
//	if arr[mid] > target {
//		return binarySearch(arr, target)
//	}
//	if arr[mid] < target {
//		return binarySearch(arr, target)
//	}
//	return -1
//}