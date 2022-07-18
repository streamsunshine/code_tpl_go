package question

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintNodeList(head *ListNode) {
	index := head
	for index != nil {
		fmt.Printf("var:%v\n", index.Val)
		index = index.Next
	}
}

func InitNodeList(len int) *ListNode {
	var next *ListNode

	for i := 0; i < len; i++ {
		tmp := &ListNode{
			Val:  len - i,
			Next: next,
		}
		next = tmp
	}
	return next
}
