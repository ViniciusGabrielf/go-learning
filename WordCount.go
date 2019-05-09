package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(WordCount("O tempo perguntou para o tempo quanto tempo o tempo tem"))
}
