package kmp

//当前字符串冲突时回退到前一位next数组对应的值
func GetNext(next []int, s string) {
	//初始化
	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		//前后缀不相同的情况
		for j > 0 && s[i] != s[j] {
			j = next[j-1] //回退
		}
		//前后缀相同的情况
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
