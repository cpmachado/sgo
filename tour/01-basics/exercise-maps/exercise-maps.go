package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for field := range strings.FieldsSeq(s) {
		res[field] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
