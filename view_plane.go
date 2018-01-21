package main

type ViewPlane struct {
	HRes       int
	VRes       int
	PixelSize  float64
	Gamma      float64
	NumSamples int
	Center     Vector
}

func NewViewPlane() ViewPlane {
	return ViewPlane{
		HRes:       400,
		VRes:       400,
		PixelSize:  1.0,
		Gamma:      1.0,
		NumSamples: 100,
		Center:     Vector{0.0, 0.0, 10.0},
	}
}

func (vp *ViewPlane) Position(u float64, v float64) Vector {
	return Vector{}
}
