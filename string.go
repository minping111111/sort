package main

import (
	"fmt"
	"strconv"
	"strings"
)

var validstr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	// res := implementStr([]string{"b", "c", "b", "c", "d", "a"}, []string{"b", "c", "d"})
	// fmt.Println(res)
	// res := stringtoInteger([]string{"-", "1", "b", "2", "a"})
	// fmt.Println(res)
	// res := addBinary([]int{1, 0, 1}, []int{1})
	// res := longestPalindromicSubstringDP([]string{"e", "b", "c", "b", "a", "b", "c", "d"}) //最长的是cbabc
	// res := wildcardMatching([]string{"a", "a", "a"}, []string{"a", "*", "b"})
	// res := longestCommonPrefix([]string{"abc", "ab", "abcd"})
	// res := ValidNumber("2.1.0e3")
	// res := Anagrams([]string{"tea", "and", "ate", "eat", "den", "nad"})
	// res := SimplifyPath("/a/./b/../c/af/bafd/")
	// fmt.Println(res)
	res := lengthofLastWord("Hello World")
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

/*
Implement wildcard pattern matching with support for '?' and '*'.
'?' Matches any single character. '*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).
The function prototype should be:
bool isMatch(const char *s, const char *p)
Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "c*a*b") → false

// **/
// func wildcardMatching(source, match []string) bool {
// 	return true
// }

/*
求最长公共前缀
[]string{"abc", "ab", "a"} --> []string{"a"}
这道题还联系了string到[]byte的转化
**/
func longestCommonPrefix(source []string) string {
	leng := len(source)
	if leng == 0 {
		return ""
	}
	maytotal := []byte(source[0])
	var m int
	for k, v := range maytotal {
		m = k
		for i := 1; i < leng; i++ {
			if k >= len([]byte(source[i])) || []byte(source[i])[k] != v {
				goto END
			}
		}
	}
END:
	resByte := maytotal[:m]
	return string(resByte)
}

