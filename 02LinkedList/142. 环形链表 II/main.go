package main

import "fmt"

//ListNode Define
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
