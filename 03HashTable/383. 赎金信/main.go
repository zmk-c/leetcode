package main

import "fmt"

/*
哈希解法：
因为题目所只有小写字母，那可以采用空间换取时间的哈希策略， 用一个长度为26的数组还记录magazine里字母出现的次数。
然后再用ransomNote去验证这个数组是否包含了ransomNote所需要的所有字母。
依然是数组在哈希法中的应用。
*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var record [26]int
	for _, char := range magazine { //记录magazine中字符出现的次数
		record[char-'a']++
	}

	for _, char := range ransomNote { //用ransomNote验证这个数组是否包含了ransomNote所需要的所有字母
		record[char-'a']--
		if record[char-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	flag := canConstruct("aa", "aab")
	fmt.Println(flag)
}
