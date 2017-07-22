package main

import (
	"fmt"
)

func main() {
	// res1, res2 := SearchforaRange([]int{5, 6, 7, 8, 8, 10}, 8)
	// fmt.Println(res1)
	// fmt.Println(res2)
	// res := SearchInsertPosition([]int{1, 3, 5, 6}, 0)
	// fmt.Println(res)
	// sa := []int{3, 2, 5, 6, 6, 4, 3, 1, 7}
	// quickSort(sa)
	// fmt.Println(sa)
	res := yihuo([]int{2, 2, 5, 3, 1, 1, 3})
	fmt.Println(res)
}

/*
Given a sorted array of integers, find the starting and ending position of a given target value.
Your algorithm’s runtime complexity must be in the order of O(log n).
If the target is not found in the array, return [-1, -1].
For example, Given [5, 7, 7, 8, 8, 10] and target value 8, return [3, 4].
**/
func SearchforaRange(sa []int, target int) (start, end int) {
	start = 0
	end = len(sa) - 1
	find := 0
	for start < end {
		if sa[start] > target || sa[end] < target {
			return -1, -1
		}
		find = (start + end) / 2
		if sa[find] > target {
			end = find - 1
		} else if sa[find] < target {
			start = find + 1
		} else {
			break
		}
	}
	//找到一个target后，对左边和右边再进行两次二分来查找等于target的左右边界
	start1 := 0
	end1 := find
	start2 := find
	end2 := len(sa)
	find1 := 0
	find2 := 0
	for start1 < end1 {

		find1 = (start1 + end1) / 2
		if sa[find1] < target {
			start1 = find1 + 1
		}
		if sa[find1] == target {
			end1 = find1 - 1
		}
	}
	//结束条件要想好
	if sa[start1] == target {
		start = start1
	} else {
		start = end1
	}

	for start2 < end2 {
		find2 = (start2 + end2) / 2
		if sa[find2] > target {
			end2 = find2 - 1
		}
		if sa[find2] == target {
			start2 = find2 + 1
		}
	}
	if sa[end2] == target {
		end = end2
	} else {
		end = start2
	}

	return start, end
}

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index
where it would be if it were inserted in order.
You may assume no duplicates in the array
Here are few examples.
[1,3,5,6], 5 → 2
[1,3,5,6], 2 → 1
[1,3,5,6], 7 → 4
[1,3,5,6], 0 → 0

**/
func SearchInsertPosition(sa []int, target int) int {
	start := 0
	end := len(sa) - 1
	find := 0
	for start < end {

		if sa[end] < target {
			return end + 1
		}
		if sa[start] > target {
			return start
		}
		find = (start + end) / 2
		if sa[find] > target {
			end = find - 1
		} else if sa[find] < target {
			start = find + 1
		} else {
			return find
		}
	}
	return find
}

/*
Write an efficient algorithm that searches for a value in an m × n matrix. This matrix has the following
properties:
• Integers in each row are sorted from left to right.
• The first integer of each row is greater than the last integer of the previous row
For example, Consider the following matrix:
[
[1, 3, 5, 7],
[10, 11, 16, 20],
[23, 30, 34, 50]
]
Given target = 3, return true.

**/
// func Searcha2DMatrix() {

// }
//[]int{3, 2, 5, 6, 6, 4, 3, 1, 7}
func quickSort(sa []int) {
	baseIndex := 0
	i := 0
	j := len(sa) - 1
	for i < j {
		for sa[j] > sa[baseIndex] && i < j {
			j--
		}
		for sa[i] <= sa[baseIndex] && i < j {
			i++
		}
		sa[i], sa[j] = sa[j], sa[i]
	}
	sa[baseIndex], sa[i] = sa[i], sa[baseIndex]
	if i-1 > 0 {
		quickSort(sa[0:i])
	}
	if i+1 < len(sa) {
		quickSort(sa[i+1:])
	}
	return
}

//一个数组中，有一个数只出现一次，其他数都出现了两次，找到这个数
func yihuo(sa []int) int {
	res := sa[0]
	for i := 1; i < len(sa); i++ {
		res = res ^ sa[i]
	}
	return res
}


