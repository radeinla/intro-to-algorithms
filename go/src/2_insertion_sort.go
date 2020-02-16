package main

import "fmt"

func insertionSort(arr []int) []int {
	fmt.Printf("初期状態： %v\n", arr)

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				fmt.Printf("行動した： %v\n", arr)
			}
		}
		fmt.Printf("繰り返し\n")
	}

	return arr
}

func main() {
	fmt.Printf("最終状態: %v\n", insertionSort([]int{5, 2, 4, 6, 1 ,3}))

	fmt.Printf("最終状態: %v\n", insertionSort([]int{2, 1}))

	fmt.Printf("最終状態: %v\n", insertionSort([]int{1}))

	fmt.Printf("最終状態: %v\n", insertionSort([]int{}))
}
