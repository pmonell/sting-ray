package main

import (
	"math"
)

const kEpsilon = 0.00001

type GeometricObject interface {
	Hit(r Ray, tmin float64, s *ShadeRec) bool
	GetColor() RGBColor
}

type Sphere struct {
	Center Vector
	Radius float64
	Color  RGBColor
}

func (s Sphere) GetColor() RGBColor {
	return s.Color
}

func (s Sphere) Hit(r Ray, tmin float64, sr *ShadeRec) bool {
	var t float64
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	disc := b*b - 4.0*a*c

	if disc < 0.0 {
		return false
	}

	e := math.Sqrt(disc)
	denom := 2.0 * a

	// check smaller root
	t = (-b - e) / denom
	if t > kEpsilon {
		tmin = t
		sr.Normal = oc.Add(r.Direction.MultiplyByScalar(t)).DivideByScalar(s.Radius)
		sr.LocalHitPoint = r.Origin.Add(r.Direction.MultiplyByScalar(t))
		return true
	}
	// check larger root
	t = (-b + e) / denom
	if t > kEpsilon {
		tmin = t
		sr.Normal = oc.Add(r.Direction.MultiplyByScalar(t)).DivideByScalar(s.Radius)
		sr.LocalHitPoint = r.Origin.Add(r.Direction.MultiplyByScalar(t))
		return true
	}
	return false
}
