package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)


func sincos(r, a float64) (x, y float64) {
	phi := -a * math.Pi / 180.0
	x = float64(r) * math.Cos(phi)
	y = float64(r) * math.Sin(phi)
	return
}


type Frame struct {
	x0, y0 float64
	x1, x2 float64 // x = x0 + x1 * x + x2 * y
	y1, y2 float64 // y = y0 + y1 * x + y2 * y
}


func (f *Frame) String() string {
	return fmt.Sprintf(
		"<x=(+%.3f +%.3f +%.3f) y=(+%.3f +%.3f +%.3f)>",
		f.x1, f.x2, f.x0,
		f.y1, f.y2, f.y0,
	)
}


func Eye() *Frame {
	f := Frame{ x1: 1.0, y2: 1.0 }
	return &f
}


func (f *Frame) Point(xx, yy float64) (x, y int) {
	x = int(math.Round(f.x0 + f.x1 * xx + f.x2 * yy))
	y = -int(math.Round(f.y0 + f.y1 * xx + f.y2 * yy))
	return
}


func (f *Frame) Polar(r, a float64) (x, y int) {
	px, py := sincos(r, a)
	return f.Point(px, py)
}


func (f *Frame) Translated(xx, yy float64) *Frame {
	g := Frame{
		x0: f.x0 + xx * f.x1 + yy * f.x2,
		y0: f.y0 + xx * f.y1 + yy * f.y2,
		x1: f.x1,
		x2: f.x2,
		y1: f.y1,
		y2: f.y2,
	}
	return &g
}


func (f *Frame) Rotated(a float64) *Frame {
	c, s := sincos(1.0, a)
	g := Frame{
		x0: f.x0,
		y0: f.y0,
		x1: c * f.x1 - s * f.y1,
		x2: s * f.x1 + c * f.y1,
		y1: c * f.x2 - s * f.y2,
		y2: s * f.x2 + c * f.y2,
	}
	return &g
}


func TextStyle(anchor string) string {
	// "text-anchor:left;font-size:20px;fill:black"
	return strings.Join(
		[]string{
			"text-anchor:", anchor,
			";font-size:20px;fill:black;stroke-width:1",
			";alignment-baseline:middle",
		},
		"",
	)
}


func ArrowHead(c *svg.SVG, f *Frame, a float64) {
	f = f.Rotated(a)
	fmt.Println("ah.r=", f)
	x1, y1 := f.Polar(10, 180 + 30)
	x2, y2 := f.Point(0, 0)
	x3, y3 := f.Polar(10, 180 - 30)
	c.Polyline(
		[]int{ x1, x2, x3 },
		[]int{ y1, y2, y3 },
	)
	return
}


func InverseKinematics() {

	fout, _ := os.Create("inverse-kinematics.svg")

	var c *svg.SVG = svg.New(fout)
	width := 600
	height := 400
	c.Start(width, height)
	c.Roundrect(0, 0, width, height, 15, 15,
		"fill:snow;stroke:black;stroke-width:3")

	var x1, y1, x2, y2 int

	e0 := Eye()
	fmt.Println("e0=", e0)

	org := e0.Translated(300, -200)
	fmt.Println("org=", org)

	c.Gstyle("fill:none;stroke:gray;stroke-width:1")
	{
		// X axis
		x1, y1 = org.Point(-280, 0)
		x2, y2 = org.Point(280, 0)
		c.Line(x1, y1, x2, y2)
		ArrowHead(c, org.Translated(280, 0), 0)
		x1, y1 = org.Point(270, 15)
		c.Text(x1, y1, "X", TextStyle("middle"))

		// Y axis
		x1, y1 = org.Point(0, -180)
		x2, y2 = org.Point(0, 180)
		c.Line(x1, y1, x2, y2)
		ArrowHead(c, org.Translated(0, 180), 90)
		x1, y1 = org.Point(15, 180)
		c.Text(x1, y1, "Y", TextStyle("middle"))
	}
	c.Gend()

	c.Gstyle("fill:none;stroke:black;stroke-width:3")
	{
		// X intercept
		u := org.Translated(30, 0)
		x1, y1 = u.Point(0, -10)
		x2, y2 = u.Point(0,  10)
		c.Line(x1, y1, x2, y2)

		x1, y1 = u.Point(0, -25)
		c.Text(x1, y1, "Px", TextStyle("middle"))

		// Y intercept
		u = org.Translated(0, 114)
		x1, y1 = u.Point(-10, 0)
		x2, y2 = u.Point( 10, 0)
		c.Line(x1, y1, x2, y2)

		x1, y1 = u.Point(-25, 0)
		c.Text(x1, y1, "Py", TextStyle("middle"))

		// P(latform)
		x1, y1 = org.Point(30, 114)
		c.Circle(x1, y1, 5, "fill:black")
		x1, y1 = org.Point(50, 115)
		c.Text(x1, y1, "P", TextStyle("middle"))
	}
	c.Gend()

	c.Gstyle("fill:none;stroke:black;stroke-width:3")
	{
		// Cross-Rail 1.
		u := org.Translated(-50, 0).Rotated(55)
		x1, y1 = u.Point(-180, 0)
		x2, y2 = u.Point(180, 0)
		c.Line(x1, y1, x2, y2)

		x1, y1 = u.Point(0, 0)
		c.Circle(x1, y1, 5, "fill:black")
		x1, y1 = org.Point(-65, 15)
		c.Text(x1, y1, "x1", TextStyle("middle"))
		x1, y1 = org.Point(-105, -120)
		c.Text(x1, y1, "R1", TextStyle("middle"))

		// Cross-Rail 2.
		v := org.Translated(110, 0).Rotated(125)
		x1, y1 = v.Point(-180, 0)
		x2, y2 = v.Point(180, 0)
		c.Line(x1, y1, x2, y2)

		x1, y1 = v.Point(0, 0)
		c.Circle(x1, y1, 5, "fill:black")
		x1, y1 = org.Point(125, 15)
		c.Text(x1, y1, "x2", TextStyle("middle"))
		x1, y1 = org.Point(165, -120)
		c.Text(x1, y1, "R2", TextStyle("middle"))

	}
	c.Gend()

	c.Gstyle("fill:none;stroke:black;stroke-width:3")
	{
		// h
		x1, y1 = org.Point(30, 15)
		x2, y2 = org.Point(30, 105)
		c.Line(x1, y1, x2, y2)
		x1, y1 = org.Point(40, 55)
		c.Text(x1, y1, "h", TextStyle("middle"))

		// b1
		x1, y1 = org.Point(-40, 0)
		x2, y2 = org.Point(25, 0)
		c.Line(x1, y1, x2, y2)
		x1, y1 = org.Point(-10, 15)
		c.Text(x1, y1, "b1", TextStyle("middle"))

		// b2
		x1, y1 = org.Point(35, 0)
		x2, y2 = org.Point(100, 0)
		c.Line(x1, y1, x2, y2)
		x1, y1 = org.Point(60, 15)
		c.Text(x1, y1, "b2", TextStyle("middle"))

	}
	c.Gend()

	c.End()
	fout.Close()
	return
}


func main() {
	InverseKinematics()

}
