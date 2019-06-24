package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 1000, 1000)
	img := image.NewRGBA(rect)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			img.Set(i, j, color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			})
		}
	}
	for j := 0; j < 10; j++ {
		a := point{
			x: float64(rand.Int31n(1000)),
			y: float64(rand.Int31n(1000)),
		}
		img.Set(int(a.x), int(a.y), color.NRGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		})

		arr := shape(a, 7, 100)
		for i := 0; i < len(arr); i++ {
			if i < len(arr)-1 {
				l := line(arr[i], arr[i+1])
				// fmt.Printf("%v %v \n", arr[i], arr[i+1])
				for j := 0; j < len(l); j++ {
					img.Set(int(l[j].x), int(l[j].y), color.NRGBA{
						R: 255,
						G: 255,
						B: 255,
						A: 255,
					})
				}

			} else {
				l := line(arr[i], arr[0])
				for j := 0; j < len(l); j++ {
					img.Set(int(l[j].x), int(l[j].y), color.NRGBA{
						R: 255,
						G: 255,
						B: 255,
						A: 255,
					})
				}
			}
		}
	}
	f, _ := os.Create("photo.png")
	defer f.Close()
	_ = png.Encode(f, img)
}
