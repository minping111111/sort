package main

import (
	"errors"
	"fmt"
)

func main() {
	// res, length := RemoveDuplicate3([]int{1, 3, 9, 7, 9, 3, 5, 7, 6, 7})
	// fmt.Println(res)
	// fmt.Println(length)
	// source := []int{1, 3, 5, 7, 9}
	// res := erfen(source, 0, len(source)-1, 5)
	// fmt.Println(res)
	// res := SearchInRotatedSortedArray([]int{4, 5, 6, 6, 7, 0, 1, 2}, 6)
	res, err := GetKthElem3([]int{1, 3, 5, 7, 9}, []int{2, 3, 5, 6, 8}, 9)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

}

/*
给出一个排好序的数组，剔除掉重复的元素，并返回新数组的长度
要求：不能分配一个新的数组空间
比如，给出A=[1,1,2]，返回值为2，新数组为[1,2]
*/
func RemoveDuplicate(source []int) ([]int, int) {
	leng := len(source)
	for i := 1; i < leng; i++ {
		if source[i] == source[i-1] {
			for j := i; j < leng-1; j++ {

				source[j] = source[j+1]

			}
			leng--
			i--
		}
	}
	return source[:leng], leng
}

/*
给出一个排好序的数组，剔除掉出现超过2此的重复的元素，并返回新数组的长度
要求：不能分配一个新的数组空间
比如，给出A=[1,1,1,2,2,3]，返回值为4，新数组为[1,1,2,2,3]
*/
func RemoveDuplicate2(source []int) ([]int, int) {
	leng := len(source)
	for i := 1; i < leng-1; i++ {
		if source[i] == source[i-1] && source[i+1] == source[i-1] {
			for j := i + 1; j < leng-1; j++ {
				source[j] = source[j+1]
			}
			leng--
			i-- //注意这个i要在这一步for循环结束后加1，要保持i不变，这里要减一
		}
	}
	return source[:leng], leng
}

/*
给出一个没有好序的数组，剔除掉出现超过2此的重复的元素，并返回新数组的长度
思路：用hashmap来记录数组中元素出现的次数
比如，给出A=[1,2,2,1,1,3]，返回值为4，新数组为[1,2,2,1,3]
*/
func RemoveDuplicate3(source []int) ([]int, int) {
	leng := len(source)
	timesHash := make(map[int]int)
	for i := 0; i < leng; i++ {
		v, ok := timesHash[source[i]]
		if ok {
			if v == 2 { //去掉，并将后面元素前移
				for j := i; j < leng-1; j++ {
					source[j] = source[j+1]
				}
				i--
				leng--
			} else {
				timesHash[source[i]] += 1
			}
		} else {
			timesHash[source[i]] = 1
		}

	}
	return source[:leng], leng
}

// 二分法
func erfen(source []int, start, end, value int) int {
	if end-start < 0 {
		return -1
	}
	index := start + (end-start)/2
	if source[index] == value {
		return index
	} else if source[index] > value {
		return erfen(source, start, index-1, value)
	} else {
		if index < end {
			return erfen(source, index+1, end, value)
		} else {
			return -1
		}

	}
}

// 假定有一个排好序的不重复的序列A(0 1 2 4 5 6 7)，经过某个轴旋转，得到序列B(4 5 6 7 0 1 2),在这个序列B中查找一个数，如果存在，返回在B中的索引值，否则返回-1,
func SearchInRotatedSortedArray(source []int, value int) int {
	leng := len(source)
	if leng == 0 {
		return -1
	}
	if leng == 1 {
		if source[0] == value {
			return 0
		} else {
			return -1
		}
	}
	for i := 1; i < leng; i++ {
		if source[i]-source[i-1] < 0 { //找到轴元素
			if source[i] == value {
				return i
			} else if source[0] > value {
				return erfen(source, i, leng-1, value)
			} else {
				return erfen(source, 0, i-1, value)
			}

		}
	}
	return -1
}

