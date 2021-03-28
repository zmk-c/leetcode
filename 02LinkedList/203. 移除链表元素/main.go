package main

import "fmt"

//ListNode Defination
type ListNode struct {
	Val  int
	Next *ListNode
}

//利用虚拟头指针对链表进行统一操作
func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{0, head} //新建虚拟头指针
	cur := dummyNode
	for cur.Next != nil {
		if cur.Next.Val == val { //找到指定的val值
			cur.Next = cur.Next.Next //指向下一个节点
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next //返回头节点
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	head = removeElements(head, 6)
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
