// https://leetcode.com/problems/two-sum/
package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, inMap := m[target-num]; inMap {
			return []int{v, i}
		} else {
			m[num] = i
		}
	}
	return nil
}

//func twoSum(nums []int, target int) []int {
//	lenght := len(nums)
//	for i := 0; i < lenght; i++ {
//		for j := i + 1; j < lenght; j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}
