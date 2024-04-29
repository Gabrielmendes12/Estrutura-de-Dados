package main
 
import (
    "fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var value string
		fmt.Scanln(&value)
		fmt.Println(countLEDs(value), "leds")
	}
}

func countLEDs(value string) int {
	leds := map[byte]int{
		'0': 6, '1': 2, '2': 5, '3': 5, '4': 4,
		'5': 5, '6': 6, '7': 3, '8': 7, '9': 6,
	}

	totalLEDs := 0
	for i := 0; i < len(value); i++ {
		digit := value[i]
		totalLEDs += leds[digit]
	}

	return totalLEDs
}
