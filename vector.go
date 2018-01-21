package main

import "math"

// class definition
type Vector struct {
	X, Y, Z float64
}

// constructor from coordinates
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// constructor from a vector pointer
func CopyVector(v Vector) Vector {
	return Vector{v.X, v.Y, v.Z}
}

// add a vector to the vector and return a vector
func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// subtract a vector from the vector and return a vector
func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// multiply by a float
func (a Vector) MultiplyByScalar(s float64) Vector {
	return Vector{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

// divide by scalar
func (a Vector) DivideByScalar(s float64) Vector {
	return Vector{
		X: a.X / s,
		Y: a.Y / s,
		Z: a.Z / s,
	}
}

// dot product
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// length of the vector
func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// cross product
func (a Vector) Cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// converts the vector to a unit vector
func (a Vector) Normalize() Vector {
	return a.MultiplyByScalar(1. / a.Length())
}

// reflect a vector
func (a Vector) Reflect(b Vector) Vector {
	return b.Sub(a.MultiplyByScalar(2 * a.Dot(b)))
}
