package main

import (
	"fmt"
	"math/rand"
)

func main() {
	moneys := [3]int{5, 10, 25}
	target := 2000
	total := 0
	for {
		money := moneys[rand.Intn(3)]
		total = total + money
		fmt.Printf("current select money is %d cent,total is %.2f$\n", money, (float64(total) / 100))
		if total > target {
			break
		}
	}

}
