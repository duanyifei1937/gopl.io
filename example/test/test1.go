package main

import (
	"fmt"
)

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func main() {

	lista := []string{"abc", "aaa", "bbb"}
	// Add(lista)
	// fmt.Println(Count(lista))
	fmt.Printf("%T", k(lista))
	fmt.Println(k(lista))
}
