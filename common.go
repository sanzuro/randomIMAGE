package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func line(a, b point) []point {
	// tengent := (a.y - b.y) / (a.x - b.x)
	arr := make([]point, 0)
	if math.Abs(a.x-b.x) > math.Abs(a.y-b.y) {
		maxDist := int(math.Abs(a.x - b.x))
		tengent := ((a.y - b.y) / (a.x - b.x))
		fmt.Printf("%v tengent \n", tengent)
		if a.x > b.x {
			for i := 0; i < maxDist; i++ {
				arr = append(arr, point{
					x: b.x + float64(i),
					y: b.y + float64(i)*tengent,
				})
			}
		} else {
			for i := 0; i < maxDist; i++ {
				arr = append(arr, point{
					x: a.x + float64(i),
					y: a.y + float64(i)*tengent,
				})
			}
		}
	} else {
		maxDist := int(math.Abs(a.y - b.y))
		tengent := ((a.x - b.x) / (a.y - b.y))
		fmt.Printf("%v tengent \n", tengent)
		if a.y > b.y {
			for i := 0; i < maxDist; i++ {
				arr = append(arr, point{
					x: b.x + float64(i)*tengent,
					y: b.y + float64(i),
				})
			}
		} else {
			for i := 0; i < maxDist; i++ {
				arr = append(arr, point{
					x: a.x + float64(i)*tengent,
					y: a.y + float64(i),
				})
			}
		}

	}
	return arr
}
func main() {
	a := point{
		x: 35,
		y: 0,
	}
	b := point{
		x: 6,
		y: 70,
	}
	fmt.Printf("%v \n ", line(a, b))
}