/*
Validate if a given string is numeric.
Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one
**/
func ValidNumber(source string) bool {
	sourceArr := []byte(source)
	leng := len(sourceArr)
	if leng < 1 {
		return false
	}
	first := string(sourceArr[0])
	numchar := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	firstchar := []string{"+", "-"}
	midchar := []string{"E", "e"}
	if !in_array(first, append(numchar, firstchar...)) {
		return false
	}
	eNums := 0
	dianNums := 0
	for i := 1; i < leng; i++ {
		if in_array(string(sourceArr[i]), numchar) {
			continue
		} else if in_array(string(sourceArr[i]), midchar) {
			eNums += 1
			if eNums > 1 {
				return false
			}
		} else if string(sourceArr[i]) == "." {
			dianNums += 1
			if dianNums > 1 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

/*
Given an integer, convert it to a roman numeral.
Input is guaranteed to be within the range from 1 to 3999
【罗马数字】
1~9: {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};
10~90: {"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
100~900: {"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
1000~3000: {"M", "MM", "MMM"}.
**/
func integertoRoman(input int) string {
	if input > 3000 || input < 1 {
		return "invalid"
	}
	qianMap := []string{}
	qianMap = append(qianMap, "M", "MM", "MMM")
	baiMap := []string{}
	baiMap = append(baiMap, "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM")
	shiMap := []string{}
	shiMap = append(shiMap, "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC")
	geMap := []string{}
	geMap = append(geMap, "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX")
	fmt.Println(geMap)
	qian := qianMap[input/1000-1]
	input = input % 1000
	bai := baiMap[input/100-1]
	input = input % 100
	shi := shiMap[input/10-1]
	input = input % 10
	ge := geMap[input%10-1]
	return qian + bai + shi + ge
}

func RomantoInteger(srt string) int {
	return 0
}

/*
The count-and-say sequence is the sequence of integers beginning as follows:
1, 11, 21, 1211, 111221, ...
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2", then "one 1" or 1211.
Given an integer n, generate the nth sequence.
Note: The sequence of integers will be represented as a string.
**/
func CountAndSay() []string {
	total := []string{"1"}
	for len(total) < 10 {
		base := total[len(total)-1]
		newstr := createSequence(base)
		total = append(total, newstr)
	}
	return total
}

/*
1211--〉111221
做这个题目的时候浪费了好多时间，原因在当时走到i的时候，用source[i]去比较source[i-1]，来决定要不要记录source[i-1]这个元素，这样很绕。
为什么在当前元素i下用source[i]去比较source[i+1]来判断要不要记录source[i]呢，这样就直观很多了是不是
**/
func createSequence(str string) (res string) {

	source := strings.Split(str, "")
	// res = res + strconv.Itoa(1) + source[0]
	leng := len(source)
	if leng == 0 {
		return ``
	}
	if leng == 1 {
		return `1` + source[0]
	}

	same := 1
	for i := 0; i < len(source)-1; i++ {

		if source[i] != source[i+1] {
			res = res + strconv.Itoa(same) + source[i]
			same = 1
		} else {
			same += 1
		}

	}
	if source[leng-1] != source[leng-2] {
		same = 1
	}
	res = res + strconv.Itoa(same) + source[leng-1]
	return
}

/*
回文构词法
Given an array of strings, return all groups of strings that are anagrams.
Note: All inputs will be in lower-case.
For example:
Input:　　["tea","and","ate","eat","den"]
Output:   ["tea","ate","eat"]
**/
func Anagrams(source []string) [][]string {
	var res [][]string
	hashMap := make(map[string][]int)
	for i := 0; i < len(source); i++ {
		tempstr := sortStr(source[i])
		_, ok := hashMap[tempstr]
		if ok {
			hashMap[tempstr] = append(hashMap[tempstr], i)
		} else {
			hashMap[tempstr] = []int{i}
		}
	}
	for _, v := range hashMap {
		if len(v) > 1 {
			var tempres []string
			for i := 0; i < len(v); i++ {
				tempres = append(tempres, source[v[i]])
			}
			res = append(res, tempres)
		}
	}
	return res
}

//冒泡法对字符串进行排序
func sortStr(str string) string {
	source := strings.Split(str, "")
	length := len(source)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if source[i] > source[j] {
				source[i], source[j] = source[j], source[i]
			}
		}
	}
	var res string
	for _, v := range source {
		res = res + v
	}
	return res
}

/*
简化路径表示法
Given an absolute path for a file (Unix-style), simplify it.
For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
Corner Cases:
• Did you consider the case where path = "/../"? In this case, you should return "/".
• Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
In this case, you should ignore redundant slashes and return "/home/foo".
**/
func SimplifyPath(str string) string {
	source := strings.Split(str, "")

	//先去除./中的. 将/a/./b/../../c/去除多余的"."转换成/a//b/../../c/
	for i := 0; i < len(source); i++ {
		if source[i] == "." {
			if source[i+1] == "/" && (i-1 < 0 || source[i-1] != ".") {
				source = append(source[:i], source[i+1:]...)
			}
		}
	}
	//将/a//b/../../c/去除多余的"/"变成/a/b/../b/a/../c/
	for i := 0; i < len(source)-1; i++ {
		if source[i] == "/" {
			if source[i+1] == "/" {
				source = append(source[:i], source[i+1:]...)
			}
		}
	}
	// //5是因为有前面的/a/..这四个字符
	var i int
	for i = 5; i < len(source); i++ {
		if source[i-2] == "." && source[i-1] == "." && source[i] == "/" {
			source = append(source[:i-4], source[i+1:]...)
		}
	}
	return strings.Join(source, "")
}

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length
of last word in the string.
If the last word does not exist, return 0.
Note: A word is defined as a character sequence consists of non-space characters only.
For example, Given s = "Hello World", return 5.
**/
func lengthofLastWord(str string) int {
	str_arr := strings.Split(str, " ")
	leng := len(str_arr)
	if leng < 0 {
		return 0
	}
	elem := str_arr[leng-1]
	elem_arr := strings.Split(elem, "")
	return len(elem_arr)
}
