package main

import (
	"fmt"
)

func main() {
	// res := SubsetsII([]int{1, 2, 2})
	// res := Permutations([]int{1, 2, 3})
	res := Combinations(0, 20, 6)
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
	isCircle := true
	permu := [][]int{sa}

	for isCircle {
		isCircle = false
		for i := len(sa) - 1; i > 0; i-- {
			if sa[i-1] < sa[i] {
				sa[i-1], sa[i] = sa[i], sa[i-1]
				//关于slice的引用类型，这里比较坑。。。
				temp := make([]int, len(sa))
				copy(temp, sa)
				permu = append(permu, temp)
				isCircle = true
				break
			}
		}
	}
	return permu
}

/*
Given two integers n and k, return all possible combinations of k numbers out of 1...n.
For example, If n = 4 and k = 2, a solution is:
[
[2,4],
[3,4],
[2,3],
[1,2],
[1,3],
[1,4],
]
**/
func Combinations(start, end, k int) [][]int {
	res := [][]int{}
	if end-start < k {
		return res
	}
	if k == 1 {
		tempRes := [][]int{}
		for j := start; j < end; j++ {
			tempRes = append(tempRes, []int{j + 1})
		}
		return tempRes
	}
	for i := start; i <= end-k; i++ {
		clildRes := Combinations(i+1, end, k-1)
		for _, v := range clildRes {
			v1 := append([]int{i + 1}, v...)
			res = append(res, v1)
		}
	}

	return res

}

var phoneTable = map[string][]string{
	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	5: []string{"m", "n", "o"},
	6: []string{"p", "q", "r", "s"},
	7: []string{"t", "u", "v"},
	8: []string{"w", "x", "y", "z"},
}

/*
Given a digit string, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below.
Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note: Although the above answer is in lexicographical order, your answer could be in any order you
want
**/
func LetterCombinationsofaPhoneNumber(selNum []int) []string {

	res := []string{}
	if len(selNum) == 1 {
		for j := 0; j < len(phoneTable[selNum[0]]); j++ {
			res = append(res, []string{phoneTable[selNum[0]][j]})
		}
		return res
	}
	for i := 0; i < len(phoneTable[selNum[0]]); i++ {
		res = LetterCombinationsofaPhoneNumber()
	}

}
