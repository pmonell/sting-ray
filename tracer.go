package main

type Tracer struct {
	World *World
}

func (t *Tracer) TraceRay(r Ray) RGBColor {
	sr := t.World.HitObjects(r)

	if sr.HitAnObject {
		return sr.Color
	}

	return t.World.BackgroundColor
}
