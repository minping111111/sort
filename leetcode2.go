package main

import (
	"errors"
	"fmt"
	// "sort"
	// "math"
)

func main() {
	// res, leng := searchConsecutiveSequence([]int{100, 4, 200, 3, 1})
	// fmt.Println(res)
	// fmt.Println(leng)

	// res := removeElem2([]int{1, 8, 3, 6, 3, 5, 4, 12, 7, 9}, 3)
	// fmt.Println(res)
	// res := nextPaiLieZuHe([]int{7, 3, 1, 2, 2, 5})
	// fmt.Println(res)
	res := PermutationSequence(5, 3)
	fmt.Println(res)
}

// 求最长的子连续序列，比如输入为100, 4, 200, 1, 3, 2，输出为1,2,3,4,要求在O(n)复杂度内完成
// o(n)所以不能进行排序
// 思路，先将数组映射成一个map，然后以map的key为中心，向左拓展，向右拓展，以争取到最大
func searchConsecutiveSequence(source []int) ([]int, int) {

	leng := len(source)
	hashTab := make(map[int]bool)
	for i := 0; i < leng; i++ {
		hashTab[source[i]] = true
	}

	index_left := 0
	index_right := 0
	var i int
	var j int
	maxConsecutiveLen := 0
	for k, _ := range hashTab {
		tempConsecutiveLen := 1
		//向左拓展
		for i = k - 1; i > -9999999; i-- {
			_, ok := hashTab[i]
			if ok {
				tempConsecutiveLen++
			} else {
				break
			}
		}

		//向右拓展
		for j = k + 1; j < 99999999; j++ {
			_, ok := hashTab[j]
			if ok {
				tempConsecutiveLen++
			} else {
				break
			}
		}
		if tempConsecutiveLen > maxConsecutiveLen {
			index_left = i + 1
			index_right = j - 1
		}
	}

	new_source := []int{}
	for m := index_left; m <= index_right; m++ {
		new_source = append(new_source, m)
	}
	return new_source, len(new_source)

}

// 给定序列，找出其中两个数之和是指定数
func searchTwoSum(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	if leng <= 1 {
		return res
	}
	hashmap := make(map[int]int)
	for i := 0; i < leng; i++ {
		hashmap[source[i]] = sum - source[i]
	}

	for j := 0; j < leng; j++ {
		index, err := search(source, hashmap[source[j]])
		if err == nil {
			res = append(res, []int{j, index})
		}
	}
	return res
}

