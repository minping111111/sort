package main

import (
	"errors"
	"fmt"
)

var source = []int{1, 5, 3, 8, 7, 9}

func main() {
	res := guibin(source)
	fmt.Println(res)
}

//data已经是从小到大排好序的，进行二分查找elem
func erfen(data []int, elem int) (index int, err error) {
	if data[len(data)-1] < elem || data[0] > elem {
		return -1, errors.New("not found")
	}
	middle := len(data) / 2
	if data[middle] > elem { //在前半部分
		index, err := erfen(data[:middle], elem)
		return index, err
	} else if data[middle] < elem { //在后半部分
		index, err := erfen(data[middle:], elem)
		return index, err
	} else {
		return middle, nil
	}
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

//从大到小冒泡
func maopaoR(source []int) []int {
	length := len(source)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if source[i] < source[j] {
				source[i], source[j] = source[j], source[i]
			}
		}
	}
	return source
}

//插入排序：将一个新数插入到已经排好序的队列中
func charu(source []int) []int {
	length := len(source)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if source[j] > source[i] {
				temp := source[i]
				for k := i - 1; k >= j; k-- {
					source[k+1] = source[k]
				}
				source[j] = temp
				break
			}
		}
	}
	return source
}

//选择排序：先选出最大数，然后选择第二大数，第三大数。。。
func xuanze(source []int) []int {
	length := len(source)
	for i := 0; i < length; i++ {
		index := 0

		for j := 0; j < length-i; j++ {
			if source[j] > source[index] {
				index = j
			}
		}
		source[length-i-1], source[index] = source[index], source[length-i-1]
	}
	return source
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

//归并排序，先将数组的1,2为一组A，3,4为一组B，5,6为一组C，7,8为一组D，先将A,B,C,D各自组内排好序，然后将A,B合并成AA，C,D合并CC，AA和CC分别组内排好序，然后将AA和CC合并，一起排好序。
func guibin(source []int) []int {
	length := len(source)
	for i := 2; i < length; i *= 2 {
		//好多子小组
		//0,1,2...i-1;      i,i+1,i+2...2i-1;          (length/i)*i,(length/i)*i+1,(length/i)*i+2...length-1
		for j := 0; j <= length/i; j++ {
			for k := j * i; k < (j+1)*i && k < length; k++ {
				for l := 0; l < k; l++ {
					if source[l] > source[k] {
						temp := source[k]
						for m := k - 1; m >= l; m-- {
							source[m+1] = source[m]
						}
						source[l] = temp
					}
				}
			}
		}
	}
	return source
}
//快排，先选最后一个元素为基准元素，从头遍历其他元素，与基准元素比较，大于基准元素的放前面，小于基准元素的放小于元素的后面，而不是基准元素的后面。最后将中间元素与基准元素交换。
//这样基准元素的左边都小于基准元素，右边都大于基准元素。左右两边各成一个子列。对每个子列再分别进行快排操作，递归下去即可。最终完成整个数组的排序。
func quickSort() []int {
	source := []int{1, 5, 3, 7, 9, 4, 2}
	realQuickSort(source)
	return source

}

func realQuickSort(source []int) {
	storeIndex := 0
	baseIndex := len(source) - 1
	if baseIndex <= storeIndex {
		return
	} else {
		for i := storeIndex; i < baseIndex; i++ {
			if source[i] < source[baseIndex] {
				source[storeIndex], source[i] = source[i], source[storeIndex]
				storeIndex++
			}
		}
		source[storeIndex], source[baseIndex] = source[baseIndex], source[storeIndex]
		realQuickSort(source[:storeIndex])
		realQuickSort(source[storeIndex+1:])
	}
}
