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
	// res := candy([]float64{1.1, 2.3, 2.1, 5.3, 9.2, 4.5, 3.6})
	// fmt.Println(res)

	// res := findSingleNumII([]int{1, 2, 3, 4, 5, 5, 4, 3, 1})
	// fmt.Println(res)
	// res := AddTwoNumbers([]int{2, 4, 3}, []int{5, 6, 4})
	// fmt.Println(res)
	// res, err := reverseLinkedListII([]int{1, 2, 3, 4, 5}, 2, 4)
	// if err == nil {
	// 	fmt.Println(res)
	// } else {
	// 	fmt.Println(err)
	// }
	res := partitionList([]int{1, 4, 3, 2, 5, 2}, 3)
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
	candyDetail := []int{}
	//每个小孩给一颗糖
	ii := 0
	for ii < leng {
		candyDetail = append(candyDetail, 1)
		ii++
	}
	// 从左到右遍历，如果第i+1个小孩的等级比第i个小孩的等级要高，则第i+1个小孩的糖火数是第i个小孩的糖果数+1
	for i := 0; i < leng-1; i++ {
		if ratingVal[i+1] > ratingVal[i] {
			candyDetail[i+1] = candyDetail[i] + 1
		}
	}
	// 从右到左遍历，如果第j-1个小孩的等级比第j个小孩的等级要高，且第j-1个小孩的糖不比第j个小孩的糖要多，则第j-1个小孩的糖火数是第j个小孩的糖果数+1
	for j := leng - 1; j > 0; j-- {
		if ratingVal[j-1] > ratingVal[j] && candyDetail[j-1] <= candyDetail[j] {
			candyDetail[j-1] = candyDetail[j] + 1
		}
	}

	for k := 0; k < leng; k++ {
		sum += candyDetail[k]
	}
	return sum
}

//给一个数组，除了一个数之外，其他数都出现了两次，找出这个数
func findSingleNum1(source []int) []int {
	hashm := make(map[int]bool)
	for _, val := range source {
		if _, ok := hashm[val]; ok {
			delete(hashm, val)
		} else {
			hashm[val] = true
		}
	}
	res := []int{}
	for k, _ := range hashm {
		res = append(res, k)
	}
	return res //这个return-1只是废话，但是得满足函数的返回值类型
}

//根据两个相同数的异或结果为0,0和x的异或结果为x，依次将每个元素异或一次，最终结果就是那个落单的元素
func findSingleNum2(source []int) int {
	num := 0
	for i := 0; i < len(source); i++ {
		num ^= source[i]
	}

	return num
}

//给一个数组，除了一个数之外，其他数都出现了三次，找出这个数
func findSingleNumII1(source []int) []int {
	hashm := make(map[int]int)
	for _, val := range source {
		if _, ok := hashm[val]; ok {
			hashm[val] += 1
		} else {
			hashm[val] = 1
		}
	}
	res := []int{}
	for k, v := range hashm {
		if v != 3 {
			res = append(res, k)
		}

	}
	return res
}

// //位运算,就是3->11，如果3出现了3次，则相加后11变成了33(二进制位运算不进位)，就是二进制位数的和也是3的倍数。这样就可以剔除掉出现3次的那些数
// func findSingleNumII2(source []int) []int {

// }

/*
给定两个链表，返回两个链表的和
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/
func AddTwoNumbers(source1, source2 []int) []int {
	leng := len(source1)
	newSource := []int{}
	extra := 0
	for i := 0; i < leng; i++ {
		sum := source1[i] + source2[i] + extra
		if sum > 9 {
			extra = sum / 10
			sum = sum % 10
		}
		newSource = append(newSource, sum)
	}
	return newSource
}

/*
Reverse a linked list from position m to n. Do it in-place and in one-pass.
For example: Given 1->2->3->4->5->nullptr, m = 2 and n = 4,
return 1->4->3->2->5->nullptr.
Note: Given m, n satisfy the following condition: 1 ≤ m ≤ n ≤ length of list
*/
func reverseLinkedListII(source []int, m, n int) ([]int, error) {
	leng := len(source)
	if !(m >= 1 && m <= n && n <= leng) {
		return nil, errors.New(`参数错误`)
	}
	// fmt.Println(source)
	for i := m - 1; i < (m+n-2)/2; i++ {
		source[i], source[m+n-2-i] = source[m+n-2-i], source[i]
	}
	return source, nil
}

/*
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater
than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.
For example, Given 1->4->3->2->5->2 and x = 3, return 1->2->2->4->3->5.
*/
func partitionList(source []int, param int) []int {
	ind := 0
	for i := 0; i < len(source); i++ {

	}
}
