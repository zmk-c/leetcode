package main

import "fmt"

/*
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。
(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)

注意：
你可以假设两个字符串均只含有小写字母。
canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ransom-note
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
