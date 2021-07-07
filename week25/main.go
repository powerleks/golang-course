package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

func makeRuneArr(str string, substr string) []rune {
	length := utf8.RuneCountInString(str) + 1 + utf8.RuneCountInString(substr)
	zString := make([]rune, length)
	i := 0
	for len(substr) > 0 {
		r, size := utf8.DecodeRuneInString(substr)
		substr = substr[size:]
		zString[i] = r
		i++
	}
	zString[i] = 1114112
	for len(str) > 0 {
		i++
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		zString[i] = r
	}
	return zString
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func zFunction(str string, substr string) []int {
	runeArr := makeRuneArr(str, substr)
	zArr := make([]int, len(runeArr))
	l, r := 0, 0
	for i, _ := range runeArr {
		if i < r {
			zArr[i] = min(r-i+1, zArr[i-l])
		}
		for i + zArr[i] < len(runeArr) && runeArr[zArr[i]] == runeArr[i+zArr[i]] {
			zArr[i]++;
		}
		if i + zArr[i] - 1 > r {
			l = i
			r = i + zArr[i]-1
		}
	}
	return zArr
}

func findSubstring(str string, substr string) bool {
	zArr := zFunction(str, substr)
	substrLen := utf8.RuneCountInString(substr)
	for _, val := range zArr {
		if val == substrLen {
			return true
		}
	}
	return false
}

func main() {
    var str string
	var substr string

	flag.StringVar(&str, "str", "default", "set str")
	flag.StringVar(&substr, "substr", "default", "set substr")

	flag.Parse()

	fmt.Println(findSubstring(str, substr))
}