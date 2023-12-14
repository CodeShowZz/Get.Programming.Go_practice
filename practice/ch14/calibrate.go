package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		// print("invoke")
		return s() + offset
	}
}

func main() {
	num := kelvin(5)
	sensor := calibrate(realSensor, num)
	num = kelvin(6)
	fmt.Println(sensor())
	sensor = calibrate(fakeSensor, 20)
	fmt.Println(sensor())
	sensor = calibrate(fakeSensor, 20)
	fmt.Println(sensor())
	sensor = calibrate(fakeSensor, 20)
	fmt.Println(sensor())
}
