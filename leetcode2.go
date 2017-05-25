package main

import (
	"errors"
	"fmt"
)

func main() {
	// res, leng := searchConsecutiveSequence([]int{100, 4, 200, 3, 1})
	// fmt.Println(res)
	// fmt.Println(leng)
	res := searchThreeSum2([]int{1, 8, 3, 6, 5, 12, 7, 9}, 9)
	fmt.Println(res)
}

//求最长的子连续序列，比如输入为100, 4, 200, 1, 3, 2，输出为1,2,3,4,要求在O(n)复杂度内完成
//o(n)所以不能进行排序
//思路，先将数组映射成一个map，然后以map的key为中心，向左拓展，向右拓展，以争取到最大
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

//给定序列，找出其中两个数之和是指定数
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

//给定序列，找出其中两个数之和是指定数,注意这个map建立的巧妙之处，key是sum-val，值是下标，因为map只是按key查找的开销才是o(1)
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

//希尔排序，先等间隔分组，然后在每组内进行插入排序，每组排好后，再减小间隔再分组，分组内插入排序，直到间隔<1为止
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
		partRes = jiabi(partSource, sum-source[i])
		if len(partRes) != 0 {
			for _, v := range partRes {
				v = append(v, source[i])
				res = append(res, v)
			}
		}
	}
	return res
}

//夹逼定理来寻找数对
func jiabi(source []int, sum int) [][]int {

}
