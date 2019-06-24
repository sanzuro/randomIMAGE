package main

import (
	"fmt"
	"math"
	"math/rand"
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
		// fmt.Printf("%v tengent \n", tengent)
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
		// fmt.Printf("%v tengent \n", tengent)
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
	// fmt.Printf("%v \n", arr)
	return arr
}

func shape(c point, n int, r float64) []point {

	arr := make([]point, 0)

	angle := float64(0)
	maxAngle := float64(360) / float64(n)
	// fmt.Printf("%v \n", maxAngle)
	for i := 0; i < n; i++ {
		// angle = (float64(i)*maxAngle + float64(rand.Int31n(int32(maxAngle)))/180) * math.Pi
		angle = ((float64(i)*maxAngle + float64(rand.Int31n(int32(maxAngle)))) / 180) * math.Pi
		fmt.Printf("%v \n", angle)
		t := point{
			x: c.x + r*math.Cos(angle),
			y: c.y + r*math.Sin(angle),
		}
		// fmt.Printf("%v\n", t)
		arr = append(arr, t)
	}
	return arr
}
