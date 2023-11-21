package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const ANSWER = 66
	random := rand.Intn(100) + 1
	for random != ANSWER {
		fmt.Print("you guess answer is ", random)
		if random > ANSWER {
			fmt.Print(" It is too big")
		} else {
			fmt.Print(" It is too small")
		}
		random = rand.Intn(100) + 1
		fmt.Println()
	}
	fmt.Println("you are right the answer is ", ANSWER)
}