func search(source []int, num int) (int, error) {
	leng := len(source)
	for i := 0; i < leng; i++ {
		if source[i] == num {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

// 给定序列，找出其中两个数之和是指定数,注意这个map建立的巧妙之处，key是sum-val，值是下标，因为map只是按key查找的开销才是o(1)
func searchTwoSum2(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	if leng <= 1 {
		return res
	}
	hashmap := make(map[int]int)
	for i := 0; i < leng; i++ {
		hashmap[sum-source[i]] = i
	}

	for j := 0; j < leng; j++ {
		v, ok := hashmap[source[j]]
		if ok && v != j { //排除3+3=6的这种情况
			res = append(res, []int{source[j], source[v]})
		}
	}
	return res
}

func searchThreeSum(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	for i := 0; i < leng; i++ {
		partSource := []int{}
		if (i - 1) >= 0 {
			partSource = append(partSource, source[:i]...)
		}
		if i+1 <= leng-1 {
			partSource = append(partSource, source[i+1:]...)
		}
		//注意，这里的partSource是新的数组，所以下标换了，后面的返回值不能用下标表示，而要用值来表示，如果非要用下标表示，再遍历一遍，换成下标即可
		partRes := searchTwoSum2(partSource, sum-source[i])
		if len(partRes) != 0 {
			for _, v := range partRes {
				v = append(v, source[i])
				res = append(res, v)
			}
		}
	}
	return res
}

// 希尔排序，先等间隔分组，然后在每组内进行插入排序，每组排好后，再减小间隔再分组，分组内插入排序，直到间隔<1为止
func xier(source []int) []int {
	length := len(source)
	for gap := length / 2; gap >= 1; gap = gap / 2 {
		//0,0+gap;0+2gap;0+3gap...
		for i := 1; i*gap < length; i++ {
			for j := 0; j < i; j++ {
				if source[j*gap] > source[i*gap] {
					temp := source[i*gap]
					for k := i - 1; k >= j; k-- {
						source[(k+1)*gap] = source[k*gap]
					}
					source[j*gap] = temp
				}
			}
		}
	}
	return source
}

func searchThreeSum2(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	source = xier(source)
	for i := 0; i < leng; i++ {
		partSource := []int{}
		if (i - 1) >= 0 {
			partSource = append(partSource, source[:i]...)
		}
		if i+1 <= leng-1 {
			partSource = append(partSource, source[i+1:]...)
		}
		partRes := jiabi(partSource, sum-source[i])
		if len(partRes) != 0 {
			for _, v := range partRes {
				v = append(v, source[i])
				res = append(res, v)
			}
		}
	}
	return res
}

// 对排好序的使用夹逼定理来寻找数对,第一个和最后一个，判断他们的和，如果大了，则大数往左移，如果小了，则小数往右移动，如果刚好相等，则拿到了一对，然后小数右移大数左移，一步一步向中间夹，复杂度感觉可以到O(N)
func jiabi(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	i := 0
	j := leng - 1
	for i < j && j > 0 {
		if source[i]+source[j] < sum { //i要往前移
			i++
		} else if source[i]+source[j] == sum {
			res = append(res, []int{source[i], source[j]})
			i++
			j--
		} else {
			j--
		}
	}
	return res
}

// // 给定一个数sum，查找数组source中的三个数，使他们的和最接近sum
// func searchThreeSumClosest(source []int, sum int) []int {
// 	res := [][]int{}
// 	leng := len(source)
// 	if leng < 3 {
// 		return [][]int{}
// 	}
// 	source = xier(source)
// 	closerSource := source[0] + source[1] + source[2]
// 	num1 := []int{source[0], source[1], source[2]}
// 	for i := 0; i < leng; i++ {
// 		partSource := []int{}
// 		if (i - 1) >= 0 {
// 			partSource = append(partSource, source[:i]...)
// 		}
// 		if i+1 <= leng-1 {
// 			partSource = append(partSource, source[i+1:]...)
// 		}
// 		partRes := jiabiCloser(partSource, sum-source[i])
// 		if len(partRes) != 0 {
// 			for _, v := range partRes {
// 				v = append(v, source[i])
// 				res = append(res, v)
// 			}
// 		}
// 	}
// 	for _, v := range res {
// 		tempsum := res[0][0] + res[0][1] + res[0][2]
// 		for _, v2 := range v {
// 			tempsum += v2
// 		}
// 		if math.Abs(tempsum-sum) < math.Abs(closerSource-sum) {
// 			closerSource = tempsum
// 			num1 = []int{v[0], v[1], v[2]}
// 		}
// 	}
// 	return num1
// }

// //夹逼找出两个数的和最接近sum的两个数
// func jiabiCloser(source []int, sum int) [][]int {

// }

func searchThreeSumHash(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	hash := make(map[int][]int)
	for i := 0; i < leng; i++ {
		for j := i + 1; j < leng; j++ {
			inex := sum - source[i] - source[j]
			hash[inex] = []int{i, j}
		}
	}
LOOP:
	for m := 0; m < leng; m++ {
		k, v := hash[source[m]]
		if v {
			for _, v1 := range k {
				if m == v1 {
					continue LOOP
				}
			}
			hash[source[m]] = append(hash[source[m]], m)
		}
	}

	for _, v2 := range hash {
		if len(v2) == 3 {
			res = append(res, v2)
		}
	}
	return res
}

//四个,先用hasha来存一个，key为sum-两个元素之和，val为两个下标，然后用hashb来存一个，key为两个数之和，val为两个下标，然后对比两个hash是否有相同的key，有相同的key的情况下，看下标是否有重复，也是是否是重复用了某个元素
func searchFourSumHash(source []int, sum int) [][]int {
	res := [][]int{}
	leng := len(source)
	hasha := make(map[int][]int)
	hashb := make(map[int][]int)
	for i := 0; i < leng; i++ {
		for j := i + 1; j < leng; j++ {
			inex := sum - source[i] - source[j]
			hasha[inex] = []int{i, j}
		}
	}

	for m := 0; m < leng; m++ {
		for n := m + 1; n < leng; n++ {
			inex := source[m] + source[n]
			hashb[inex] = []int{m, n}
		}
	}

	for k1, v1 := range hasha {
		for k2, v2 := range hashb {
			if k1 == k2 {
				if hasCle := checkCle(v1, v2); hasCle {
					res = append(res, append(v1, v2...))
				}
			}
		}
	}
	return res
}

//判断两个数组是否有交集
func checkCle(source1, source2 []int) bool {
	for i := 0; i < len(source1); i++ {
		for j := 0; j < len(source2); j++ {
			if source2[j] == source1[i] {
				return false
			}
		}
	}
	return true
}

// 从数组中移除所有等于某个值的元素,返回新数组的长度

// 时间复杂度0(n),空间复杂度o(n)，利用了新数组
func removeElem(source []int, elem int) int {
	newSource := []int{}
	for i := 0; i < len(source); i++ {
		if source[i] != elem {
			newSource = append(newSource, source[i])
		}
	}
	return len(newSource)
}

// 时间复杂度0(n),空间复杂度o(1),发现其实不需要利用新数组，直接在原数组中赋值即可，只要保证i走在index的前面就行
func removeElem2(source []int, elem int) int {
	index := 0
	for i := 0; i < len(source); i++ {
		if source[i] != elem {
			source[index] = source[i]
			index++
		}
	}
	return len(source[:index])
}

//输入一个序列，输出这个序列的排列组合的下一个序列
/*比如“1，2，3”的全排列，依次是：
1 2 3
1 3 2
2 1 3
2 3 1
3 1 2
3 2 1
*/
/* 所以
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPaiLieZuHe(source []int) []int {
	leng := len(source)
	everflag := false
	for k := leng - 2; k >= 0; k-- {
		if source[k] < source[k+1] {
			everflag = true
			abstu := 0
			flag := false
			index := k + 1
			for j := k + 1; j < leng; j++ {
				if source[j] > source[k] {
					if !flag {
						abstu = source[j] - source[k]
						index = j
						flag = true
					} else {
						if source[j]-source[k] < abstu {
							abstu = source[j] - source[k]
							index = j
						}
					}
				}
			}
			source[k], source[index] = source[index], source[k]
			//对 k+1,k+2...进行从大到小排序
			source = append(source[:k+1], maopao(source[k+1:])...)
			break
		}

	}
	if !everflag { //如果已经到末尾了，返回第一个排序数
		source = maopao(source)
	}
	return source
}

//从小到大冒泡
func maopao(source []int) []int {
	length := len(source)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if source[i] > source[j] {
				source[i], source[j] = source[j], source[i]
			}
		}
	}
	return source
}

/*
The set [1,2,3,...,n] contains a total of n! unique permutations.
By listing and labeling all of the permutations in order, We get the following sequence (ie, for n = 3):
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.
Note: Given n will be between 1 and 9 inclusive.
**/
func PermutationSequence(n, k int) []int {
	source := []int{}
	for i := 1; i <= n; i++ {
		source = append(source, i)
	}
	for j := 1; j < k; j++ {
		source = nextPaiLieZuHe(source)
	}
	return source
}

//k>=1，求阶乘
func jiecheng(k int) int {
	res := 1
	for i := 1; i <= k; i++ {
		res = res * i
	}
	return res
}

//todo
func PermutationSequence2(n, k int) []int {
	// source := []int{}
	// var newSource [n]int

	// for i := 0; i <= n; i++ {
	// 	source = append(source, i)
	// }
	// tempSource := source
	// for j := 1; j < len(tempSource); j++ {
	// 	newSource[j] =  k/jiecheng(len(tempSource)-1-j) +1
	// 	tempSource = append(tempSource[:newSource[j]+1],tempSource[newSource[j]+1:])
	// }

	// return source
}
