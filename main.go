package main

import "fmt"

func main() {
	arr := []int{-1, 1, 0, -3, 3}
	var res []int
	res = productExceptSelf(arr)
	fmt.Println(res)
}
func productExceptSelf(nums []int) []int {
	var res []int
	total := 1
	for _, v := range nums {
		if v != 0 {

			total *= v
		}
	}
	for _, v := range nums {
		if v != 0 {
			res = append(res, total/v)
		} else {
			res = append(res, total)
		}

	}
	return res
}
