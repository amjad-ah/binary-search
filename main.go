package main

import (
	"fmt"
)

func search(arr []int, n int) func() (bool, int) {
	r, l := len(arr)-1, 0
	var i int
	return func() (bool, int) {
		i = (r + l) / 2
		if arr[i] == n {
			return true, i
		} else if arr[i] < n {
			l = i + 1
		} else if arr[i] > n {
			r = i - 1
		} else if i == (r+l)/2 {
			return true, -1
		}
		return false, 0
	}
}

func main() {
	arr := []int{1, 2, 3, 6, 7, 9, 12, 21, 39} // your SORTED array
	n := 9                                     // the number you would like to search for

	search := search(arr, n)
	checker := false
	var i int
	for !checker {
		checker, i = search()
		if checker {
			if i < 0 {
				fmt.Println("Not Found!")
			} else {
				fmt.Printf("Found, index is: %d\n", i)
			}
		}
	}
}
