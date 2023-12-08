package main

import "fmt"

// WEDIGYOULUVTHEGOPHERS
func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] - 'A'
		key := keyword[keyIndex] - 'A'
		c = (c-key+26)%26 + 'A'
		message = message + string(c)
		keyIndex++
		keyIndex = keyIndex % len(keyword)
	}
	fmt.Println(message)
}
