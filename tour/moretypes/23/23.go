package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount goes brr
func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	fmt.Printf("%#v\n", fields)

	m := make(map[string]int)

	for i, s := range fields {
		fmt.Println(i, s)
		m[s] = m[s] + 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
