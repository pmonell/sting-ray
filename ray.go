package main

// class definition
type Ray struct {
	Origin, Direction Vector
}

// calculates the point on a vector at distance t from the origin
func (r Ray) Position(t float64) Vector {
	return r.Origin.Add(r.Direction.MultiplyByScalar(t))
}
