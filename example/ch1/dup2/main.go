package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
)


func main() {
	// content, err := ioutil.ReadFile("./hello")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("File contents: %s", content)
	// fmt.Printf("type: %T", content)
	names := []string{"aa", "bb", "cc"}
	for _, name := range names {
		fmt.Println(name)
		fmt.Println(string(name))
	}

}
