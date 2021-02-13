package main

import "fmt"

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。

进阶：
你是否可以使用 O(1) 空间解决此题？

提示：
链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*环形链表II 主要考虑两个问题：
1.判断链表是否有环？
A:利用双指针(快慢指针，fast指针两步一单位，slow指针一步一单位，如果两者相遇  说明存在环)
2.如果有环如何找到这个环的入口？
A:在两个指针相遇的位置设置一个指针index1,再在头节点设置一个指针index2,两者相遇的位置就是环入口处
*/
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { //快慢指针相遇位置
			index1 := head
			index2 := slow
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil //没有环
}
func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}

	head.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}

	head.Next.Next.Next = head.Next

	cycle := detectCycle(head)
	if cycle == nil {
		fmt.Println("不存在环")
	} else {
		fmt.Println("存在环，环入口为", cycle)
	}
}
