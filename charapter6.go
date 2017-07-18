package main

import (
	"fmt"
	// "math/rand"
)

type node struct {
	data int
	next *node
}

//初始化一个链表
var n0 = NewList([]int{0, 1, 3, 3, 4, 6, 10})

//初始化一个链表
var p0 = NewList([]int{0, 1, 1, 3, 7, 9, 15, 16})

//初始化一个链表
var t0 = NewList([]int{0, 2, 4, 6, 8, 13, 19, 38})

var np0 = NewList([]int{3, 2, 13, 6, 8, 100, 5, 49, 38})

func main() {
	//
	// res := MergeSortedArrayChg3([]int{1, 3, 5, 7, 9}, []int{2, 3, 5, 6, 7})
	// res := MergekSortedLists(n0, p0, t0)
	// res := InsertionSortList(n0, -1)
	res := SortList(np0)
	for res != nil {
		// fmt.Println(res.data)
		fmt.Println(res)
		res = res.next
	}
	// res := NewList([]int{1, 5, 3, 7, 9})
	// for res != nil {
	// 	// fmt.Println(res.data)
	// 	fmt.Println(res)
	// 	res = res.next
	// }
	// res1 := res.AddNewNode(4)
	// for res1 != nil {
	// 	// fmt.Println(res1.data)
	// 	fmt.Println(res1)
	// 	res1 = res1.next
	// }

}

func NewList(sa []int) *node {
	pa := &node{0, nil}
	pa1 := pa
	for i := 0; i < len(sa); i++ {
		pa.next = &node{sa[i], nil}
		pa = pa.next
	}
	return pa1.next
}

func (sa *node) AddNewNode(n int) (sa1 *node) {
	var prev *node
	if sa == nil {
		sa1 = &node{n, nil}
		return
	}
	prev = sa
	sa1 = sa
	for sa != nil {
		prev = sa
		sa = sa.next
	}
	prev.next = &node{n, nil}
	return

}

/*
Given two sorted integer arrays A and B, merge B into A as one sorted array.
Note: You may assume that A has enough space to hold additional elements from B. The number of
elements initialized in A and B are m and n respectively
**/
//假设sa和sb都从小到大排好序，新数组也从小到大排序
func MergeSortedArray(sa, sb []int) (s []int) {
	for i := 0; i < len(sb); i++ {
	LOOP:
		for j := len(sa) - 1; j >= 0; j-- {
			if sa[j] < sb[i] {
				sa = append(sa, 0)
				for k := len(sa) - 2; k > j; k-- {
					sa[k+1] = sa[k]
				}
				sa[j+1] = sb[i]
				break LOOP
			}
		}
	}
	return sa
}

/*
Given two sorted integer arrays A and B, merge B into A as one sorted array.
Note: You may assume that A has enough space to hold additional elements from B. The number of
elements initialized in A and B are m and n respectively
**/
//假设sa和sb都从小到大排好序，新数组也从小到大排序,相比较前一个方法，这里不是从len(sa)开始的，而是从flag开始的
func MergeSortedArrayChange(sa, sb []int) (s []int) {
	flag := len(sa) - 1
	for i := 0; i < len(sb); i++ {
	LOOP:
		for j := flag; j >= 0; j-- {
			if sa[j] < sb[i] {
				sa = append(sa, 0)
				for k := len(sa) - 2; k > j; k-- {
					sa[k+1] = sa[k]
				}
				sa[j+1] = sb[i]
				flag = j + 2
				break LOOP
			}
		}
	}
	return sa
}

//前面两种方法都比较笨，为啥不维护三个index，分别对应数组A，数组B和新数组C，然后A和B同时从后往前扫。
func MergeSortedArrayChg3(sa, sb []int) (s []int) {
	i := len(sa) - 1
	j := len(sb) - 1
	var sc []int
	for {
		if i == -1 && j == -1 {
			break
		}
		if i == -1 {

			sc = append(sc, sb[j])
			j--
		} else if j == -1 {

			sc = append(sc, sa[i])
			i--
		} else {
			if sa[i] > sb[j] {

				sc = append(sc, sa[i])
				i--
			} else {

				sc = append(sc, sb[j])
				j--
			}
		}

	}
	sc = reverse(sc)
	return sc
}

