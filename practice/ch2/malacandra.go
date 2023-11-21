package main 
import "fmt"

const (
DISTANCE = 56000_000
HOURS_OF_DAY = 24
)

func main() {
	speed := DISTANCE / 28 / HOURS_OF_DAY
	fmt.Println(speed,"km/h")
}