// 找出两列排好序的数组的全部元素的第K大的元素
func GetKthElem(sourceA, sourceB []int, k int) (int, error) {
	//先对A，B自身去重
	sourceA, _ = RemoveDuplicate(sourceA)
	sourceB, _ = RemoveDuplicate(sourceB)
	source := []int{}
	lengB := len(sourceB)
	index := 0
LOOP1:
	for i := 0; i < lengB; i++ {
		for j := index; j < len(sourceA); j++ {
			if sourceA[j] == sourceB[i] { //如果B中元素在A中已经存在，忽略
				continue LOOP1
			}
			if sourceA[j] > sourceB[i] { //把sourceB[i]插入到排好序的sourceA中
				temp := sourceB[i]
				sourceA = append(sourceA, 0) //先暂时给个领
				for k := len(sourceA) - 2; k >= j; k-- {
					sourceA[k+1] = sourceA[k]
				}
				sourceA[j] = temp
				break
			}
		}
	}
	// //合并后的整个数组
	source = append(source, sourceA...)
	if len(source) >= k {
		return source[k-1], nil
	}
	return -1, errors.New(fmt.Sprintf(`the %dth older elem do not found`, k))

}

// 找出两列排好序的数组的全部元素的第K小的元素,第二种方法
func GetKthElem2(sourceA, sourceB []int, k int) (int, error) {
	sourceA, _ = RemoveDuplicate(sourceA)
	sourceB, _ = RemoveDuplicate(sourceB)
	lengA := len(sourceA)
	lengB := len(sourceB)
	indexA := 0
	indexB := 0
	index := 0
	elem := 0

	for index < k {
		if indexA <= lengA-1 && indexB <= lengB-1 {
			if sourceA[indexA] > sourceB[indexB] {
				elem = sourceB[indexB]
				indexB++
			} else if sourceA[indexA] < sourceB[indexB] {
				elem = sourceA[indexA]
				indexA++
			} else {
				elem = sourceA[indexA]
				indexA++
				indexB++
			}
		} else if indexA > lengA-1 && indexB <= lengB-1 {
			elem = sourceB[indexB]
			indexB++
		} else if indexB > lengB-1 && indexA <= lengA-1 {
			elem = sourceA[indexA]
			indexA++
		} else {
			return -1, errors.New(fmt.Sprintf(`the %dth older elem do not found`, k))
		}
		index++
	}

	return elem, nil

}

//第三种方法。先去除A数组中的第k/2个元素和B数组中的第K/2个元素，然后再按照第二种方法继续寻找
func GetKthElem3(sourceA, sourceB []int, k int) (int, error) {
	sourceA, _ = RemoveDuplicate(sourceA)
	sourceB, _ = RemoveDuplicate(sourceB)
	delmap := make(map[int]bool)
	lengA := len(sourceA)
	lengB := len(sourceB)
	delindex := 0
	elem := -1
	i := 0
	j := 0
	for i = 0; i < k/2 && i < lengA; i++ {
		delmap[sourceA[i]] = true
		delindex++
	}
	for j = 0; j < k/2 && j < lengB; j++ {
		if _, ok := delmap[sourceB[j]]; !ok {
			delmap[sourceB[j]] = true
			delindex++
		}
	}

	for delindex < k {
		if i < lengA && j < lengB {
			if sourceA[i] < sourceB[j] {
				if _, ok := delmap[sourceA[i]]; !ok {
					elem = sourceA[i]
					i++
					delindex++
				} else {
					i++
				}
			} else if sourceA[i] > sourceB[j] {
				if _, ok := delmap[sourceB[j]]; !ok {
					elem = sourceB[j]
					j++
					delindex++
				} else {
					j++
				}
			} else {
				elem = sourceA[i]
				i++
				j++
				delindex++
			}
		} else if i < lengA && j >= lengB { //B遍历完了
			if _, ok := delmap[sourceA[i]]; !ok {
				elem = sourceA[i]
				i++
				delindex++
			} else {
				i++
			}
		} else if j < lengB && i >= lengA { //A遍历完了
			if _, ok := delmap[sourceB[j]]; !ok {
				elem = sourceB[j]
				j++
				delindex++
			} else {
				j++
			}
		} else {
			return -1, errors.New(fmt.Sprintf(`the %dth older elem do not found`, k))
		}
	}
	return elem, nil

}
