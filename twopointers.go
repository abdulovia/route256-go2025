package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) == 2 {
		fmt.Printf("Indices of elements that sum up to %d: %v\n", target, result)
	} else {
		fmt.Printf("No two elements sum up to %d\n", target)
	}
}
