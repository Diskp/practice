package main

import "fmt"

func testRemoveDuplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	x := removeDuplicates(nums)
	fmt.Println(nums, x)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	cur := 10001
	index := 0
	for i := 0; i < length; i++ {
		if nums[index] != cur {
			cur = nums[index]
			index++
		} else {
			if index == len(nums)-1 {
				nums = nums[:index]
			} else {
				nums = append(nums[:index], nums[index+1:]...)
			}
		}
	}
	return len(nums)
}
