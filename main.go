package main

import "fmt"

func findTheMedian(arr []int) int {
	// Write your code here
	var median int
	if len(arr)%2 == 0 {
		median = (arr[len(arr)/2] + arr[len(arr)/2-1]) / 2
	} else {
		median = arr[len(arr)/2]
	}
	return median
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findTheMedian(arr))
}
