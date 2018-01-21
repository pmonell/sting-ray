package main

type ShadeRec struct {
	HitAnObject   bool
	LocalHitPoint Vector
	Normal        Vector
	Color         RGBColor
	W             *World
}

func NewShadeRec(w *World) ShadeRec {
	return ShadeRec{
		HitAnObject:   false,
		LocalHitPoint: Vector{},
		Normal:        Vector{},
		Color:         w.BackgroundColor,
		W:             w,
	}
}
