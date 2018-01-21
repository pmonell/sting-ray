package main

type RGBColor struct {
	R, G, B float64
}

// constuctor for all one
func MakeRGBColor(c float64) RGBColor {
	return RGBColor{c, c, c}
}

func (rgb RGBColor) DivideByScalar(ns float64) RGBColor {
	return RGBColor{
		R: rgb.R / ns,
		G: rgb.G / ns,
		B: rgb.B / ns,
	}
}

func (rgb RGBColor) Add(color RGBColor) RGBColor {
	return RGBColor{
		R: rgb.R + color.R,
		G: rgb.G + color.G,
		B: rgb.B + color.B,
	}
}
