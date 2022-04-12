package main

import "fmt"

//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	// 计算两个数组的总长度
//	totalLen := len(nums1) + len(nums2)
//
//	var a int // in nums1
//	var b int // in nums2
//
//	if totalLen%2 == 0 {
//		// 如果总长度为偶数，中位数的结果应该是(a+b)/2
//		// 要找到 totalLen/2 和 totalLen/2 + 1
//		i := 0
//		j := 0
//		count := 0
//		for ; count != totalLen / 2 && i < len(nums1) && j < len(nums2);{
//			if nums1[i] <= nums2[j] {
//				i++
//			}else {
//				j++
//			}
//			count++
//		}
//		if count != totalLen / 2 && i < len(nums1) {
//			for ;count != totalLen / 2 && i < len(nums1); {
//				i++
//				count++
//			}
//
//			return
//		}
//
//		return (float64(a) + float64(b)) / 2
//	} else {
//		// 如果总长度为奇数，中位数的结果应该是 a 或者 b
//		// 要找到 totalLen/2 + 1
//
//	}
//
//	return 0
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	temp := make([]int, 0)
	var index1 = 0
	var index2 = 0

	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] <= nums2[index2] {
			temp = append(temp, nums1[index1])
			index1++
		} else {
			temp = append(temp, nums2[index2])
			index2++
		}
	}
	for ; index1 < len(nums1); index1++ {
		temp = append(temp, nums1[index1])
	}
	for ; index2 < len(nums2); index2++ {
		temp = append(temp, nums2[index2])
	}

	if len(temp)%2 == 0 {
		return (float64(temp[len(temp)/2-1]) + float64(temp[len(temp)/2])) / 2
	} else {
		return float64(temp[len(temp)/2])
	}
}

func main() {
	array1 := []int{1, 3}
	array2 := []int{2}
	fmt.Println(findMedianSortedArrays(array1, array2))
}
