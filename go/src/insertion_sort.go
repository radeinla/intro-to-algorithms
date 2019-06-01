package main

import "fmt"

func insertionSort(arr []int) []int {
	fmt.Printf("%v\n", arr)

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
			fmt.Printf("%v\n", arr)
		}
	}

	return arr
}

func main() {
	fmt.Printf("%v\n", insertionSort([]int{5, 2, 4, 6, 1 ,3}))

	fmt.Printf("%v\n", insertionSort([]int{2, 1}))

	fmt.Printf("%v\n", insertionSort([]int{1}))

	fmt.Printf("%v\n", insertionSort([]int{}))
}
