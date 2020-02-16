package main

import "fmt"


type searchRequest struct {
	haystack []int
	needle int
}

type searchResult struct {
	found bool
	index int
	searchRequest searchRequest
}

func (per *searchResult) show() {
	if (per.found) {
		fmt.Printf("%dで見つけた\n", per.index)
	} else {
		fmt.Printf("見つけなかった\n", per.index)
	}
}

func (searchRequest *searchRequest) execute() searchResult {
	for i := 0; i < len(searchRequest.haystack); i++ {
		if searchRequest.haystack[i] == searchRequest.needle {
			return searchResult{true, i, *searchRequest}
		}
	}

	return searchResult{ false, -1, *searchRequest}
}


func (searchRequest *searchRequest) show() {
	fmt.Printf("%vで%dを探して\n", searchRequest.haystack, searchRequest.needle)
}



func main() {
	searchRequest := searchRequest {
		[]int {1, 2, 5, 7},
		5,
	}

	searchRequest.show()
	searchResult := searchRequest.execute()
	searchResult.show()
}
