package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
思路：对两个链表同时遍历，对应位进行累加，如果有进位carry(初始进位carry=0)，则对应的累加和为两个链表对应位的值加上进位。
对应位数不匹配的，可以认为链表后面表示为0
结果链中的和sum=（n1+n2+carry）%10,新进位=(n1+n2+carry)/10
如果最后遍历结束，还有进位，则直接将结果=该进位
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	//初始化进位为0
	carry := 0
	for l1 != nil || l2 != nil{
		n1,n2 := 0,0
		if l1 != nil{
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil{
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 +carry) / 10

		if head == nil{
			head = &ListNode{Val:sum}
			tail = head
		}else{
			tail.Next = &ListNode{Val:sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val:carry}
	}
	return
}

func main(){
	/*测试数据的时候遗忘了 应该使用嵌套结构体因为开辟的是同一个地址的数据，而不应该创建多个结构体，让Next指向，这是独立的无关*/
	//嵌套struct
	l1 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	head := addTwoNumbers(l1, l2)
	fmt.Println("最终输出的结果为:")
	for head != nil{
		if head.Next == nil{
			fmt.Print(head.Val)
		}else {
			fmt.Print(head.Val,"->")
		}
		head = head.Next
	}

}
