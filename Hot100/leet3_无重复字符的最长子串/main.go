package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0
 

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*思路：
	只增大不减小的滑动窗口
	定义两个游标start,end,遍历字符串，其中end的游标随着遍历实现无重复的增大
	若出现重复字符，则两个游标都增大index+1位（窗口大小不变，start游标滑动到重复位置后一位）
	这时候由于窗口大小不变，窗口内可能存在重复，恰好遍历从start游标开始，如果没有重复，需要判断i+1与end的关系，
超过的话，即i遍历到窗口之外，end再增大
	遍历结束，end-start即为结果。
*/
func lengthOfLongestSubstring(s string) int {
	start,end := 0,0
	for i:=0;i<len(s);i++{
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i + 1 > end {
				end++
			}
		}else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

func main(){
	s := "abcabcbb"
	lenLongSubString := lengthOfLongestSubstring(s)
	fmt.Println(lenLongSubString)
}