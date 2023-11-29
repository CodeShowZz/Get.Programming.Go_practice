package main

import (
	"fmt"
	"math/rand"
)

const SECOND_OF_DAYS = 24 * 60 * 60

func main() {
	distance := 62100000
	fmt.Printf("%-20v %5v %-10v %3v", "Spaceline", "Days", "Trip type", "Price")
	spacelines := [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	fmt.Println()
	for i := 0; i < 10; i++ {
		spaceline := spacelines[rand.Intn(3)]
		speed := rand.Intn(15) + 16
		cost := speed + 20
		days := distance / SECOND_OF_DAYS / speed
		trip := rand.Intn(2)
		trip_type := ""
		if trip == 0 {
			trip_type = "One-way"
		} else {
			trip_type = "Round-trip"
			cost = cost * 2
		}
		fmt.Printf("%-20v %5v %-10v $ %3v", spaceline, days, trip_type, cost)
		fmt.Println()
	}

}
