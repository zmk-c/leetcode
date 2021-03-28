package main

import "fmt"

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; j-- {
		s[i], s[j] = s[j], s[i]
		i++
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}
