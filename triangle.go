package pencil

import (
	"image"
	"image/color"
	"image/draw"
)

func FillTriangle(a, b, c image.Point, color color.Color, img draw.Image) {
	var ymin, ymax int
	var xmin, xmax int
	for _, pt := range []image.Point{a, b, c} {
		if pt.Y < ymin {
			ymin = pt.Y
		}
		if pt.Y > ymax {
			ymax = pt.Y
		}
		if pt.X < xmin {
			xmin = pt.X
		}
		if pt.X > xmax {
			xmax = pt.X
		}
	}

	var triangles [3][2]float64
	for i, pt := range []image.Point{a, b, c} {
		triangles[i] = [2]float64{
			float64(pt.X), float64(pt.Y),
		}
	}
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y < ymax; y++ {
			u, v, _ := weightsFor(triangles, [2]float64{float64(x), float64(y)})
			if u >= 0 && v >= 0 && (u+v) <= 1 {
				img.Set(x, y, color)
			}
		}
	}
}

func weightsFor(tri [3][2]float64, pt [2]float64) (u, v, w float64) {
	a := tri[0]
	b := tri[1]
	c := tri[2]
	x, y := pt[0], pt[1]
	x1, y1 := a[0], a[1]
	x2, y2 := b[0], b[1]
	x3, y3 := c[0], c[1]

	// https://en.wikipedia.org/wiki/Barycentric_coordinate_system#Conversion_between_barycentric_and_Cartesian_coordinates
	coeff := (y2-y3)*(x1-x3) + (x3-x2)*(y1-y3)
	lambda1 := ((y2-y3)*(x-x3) + (x3-x2)*(y-y3)) / coeff
	lambda2 := ((y3-y1)*(x-x3) + (x1-x3)*(y-y3)) / coeff
	return lambda1, lambda2, 1 - lambda1 - lambda2
}
