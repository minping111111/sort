package main

import (
	"fmt"
	"strconv"
)

var validstr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	// res := implementStr([]string{"b", "c", "b", "c", "d", "a"}, []string{"b", "c", "d"})
	// fmt.Println(res)
	// res := stringtoInteger([]string{"-", "1", "b", "2", "a"})
	// fmt.Println(res)
	// res := addBinary([]int{1, 0, 1}, []int{1})
	// res := longestPalindromicSubstringDP([]string{"e", "b", "c", "b", "a", "b", "c", "d"}) //最长的是cbabc
	res := regularExpressionMatching([]string{"a", "a", "a"}, []string{"a", "*", "b"})
	fmt.Println(res)
}

func implementStr(source, strsou []string) int {
	soulen := len(source)
	strsoulen := len(strsou)
	if soulen < strsoulen {
		return -1
	}
	// i := 0
LOOP:
	for i := 0; i <= soulen-strsoulen; i++ {
		for j := 0; j < strsoulen; j++ {
			if source[i+j] != strsou[j] {
				continue LOOP

			}
		}
		return i
	}
	return -1
}

//字符串转整形
func stringtoInteger(source []string) int {
	//validstr
	leng := len(source)
	if leng < 1 {
		return 0
	}
	baseindex := 0
	var i int
	if source[0] == "+" || source[0] == "-" {
		for i = 1; i < leng; i++ {
			if !in_array(source[i], validstr) {
				if i > 1 {
					break
				} else {
					return 0
				}
			}
		}
		baseindex = i
	} else if in_array(source[0], validstr) {
		for i := 0; i < leng; i++ {
			if !in_array(source[i], validstr) {
				if i > 1 {
					break
				} else {
					return 0
				}
			}
		}
		baseindex = i
	} else {
		return 0
	}
	newslice := source[:baseindex]
	var newstr string
	for j := 0; j < len(newslice); j++ {
		newstr = newstr + newslice[j]
	}
	res, err := strconv.Atoi(newstr)
	if err == nil {
		return res
	} else {
		return 0
	}
}

func in_array(key string, array []string) bool {
	is_valid := false
	for i := 0; i < len(array); i++ {
		if key == array[i] {
			is_valid = true
			break
		}
	}
	return is_valid
}

/*
a = "11"
b = "1"
Return "100"
*/
func addBinary(source1, source2 []int) []int {
	leng1 := len(source1)
	leng2 := len(source2)
	newSource := []int{}
	jin := 0
	if leng1 > leng2 {
		for i := leng1 - 1; i >= 0; i-- {
			var tempSum int
			if leng2-1-(leng1-1-i) >= 0 {
				tempSum = source1[i] + source2[leng2-1-(leng1-1-i)] + jin
			} else {
				tempSum = source1[i] + jin
			}
			if tempSum > 1 {
				jin = 1
			} else {
				jin = 0
			}
			tempSum = tempSum % 2
			newSource = append(newSource, tempSum)
		}

	} else {
		for i := leng2 - 1; i >= 0; i-- {
			var tempSum int
			if leng1-1-(leng2-1-i) >= 0 {
				tempSum = source1[i] + source2[leng1-1-(leng2-1-i)] + jin
			} else {
				tempSum = source1[i] + jin
			}
			if tempSum > 1 {
				jin = 1
			} else {
				jin = 0
			}
			tempSum = tempSum % 2
			newSource = append(newSource, tempSum)
		}
	}
	if jin == 1 {
		newSource = append(newSource, 1)
	}
	return reverse(newSource)

}

func reverse(source []int) []int {
	leng := len(source)
	for i := 0; i <= leng/2; i++ {
		source[i], source[leng-1-i] = source[leng-1-i], source[i]
	}
	return source
}

//返回最长回文子串 ebcbabcd最长的是cbabc，自己写的很欣慰
func longestPalindromicSubstring(source []string) []string {
	huiwen := [][]string{}
	leng := len(source)
	if leng <= 1 {
		return source
	}
	for i := 0; i < leng; i++ {
		k := i
		for j := leng - 1; j > k; j-- {
			v := j
			toend := true
			for v > k {
				if source[v] == source[k] {
					v--
					k++
				} else {
					toend = false
					break
				}
			}
			if toend { //合法的回文字符串都放进来，后面再选取最长的
				huiwen = append(huiwen, source[i:j+1])
			}
			k = i
		}
	}
	maxlen := len(huiwen[0])
	result := huiwen[0]
	for _, v := range huiwen {
		if len(v) > maxlen {
			maxlen = len(v)
			result = v
		}
	}
	return result
}

type d map[int]bool

// type DD map[int]d

//返回最长回文子串 ebcbabcd最长的是cbabc，使用DP
/*
D[i][j]为true表示source[i:j+1]是回文字符串,特殊的，本来是想用DP的，但是写着写着，就没用用DP了，因为回文这个规律其实很规整，左一个，右一个的，其实很好判断，
自己写的还是很高兴，这个方法用到了对称性，O(N*2),比上面一个方法的0(N*3)要好点
*/
func longestPalindromicSubstringDP(source []string) []string {
	leng := len(source)
	if leng <= 1 {
		return source
	}
	D := make(d)
	for i := 0; i < leng; i++ {
		D[i] = true
	}

	result := [][]string{}

	for k, _ := range D {
		var i int
		for i = 1; k-i >= 0 && k+i < leng; i++ {
			if source[k-i] != source[k+i] {
				break
			}
		}
		result = append(result, source[k-i+1:k+i])
	}

	maxlen := 0
	res := []string{}
	for _, v := range result {
		if len(v) > maxlen {
			maxlen = len(v)
			res = v
		}
	}
	return res
}

//正则匹配
/*
Implement regular expression matching with support for '.' and '*'.
'.' Matches any single character. '*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
The function prototype should be:
bool isMatch(const char *s, const char *p)
Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true
**/
func regularExpressionMatching(source, match []string) bool {

	leng1 := len(source)
	leng2 := len(match)
	for i := 0; i < leng2; i++ {
		if match[i] == "*" {
			return true
		}
	}
	j := 0
	i := 0
LOOP:
	for i = 0; i < leng1; i++ {
		if j+i > leng2-1 {
			return false
		}
		if !(source[i] == match[j+i] || match[j+i] == ".") {
			j++
			goto LOOP
		}
	}
	if i == leng1 {
		return true
	}
	return false
}
