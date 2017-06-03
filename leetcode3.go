package main

import (
	// "errors"
	"fmt"
	// "sort"
	"math"
	// "unsafe"
	// "strconv"
)

func main() {

	// res := RotateImageClockwise([][]int{[]int{1, 2}, []int{3, 4}})
	// fmt.Println(res)

	// res := plusOne([]int{8, 9, 9, 3, 9})
	// fmt.Println(res)
	// res := climbingStairs(5)
	// fmt.Println(res)
	// res := GrayCode(3)
	// fmt.Println(res)
	// res := gasStation([]int{1, 3, 5, 7, 9}, []int{4, 4, 5, 9, 6})
	// fmt.Println(res)
	res := candy([]float64{1.1, 2.3, 2.1, 5.3, 9.2, 4.5, 3.6})
	fmt.Println(res)
}

//将n*n矩阵顺时针旋转90度
func RotateImageClockwise(source [][]int) [][]int {
	leng := len(source)
	//先沿上下中轴反转
	for i := 0; i < leng/2; i++ {
		for j := 0; j < leng; j++ {
			source[i][j], source[leng-1-i][j] = source[leng-1-i][j], source[i][j]
		}
	}
	//再沿主对角线反转(i,j->j,i),注意是主对角线反转，如果用副对角线反转就是i,j->n-1-j,n-1-i
	for i := 0; i < leng; i++ {
		for j := 0; j < leng && j < i; j++ {
			source[i][j], source[j][i] = source[j][i], source[i][j]
		}
	}
	return source
}

//给定一个数组代表一个很大的数，给这个数加1，返回代表这个新数的新数组
func plusOne(source []int) []int {
	leng := len(source)
	newArray := []int{}
	isOlder := true
	for i := leng - 1; i >= 0; i-- {
		ind := source[i]
		if isOlder {
			ind += 1
		}
		if ind < 10 {
			isOlder = false
			newArray = append(newArray, ind)
		} else {
			isOlder = true
			newArray = append(newArray, 0)
			//当原本是9999时，加了1需要再加一个1进一位
			if i == 0 {
				newArray = append(newArray, 1)
			}
		}
	}

	newLeng := len(newArray)
	for j := 0; j < newLeng/2; j++ {
		newArray[j], newArray[newLeng-1-j] = newArray[newLeng-1-j], newArray[j]
	}
	return newArray
}

//梯子一共n级，每步可以上1级或者2级，从梯底上到梯顶，一共有多少种方法.
// 其实是一个斐波那契数列
func climbingStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return climbingStairs(n-1) + climbingStairs(n-2)
	}
}

//给定一个二进制位数n，给出符合格雷码规范的二进制序列对应的十进制数的序列
/*
比如给定n=2
00 --- 0
01 --- 1
11 --- 3
10 --- 2
给出序列为：0,1,3,2
**/
func GrayCode(n int) []int {
	totalSource := createSerises(n)
	fmt.Println(totalSource)
	res := []int{}
	for _, v := range totalSource {
		val1 := 0
		for k1, v1 := range v {
			val2 := v1 * (int(math.Pow(float64(2), float64(n-1-k1))))
			val1 += val2
		}
		res = append(res, val1)
	}
	return res
}

//产生格雷码
//步骤：第一步，改变最右边的位元值；第二步，改变右起第一个为1的位元的左边位元，重复第一步，重复第二步
//这个题碰到了golang中slice这种引用传值的极大不方便，如果将slice2:=slice1,后面对slice2的所有操作对slice1都是可见的，很不方便。。。
func createSerises(n int) [][]int {

	// source := []int{}
	grayCode := [][]int{}
	total := int(math.Pow(float64(2), float64(n)))
	var tempsource [][]int
	for m := 0; m < total; m++ {
		tempsource = append(tempsource, []int{})
	}
	for i := 0; i < n; i++ {
		for m := 0; m < total; m++ { //要不是slice这种变态的指针传递属性，这个for循环就不用了
			tempsource[m] = append(tempsource[m], 0)
		}
	}
	num := 1
	ind := 0
	for num < total {
		if ind == 0 {
			for m := num; m < total; m++ { //要不是slice这种变态的指针传递属性，这个for循环就不用了
				tempsource[m][n-1] = 1 - tempsource[m][n-1]
			}
			ind = 1
		} else {
			ind = 0
			for j := n - 1; j >= 1; j-- {
				if tempsource[num][j] == 1 {
					for m := num; m < total; m++ { //要不是slice这种变态的指针传递属性，这个for循环就不用了
						tempsource[m][j-1] = 1 - tempsource[m][j-1]
					}
					break
				}
			}
		}
		num += 1
	}
	for _, v := range tempsource {
		grayCode = append(grayCode, v)
	}
	return grayCode
}

//没看懂题目是什么意思？？
func setMatrixZeroes(source [][]int) [][]int {
	return [][]int{}
}

//车在气站圈兜，能否回到终点的问题
func gasStation(gas []int, cost []int) (index int) {
	leng := len(gas)
LOOP:
	for ind := 0; ind < leng; ind++ {
		for i := ind; i < leng+ind; i++ {
			getGas := 0
			costGas := 0
			for j := ind; j <= i; j++ {
				getGas += gas[j%leng]
			}
			for k := ind; k <= i; k++ {
				costGas += cost[k%leng]
			}
			if costGas > getGas {
				continue LOOP
			}
		}
		return ind
	}

	return -1

}

//糖果题，N个小孩站成一列，规则：1，每个小孩至少一颗糖；2，分数多的小孩比隔壁小孩的糖要给的多
func candy(ratingVal []float64) (sum int) {
	//首先找出分数最小的小孩，给 一颗糖，然后以这个小孩为准，分别向左left移动，向右right移动
	leng := len(ratingVal)
	if leng < 1 {
		return 0
	}
	candyDetail :=[leng]int{}
	min := 0
	for i := min; i < leng; i++ {
		if ratingVal[i] < ratingVal[min] {
			min = i
		}
	}
	//最少分数的小孩给一颗糖果
	candyDetail[min] = 1
	if min > 0 {
		indleft := 1
		for i := min - 1; i > 0; i-- {
			if ratingVal[i] > ratingVal[i+1] {
				indleft += 1
				sum += indleft
			} else {
				//Todo
				sum += candy(ratingVal[:])
			}

		}
	}

	if min < leng-1 {
		indright := 1
		for j := min + 1; j < leng; j++ {
			if ratingVal[j] > ratingVal[j-1] {
				indright += 1
				sum += indright
			} else {
				//Todo
				res, ind := candy(ratingVal[:])
			}

		}
	}
	return sum
}
