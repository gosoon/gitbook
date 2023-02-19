package main

import "fmt"

func main() {
	s := "abcdasd"

	for i, c := range s {
		fmt.Println(i, c)
	}
}
