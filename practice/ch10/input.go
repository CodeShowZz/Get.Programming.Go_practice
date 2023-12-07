package main

import "fmt"

func main() {
	var str string = "12"
	var res bool
	switch str {
	case "yes", "true", "1":
		res = true
	case "no", "false", "0":
		res = false
	default:
		fmt.Println(str, "is not valid")
		return
	}
	fmt.Println("res is", res)
}
