package main

/*
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：
void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false


说明：
你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


进阶：
你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。


示例：
输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]
解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false


提示：
1 <= x <= 9
最多调用 100 次 push、pop、peek 和 empty
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
