package hyperloglog

import (
	"fmt"
	"math/bits"
)

var users = []uint8{
	0b01111101,
	0b11010001,
	0b10000111,
	0b01110001,
	0b01011010,
	0b01000001,
	0b01010011,
	0b00001001,
	0b01101001,
	0b10001100,
	0b00111110,
	0b00010000,
	0b11110001,
	0b01101110,
	0b00010010,
}

// HLog8 - повторяем вычисления с лекции.
func HLog8() (z float64) {
	buckets := [4]int{}

	//Determine buckets
	for _, v := range users {
		idx := v >> 6
		zeroes := bits.LeadingZeros8(v<<2) + 1
		if buckets[idx] < zeroes {
			buckets[idx] = zeroes
		}
	}

	//Считаем среднее
	for _, v := range buckets {
		divider := 2 << (v - 1)
		z = z + 1.0/float64(divider)
	}
	fmt.Println("Buckets:", buckets, "\nZ: ", z)

	const alpha = 0.5324
	z = alpha * 4 * 4 / z

	fmt.Printf(" Count: %f\n", z)
	return
}
