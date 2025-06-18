package main

import (
	"fmt"
	"sort"
)

type targetObj struct {
	i int
	j int
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func task1(nums []int, targetNum int) []targetObj {
	var result []targetObj
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := num + nums[j]
			if sum == targetNum {
				result = append(result, targetObj{i: num, j: nums[j]})
			}
		}
	}
	return result
}

/*
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，
将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0                           // 慢指针，记录不重复元素的位置
	for j := 1; j < len(nums); j++ { // 快指针，遍历数组
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1 // 新长度
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较
，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func merge(intervals [][]int) [][]int {
	// 特殊情况处理
	if len(intervals) <= 1 {
		return intervals
	}

	// 按照区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果集
	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// 如果有重叠，合并区间
		if current[0] <= last[1] {
			// 更新结束位置为较大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 无重叠，直接添加
			merged = append(merged, current)
		}
	}

	return merged
}

/*
给定一个排序数组，你需要在原地删除重复出现的元素
*/
func plusOne(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] < 9 {
			nums[j]++
			return nums
		}
		nums[j] = 0
	}
	return append([]int{1}, nums...) // 如果所有位都是9，则需要在前面加1
}

// 查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准
	for i := 0; i < len(strs[0]); i++ {
		currentChar := strs[0][i]

		// 检查其他字符串的相同位置
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够或字符不匹配
			if i == len(strs[j]) || strs[j][i] != currentChar {
				return strs[0][:i]
			}
		}
	}

	return ""
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if len(stack) > 0 && stack[len(stack)-1] == mapping[char] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func findOnceNums(nums []int) int {
	numCount := make(map[int]int)

	for _, num := range nums {
		numCount[num]++
	}

	for num, count := range numCount {
		if count == 1 {
			return num
		}
	}
	return -1 // 如果没有找到，返回-1
}

func main() {

	//测试task1
	nums := []int{1, 2, 3, 4, 5}
	targetNum := 5
	results := task1(nums, targetNum)
	fmt.Println("task1 test result is :", results)

	// 测试removeDuplicates
	// 给定一个有序数组，删除重复元素
	// 返回新数组的长度
	removeDuplicatesNums := []int{1, 1, 2, 2, 3, 4, 4, 7, 8, 9, 10}
	removeDuplicatesResult := removeDuplicates(removeDuplicatesNums)
	fmt.Println("removeDuplicates result is ", removeDuplicatesResult, removeDuplicatesNums[:removeDuplicatesResult])

	test3 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {17, 20}}
	// 测试merge
	mergeResult := merge(test3)
	fmt.Println("merge result is ", mergeResult)

	plusNums := []int{9, 8, 9}
	// 测试plusOne
	plusResult := plusOne(plusNums)
	fmt.Println("plusOne result is ", plusResult)

	str := []string{"flower", "flow", "flight"}
	// 测试longestCommonPrefix
	longestPrefix := longestCommonPrefix(str)
	fmt.Println("longestCommonPrefix result is ", longestPrefix)

	// 测试isValid
	isValidStr := "(test"
	validResult := isValid(isValidStr)
	fmt.Println("validResult result is ", validResult)

	fmt.Println("findOnceNums result is ", findOnceNums([]int{4, 1, 2, 1, 2}))
}
