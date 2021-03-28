package main

import "fmt"

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//改变Next指针的指向  双指针法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var tmp *ListNode
	cur := head
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//打印输出列表
func (head *ListNode) printList() {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
