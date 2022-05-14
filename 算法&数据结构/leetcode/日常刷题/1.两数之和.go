package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int, len(nums))
	map1[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if (map1[target-nums[i]] != 0) || ((target - nums[i]) == nums[0]) {
			return []int{i, map1[target-nums[i]]}
		}
		map1[nums[i]] = i
	}
	return []int{}
}

// @lc code=end
