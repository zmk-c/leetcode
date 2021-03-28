package main

type MyStack struct {
	queue []int //实现具体的入栈出栈操作
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	//入栈 即直接追加入队queue即可
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	//队列queue中的最后一个元素  即为栈顶元素 返回  出栈
	val := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]

	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	val := this.queue[len(this.queue)-1]

	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
