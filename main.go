package main

import "fmt"

func main() {

	letters := "abcbdefghi"
	runes := []rune(letters)

	fmt.Printf("-> %s", string(runes))

}
