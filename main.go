package main

func main() {
	// initialize empty objects array for world initialization
	var objects []GeometricObject

	// initialize the world
	world := World{
		vp:              NewViewPlane(),
		BackgroundColor: RGBColor{1.0, 1.0, 1.0},
		Elements:        objects,
	}

	world.AddObject(Sphere{
		Center: Vector{0., 0., 20.},
		Radius: 20.,
		Color:  RGBColor{1., 0., 0.},
	})

	world.AddObject(Sphere{
		Center: Vector{10., 10., -20.},
		Radius: 10.,
		Color:  RGBColor{0., 1., 0.},
	})

	world.RenderPerspective()
}