func reverse(sa []int) []int {
	for i := 0; i < (len(sa)-1)/2; i++ {
		sa[i], sa[len(sa)-1-i] = sa[len(sa)-1-i], sa[i]
	}
	return sa
}

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together
the nodes of the first two lists.

**/
//两个链表已经是从小到大有序的
//??为啥用sc = sa就不行，就断掉了呢？？那sc = sc.next为啥就没有断掉呢？？

func MergeTwoSortedLists(sa, sb *node) (s *node) {
	var sc = &node{}
	sc1 := sc
	for {
		if sa == nil && sb == nil {
			break
		}
		if sa == nil {
			sc.data = sb.data
			sc.next = sb.next
			break
		} else if sb == nil {
			sc.data = sa.data
			sc.next = sa.next
			break
		} else {
			if sa.data < sb.data {
				sc.data = sa.data
				sa = sa.next
			} else {
				sc.data = sb.data
				sb = sb.next
			}
			sc.next = &node{-1, nil}
			sc = sc.next
		}
	}
	return sc1
}

//复杂度为O(M+N)
func MergeTwoSortedListsChg(sa, sb *node) (s *node) {
	var sc = &node{}
	sc1 := sc
	for {
		if sa == nil && sb == nil {
			break
		}
		if sa == nil {
			sc.next = sb
			break
		} else if sb == nil {
			sc.next = sa
			break
		} else {
			if sa.data < sb.data {
				sc.next = sa
				sa = sa.next
			} else {
				sc.next = sb
				sb = sb.next
			}
			sc = sc.next
		}
	}
	return sc1.next
}

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
**/
//复杂度为O(M+N+P)相当理想
func MergekSortedLists(sa, sb, sc *node) *node {
	res := MergeTwoSortedListsChg(sa, sb)
	res1 := MergeTwoSortedListsChg(res, sc)
	return res1
}

func InsertionSortList(sa *node, s int) *node {
	var prev *node
	sa1 := sa
	if s < sa.data {
		prev = &node{s, sa}
		sa1 = prev
	} else {
		for sa != nil {
			if sa.data >= s {
				prev.next = &node{s, sa}
				break
			}
			prev = sa
			sa = sa.next
		}
	}
	return sa1
}

/*单链表的归并排序，
知识点：
1：找到中间节点
2：归并排序的整体思想-递归
3：两个有序链表合并成一个有序链表
**/
func SortList(sa *node) *node {
	//第一步，找到中间节点
	if sa == nil || sa.next == nil {
		return sa
	}
	midNode := getMidNode(sa)
	right := midNode.next
	midNode.next = nil
	//第二步：递归，分解成最小单元，
	//第三步，从最小单元开始两两合并
	//第二步和第三步是密不可分的，最难想的
	// fmt.Println(sa)
	// fmt.Println(right)
	res1 := SortList(sa)
	res2 := SortList(right)
	res := MergeTwoSortedListsChg(res1, res2)
	return res
}

//获取一个链表的中间节点，如果这个链表只有首节点，则返回首节点
func getMidNode(sa *node) *node {
	//一个指针每次走一步，一个指针每次走两步
	var p0 = sa
	var p1 = sa
	// prev := p1
	//如果p0.next = nil，因为中间是&&运算符，所以不会走p0.next.next，所以能保证不会因p0.next.next造成nil.next而崩溃，另外用prev这种方式也行
	for p0.next != nil && p0.next.next != nil {
		p0 = p0.next.next
		// prev = p1
		p1 = p1.next
	}
	//p1即为所找的中间节点
	return p1
}

func FirstMissingPositive() {

}
