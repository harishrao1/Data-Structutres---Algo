package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{3, 4, 6, 8}
	result := mergeSoredArrays(arr1, arr2)
	fmt.Println("Merged array:", result)
}

func mergeSoredArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	merged := make([]int, 0)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}
	return merged
}
