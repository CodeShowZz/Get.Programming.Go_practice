package main

import "fmt"

type kelvin float64
type fahrenheit float64
type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	c = c + 273.15
	return kelvin(c)
}

func (k kelvin) celsius() celsius {
	k -= 273.15
	return celsius(k)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	kelvin := kelvin(233.0)
	celsius := kelvin.celsius()
	fmt.Print(kelvin, "º K is ", celsius, "º C")

	fmt.Println()

	kelvin = 0.0
	fahrenheit := kelvin.fahrenheit()
	fmt.Print(kelvin, "º K is ", fahrenheit, "º F")
}
