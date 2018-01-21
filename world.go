package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
)

const color = 255.99

func failOnError(e error, s string) {
	if e != nil {
		panic(e)
	}
}

type World struct {
	vp              ViewPlane
	BackgroundColor RGBColor
	Elements        []GeometricObject
}

func (w *World) RenderScene() {
	var buffer bytes.Buffer

	// ViewPlane specs
	hres := w.vp.HRes
	vres := w.vp.VRes
	pixelSize := w.vp.PixelSize
	numSamples := w.vp.NumSamples

	// tracer
	tracer := Tracer{World: w}

	ray := Ray{}
	ray.Direction = Vector{0, 0, -1}

	for j := vres - 1; j >= 0; j-- {
		for i := 0; i < hres; i++ {
			rgb := RGBColor{}

			for s := 0; s < numSamples; s++ {
				nx := pixelSize * (float64(i) - 0.5*float64(hres) + rand.Float64())
				ny := pixelSize * (float64(j) - 0.5*float64(vres) + rand.Float64())
				ray.Origin = Vector{nx, ny, 100.0}

				trace_rgb := tracer.TraceRay(ray)
				rgb = rgb.Add(trace_rgb)
			}

			rgb = rgb.DivideByScalar(float64(numSamples))

			ir := int(color * rgb.R)
			ig := int(color * rgb.G)
			ib := int(color * rgb.B)

			buffer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	writeToFile(buffer, hres, vres)
}

func (w *World) RenderPerspective() {
	var buffer bytes.Buffer
	tracer := Tracer{World: w}
	d := 10.0
	eye := -10.0
	ray := Ray{}
	ray.Origin = Vector{0.0, 0.0, eye}
	for r := w.vp.VRes - 1; r >= 0; r-- {
		for c := 0; c < w.vp.HRes; c++ {
			ray.Direction = Vector{
				w.vp.PixelSize * (float64(c) - 0.5*(float64(w.vp.HRes)-1.0)),
				w.vp.PixelSize * (float64(r) - 0.5*(float64(w.vp.VRes)-1.0)),
				-d,
			}
			ray.Direction = ray.Direction.Normalize()
			rgb := tracer.TraceRay(ray)

			ir := int(color * rgb.R)
			ig := int(color * rgb.G)
			ib := int(color * rgb.B)

			buffer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
	writeToFile(buffer, w.vp.HRes, w.vp.VRes)
}

func writeToFile(b bytes.Buffer, hres int, vres int) {
	f, err := os.Create("out.ppm")
	failOnError(err, "Couldn't create output file")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", hres, vres)
	failOnError(err, "Error writing to file")

	_, err = fmt.Fprintf(f, b.String())
	failOnError(err, "Error writing to file")
}

func (w *World) AddObject(object GeometricObject) {
	w.Elements = append(w.Elements, object)
}

func (w *World) HitObjects(r Ray) ShadeRec {
	sr := NewShadeRec(w)
	t := 0.0
	tmin := math.MaxFloat64

	for _, object := range w.Elements {
		if object.Hit(r, tmin, &sr) && t < tmin {
			sr.HitAnObject = true
			tmin = t
			sr.Color = object.GetColor()
		}
	}
	return sr
}
