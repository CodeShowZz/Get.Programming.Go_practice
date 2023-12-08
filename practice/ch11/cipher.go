package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	plainText = strings.ReplaceAll(plainText, " ", "")
	cipherText := strings.ToUpper(plainText)
	//加密后:ECFRZKYGLGRMUSDHRXK
	//加密前:YOURMESSAGEGOESHERE
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	for i := 0; i < len(plainText); i++ {
		c := cipherText[i]
		key := keyword[keyIndex]
		fmt.Println("c=", c, "key=", key)
		c = c + (key - 'A')
		if c > 'Z' {
			c = c - 91 + 'A'
		}
		// c = (c-key+26)%26 + 'A'
		message = message + string(c)
		keyIndex++
		keyIndex = keyIndex % len(keyword)
	}
	fmt.Println(message)

}

//解密过程 减法X-key=Y
//加密过程 加法X+key=Y
