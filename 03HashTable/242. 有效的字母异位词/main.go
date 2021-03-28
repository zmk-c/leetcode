package main

import "fmt"

//如果采取暴力解法需要时间为O(n^2)  两层for循环  t中的每一位都与s中的每一位字符进行比较并统计是否重复出现

//采取哈希表法  (数组record)  长度为26即字母长度，用来记录s中字符的个数，存在则+1,然后遍历t中的字符，如果该字符存在则-1，
//最后遍历record中的值与0比较，如果record[i]！=0,返回false，否则返回true          时间复杂度为o(n)  空间复杂度O(n)
func isAnagram(s string, t string) bool {
	var record = [26]byte{}

	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		record[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if record[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	flag := isAnagram(s, t)
	fmt.Println(flag)
}
