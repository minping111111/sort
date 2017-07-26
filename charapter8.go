package main

import (
	"fmt"
)

func main() {
	res := SubsetsII([]int{1, 2, 2})
	fmt.Println(res)
}

/*
Given a set of distinct integers, S, return all possible subsets.
Note:
• Elements in a subset must be in non-descending order.
• The solution set must not contain duplicate subsets.
For example, If S = [1,2,3], a solution is:
[
[3],
[1],
[2],
[1,2,3],
[1,3],
[2,3],
[1,2],
[]
]
**/
func Subsets(sa []int) [][]int {
	subs := [][]int{[]int{}}
	for i := 0; i < len(sa); i++ {
		for _, v := range subs {
			subs = append(subs, append(v, sa[i]))
		}
	}
	return subs
}

/*
Given a collection of integers that might contain duplicates, S, return all possible subsets.
Note:
Elements in a subset must be in non-descending order. The solution set must not contain duplicate
subsets. For example, If S = [1,2,2], a solution is:
[
[2],
[1],
[1,2,2],
[2,2],
[1,2],
[]
]
**/
func SubsetsII(sa []int) [][]int {
	subs := [][]int{[]int{}}
	flag := 0
	for i := 0; i < len(sa); i++ {
		flag++
		if flag <= len(sa) && len(subs) > 1 {
			if sa[i] == subs[len(subs)-1][0] {
				continue
			}
		}
		for _, v := range subs {
			subs = append(subs, append(v, sa[i]))
		}
	}
	return subs
}

/*
Given a collection of numbers, return all possible permutations.
For example, [1,2,3] have the following permutations: [1,2,3], [1,3,2], [2,1,3], [2,3,1],
[3,1,2], and [3,2,1].
**/
//找一个排列的下一个排列的方法：
/*
http://www.cnblogs.com/easonliu/p/3632442.html
第一步：从后到前，找到相邻的两个数i和i+1,如果sa[i]< sa[i+1],则进行第二步，如果没找到，则结束
第二步：交换sa[i]和sa[i+1]的值，从头开始走第一步
**/
func Permutations(sa []int) [][]int {

}
