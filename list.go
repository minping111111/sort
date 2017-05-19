package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

//初始化一个链表
var (
	node6 = &node{6, nil}
	node5 = &node{5, node6}
	node4 = &node{4, node5}
	node3 = &node{3, node4}
	node2 = &node{2, node3}
	node1 = &node{1, node2}
	node0 = &node{0, node1}
)

//打印链表
func printList(head *node) *node {
	p := head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
	return p
}

func main() {
	// printList(node0)
	// newHead := listReverse_recursive(node0, nil)
	newHead := listReverse_recursive2(node0)
	// newHead := listReverse_for(node0)
	_ = printList(newHead)
}

//递归实现链表反转，传入头结点，返回新的头结点
func listReverse_recursive(current, prev *node) *node {
	if current == nil || current.next == nil {
		current.next = prev
		return current
	}
	//将下一次循环需要的节点保存
	tempNext := current.next
	current.next = prev
	prev = current
	return listReverse_recursive(tempNext, prev)

}

//递归实现链表反转，与上一个函数不同的是，这里只需要一个参数即可
var prev *node

func listReverse_recursive2(current *node) *node {
	if current != nil {
		tempnext := current.next
		current.next = prev
		prev = current
		return listReverse_recursive2(tempnext)
	} else {
		return prev
	}
}

func listReverse_for(current *node) *node {
	var prev *node
	for current != nil {
		tempnext := current.next
		current.next = prev
		prev = current
		current = tempnext
	}
	return prev

}
