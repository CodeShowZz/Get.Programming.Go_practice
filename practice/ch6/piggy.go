package main

import (
	"fmt"
	"math/rand"
)

func main() {
	moneys := [3]float64{0.05, 0.10, 0.25}
	target := 20.0
	total := 0.0
	for {
		money := moneys[rand.Intn(3)]
		total = total + money
		fmt.Printf("current select money is %.2f,total is %.2f\n", money, total)
		if total > target {
			break
		}
	}
}
