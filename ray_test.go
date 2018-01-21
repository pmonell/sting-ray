package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostion(t *testing.T) {
	ray := Ray{
		Origin:    Vector{0., 0., 0.},
		Direction: Vector{3., 3., 3.},
	}
	assert.Equal(t, Vector{9., 9., 9.}, ray.Position(3.), "should calculate the correct positioning")
}
