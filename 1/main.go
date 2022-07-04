package main

import (
	"fmt"
)

/**
  TODO 第一题：字符串组合问题
 */

func main() {
	s := "abc"
	res := perm(s)
	fmt.Println("res=", res)
}

func perm(str string) []string {

	result := make([]string, 0)
	flags := make([]bool, len(str))
	var handler func(s string, sb string, k int, flags []bool)
	handler = func(s string, sb string, k int, flags []bool) {
		//完成一组
		if len(str) == k {
			result = append(result, sb)
			return
		}
		for i := 0; i < len(str); i++ {
			if !flags[i] {
				flags[i] = true
				sb += string(str[i])
				handler(str, sb, k+1, flags) //abc a 1  0=>true
				sb = sb[:len(sb)-1]
				flags[i] = false
			}
		}
	}
	handler(str, "", 0, flags)
	return result
}
