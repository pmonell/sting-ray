package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"fmt"
	"math"
)

func TestAdd(t *testing.T) {
	v1 := Vector{1., 1., 1.}
	v2 := Vector{2., 2., 2.}
	v3 := v1.Add(v2)
	assert.Equal(t, v3, Vector{3., 3., 3.}, "should add correctly")
}

func TestSub(t *testing.T) {
	v3 := Vector{2., 2., 2.}.Sub(Vector{1., 1., 1.})
	assert.Equal(t, v3, Vector{1., 1., 1.}, "should get Vector{1,1,1}")
}

func TestMultiplyByScalar(t *testing.T) {
	v2 := Vector{1., 1., 1.}.MultiplyByScalar(2.)
	assert.Equal(t, v2, Vector{2., 2., 2.}, "should be {2., 2., 2.}")
}

func TestDot(t *testing.T) {
	v2 := Vector{1., 1., 1.}.Dot(Vector{2., 2., 2.})
	assert.Equal(t, v2, 6., "should be equal to 6")
}

func TestLength(t *testing.T) {
	vLength := Vector{1., 1., 1.}.Length()
	assert.Equal(t, vLength, math.Sqrt(3.), "should be equal to sqrt of 3")
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, Vector{1., 0., 0.}, Vector{10., 0., 0.}.Normalize())
}

func TestReflect(t *testing.T) {
	vReflection := Vector{10., 0, 0}.Reflect(Vector{0., 10., 0.})
	assert.Equal(t, Vector{0., 10., 0.}, vReflection, "should calculate the reflection correctly")
}
