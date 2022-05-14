// 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。
// 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
package main

import (
	"fmt"
	"math"
)

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	xLen := math.Abs(float64(len(first) - len(second)))
	if xLen >= 2 {
		return false
	}

	//长度一样，只能交换
	// if xLen == 0 {
	// 	var a byte
	// 	var b byte
	// 	for i := 0; i < len(first); i++ {
	// 		if first[i] != second[i] {
	// 			if a == 0 {
	// 				a = first[i]
	// 				b = second[i]
	// 				continue
	// 			}
	// 			return a == second[i] && b == first[i]
	// 		}
	// 	}
	// }
	if xLen == 0 {
		flag := false
		for i := 0; i < len(first); i++ {
			if first[i] == second[i] {
				continue
			} else if flag {
				return false
			}
			flag = true
		}
	}

	//长度差1，只能在长的里面删除，或者短的里面插入
	if xLen == 1 {
		var l string
		var s string
		if len(first) > len(second) {
			l, s = first, second
		} else {
			l, s = second, first
		}

		for i := 0; i < len(s); i++ {
			if l[i] == s[i] {
				continue
			} else {
				return l[i+1:] == s[i:]
			}
		}

	}

	return true

}

func main() {
	fmt.Printf("%s and %s => %t", "pale", "ple", oneEditAway("pale", "ple"))

	fmt.Printf("%s and %s => %t", "pales", "pal", oneEditAway("pales", "pal"))

	fmt.Printf("%s and %s => %t", "abcdxabcde", "abcdeabcdx", oneEditAway("abcdxabcde", "abcdeabcdx"))
}
