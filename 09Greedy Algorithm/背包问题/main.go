package main

import "fmt"

/*
用贪心算法解决背包问题：
	n个物品的重量分别为wi，价值分别为vi，背包重量为C
	首先计算每种物品的单位重量的价值vi/wi，进行排序，然后依赖贪心选择更多的性价比最高的物品装入背包
	若将这种物品全部装入背包后，若背包的物品总重量未超过C，则选择性价比次高的装入...依次类推
*/

//参数：分别表示背包的容量，物品的价值，重量
//返回值：分别表示返回的最大价值和物品的选择状态(1表示全部，0表示没有放入，0.x表示物体放入部分)
func Knapsack(c float32, value []float32, weight []float32) (float32, []float32) {
	//1.计算物品的性价比
	length := len(weight)
	r := make([]float32, length) //r[] 存放的是物品的性价比
	for i := 0; i < length; i++ {
		r[i] = value[i] / weight[i]
	}

	//2.将性价比按照降序进行排序  这里就采用简单的冒泡排序了（当然可以采用快排）
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if r[i] <= r[j] {
				tmp := r[i]
				r[i] = r[j]
				r[j] = tmp

				//將價格和重量也按照性價比的順序進行排序
				k := value[i]
				value[i] = value[j]
				value[j] = k

				g := weight[i]
				weight[i] = weight[j]
				weight[j] = g
			}
		}
	}
	fmt.Println("物品的性价比数组：", r)
	fmt.Println("按照性价比排序后的物品的价值数组：", value)
	fmt.Println("按照性价比排序后的物品的重量数组：", weight)

	//3.贪心算法
	var i int
	var maxValue float32
	x := make([]float32, length) //存放物品是否选择放入背包的状态(1表示全部，0表示没有放入，0.x表示物体放入部分)
	for i = 0; i < length; i++ {
		if weight[i] > c { //当前物品的重量大于当前容量  则退出进入下步判断
			break
		}
		x[i] = 1             //说明当前物品的重量还<=当前容量   该物品可以完全放入
		maxValue += value[i] //加上该物品的价值
		c -= weight[i]       //当前容量减去放入的物品的重量
	}
	if i < length { // 当前物品的重量大于当前容量 则截取该物品的部分
		x[i] = c / weight[i]
		maxValue += x[i] * value[i]
	}
	return maxValue, x
}

func main() {
	value := []float32{4, 6, 3, 5, 6, 10}
	weight := []float32{5, 4, 2, 6, 2, 2}
	var c float32 = 10
	maxValue, x := Knapsack(c, value, weight)
	fmt.Printf("按照性价比排序后的选择：%v \n最大价值为%v", x, maxValue)
}
