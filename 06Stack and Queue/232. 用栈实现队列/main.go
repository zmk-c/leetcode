package main

//MyQueue need a stackIn and a stackOut
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	//入队 只需要把数据存入stackIn栈即可
	this.stackIn = append(this.stackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//出队  则需要先判断stackOut栈中是否为空  如果不为空则直接从stackOut出栈
	//如果为空 则需要先将stackIn栈中的数据全部导入到stackOut中  再从stackOut出栈
	if len(this.stackOut) == 0 {
		for len(this.stackIn) != 0 {
			val := this.stackIn[len(this.stackIn)-1]
			//从stackIn中删除val
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
			//导入stackOut栈
			this.stackOut = append(this.stackOut, val)
		}
	}

	//获取stackOut栈顶元素
	val := this.stackOut[len(this.stackOut)-1]
	//移除stackOut栈顶元素
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	//返和出队的第一个元素
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	//复用Pop方法
	val := this.Pop()
	//再将val还回去  即进入stackOut栈中
	this.stackOut = append(this.stackOut, val)

	return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	//判断队列为空 即判断stackIn和stackOut两个栈是否都为空
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
