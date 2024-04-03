package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// fmt.Println(length)
	// fmt.Println(input)
	// fmt.Println(delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, delta, alphabetUpper))
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, delta, alphabetLower))
		default:
			ret = ret + string(ch)
		}

	}

	fmt.Println(ret)
}

func rotate(s rune, delta int, key string) rune {
	idx := strings.IndexRune(key, s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return []rune(key)[idx]
}
