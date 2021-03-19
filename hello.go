package main

import "fmt"

func testFunc(messageChanger string) {
	fmt.Printf("You passed this string in: %s", messageChanger)
}

func main() {
	testFunc("Blah")
}
