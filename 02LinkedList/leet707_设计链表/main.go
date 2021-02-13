package main

import "fmt"

/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：
get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：
MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：
所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 02LinkedList 库。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//定义结构体类型
type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

type MyLinkedList struct {
	dummyNode *LinkedNode //虚拟头节点  操作方便
	size      int         //整个链表的长度
}

/** Initialize your data structure here.*/ //构造函数  初始化为空
func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyNode: &LinkedNode{},
		size:      0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 { //不合法的索引
		return -1
	}
	curNode := this.dummyNode.Next
	for index != 0 {
		curNode = curNode.Next
		index--
	}
	return curNode.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &LinkedNode{ //新建准备插入的节点
		Val:  val,
		Next: nil,
	}

	newNode.Next = this.dummyNode.Next
	this.dummyNode.Next = newNode
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &LinkedNode{ //新建准备插入的节点
		Val:  val,
		Next: nil,
	}

	curNode := this.dummyNode
	index := this.size
	for index != 0 {
		curNode = curNode.Next
		index--
	}
	curNode.Next = newNode
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 { //如果要插入的位置大于链表的长度 或者插入的位置不合法 则无法插入
		return
	}
	/*插入的三种情况
	1.index=0,插入链表头
	2.index=this.size,插入链表尾
	3.0<index<this.size,在中间插入
	*/
	newNode := &LinkedNode{ //新建准备插入的节点
		Val:  val,
		Next: nil,
	}
	curNode := this.dummyNode //当前节点暂等于虚拟头节点
	for index != 0 {          //找插入位置的前一个索引
		curNode = curNode.Next
		index--
	}
	newNode.Next = curNode.Next
	curNode.Next = newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size || index < 0 {
		return
	}
	curNode := this.dummyNode
	for index != 0 {
		curNode = curNode.Next
		index--
	}
	curNode.Next = curNode.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated（举例说明） and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//打印列表
func (this *MyLinkedList) PrintList() {
	curNode := this.dummyNode.Next
	for curNode != nil {
		fmt.Print(curNode.Val, " ")
		curNode = curNode.Next
	}
}

func main() {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)     // 1
	myLinkedList.AddAtTail(10)    // 1->10
	myLinkedList.AddAtHead(7)     // 7->1->10
	myLinkedList.AddAtHead(3)     // 3->7->1->10
	myLinkedList.Get(1)           //返回1
	myLinkedList.AddAtIndex(2, 3) //3->7->3->1->10
	myLinkedList.DeleteAtIndex(2) //3->7->1->10
	myLinkedList.PrintList()
}
