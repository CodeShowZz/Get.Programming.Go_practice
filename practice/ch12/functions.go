package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "º K is ", celsius, "º C")

	fmt.Println()

	kelvin = 0.0
	fahrenheit := kelvinToFahrenheit((kelvin))
	fmt.Print(kelvin, "º K is ", fahrenheit, "º F")

}
