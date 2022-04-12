package main

import (
	"fmt"
	"sort"
)

//这种方法采用了两个循环嵌套，时间复杂度n方
func twoSum1(nums []int, target int) []int {
	for index, value := range nums {
		for j := index + 1; j < len(nums); j++ {
			if nums[j] == target-value {
				return []int{index, j}
			}
		}
	}
	return []int{}
}

// 先对数组进行排序，然后在寻找target-nums[i]的过程中使用二分查找，这样可以将时间复杂度降低到nlogn
func twoSum2(nums []int, target int) []int {
	sort.Ints(nums)
	var left int
	var right int
	var middle int

	for i := 0; i < len(nums)-1; i++ {
		left = i + 1
		right = len(nums) - 1

		for left <= right {
			middle = (left + right) / 2
			if nums[middle] == target-nums[i] {
				return []int{i, middle}
			} else if nums[middle] < target-nums[i] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return []int{}
}

// 使用哈希表对数组中所有对数及对应对下标进行存储，这样在找target-nums[i]的过程时间复杂度就是O1
// 并不是先进行哈希映射，而是边找边存，如果在哈希表中找不到target-nums[i]的话，就将现在正在查询的元素存储进去
func twoSum3(nums []int, target int) []int {
	//map[int]int -> map[value]index
	temp := make(map[int]int)

	for index, value := range nums {
		p, ok := temp[target-value] // p: index
		if ok {
			return []int{index, p}
		}

		temp[value] = index
	}

	return nil
}

func main() {
	myArray := []int{0, 4, 3, 0}
	target := 0

	fmt.Println(twoSum1(myArray, target))
	//fmt.Println(twoSum2(myArray, target))
	fmt.Println(twoSum3(myArray, target))
}
