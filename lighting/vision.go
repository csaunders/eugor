package lighting

type Vision struct {
	intensity int
	x         int
	y         int
	raycaster Raycaster
}

func NewVision(x, y, intensity int, m *dungeon.TileMap) *Vision {
	raycaster := MakeRaycaster(m)
	raycaster.CastRays(x, y, intensity)
	return &Vision{x: x, y: y, intensity: intensity, raycaster: raycaster}
}

func (v *Vision) X() int {
	return v.x
}

func (v *Vision) Y() int {
	return v.y
}

func (v *Vision) Intensity() int {
	return v.intensity
}

func (v *Vision) IsLighting(x, y int) bool {
	return v.raycaster.IsLighting(x, y)
	// return math.Abs(float64(v.x-x)) <= float64(v.intensity) && math.Abs(float64(v.y-y)) <= float64(v.intensity)
}

func (v *Vision) Tick() {

}

func (v *Vision) Projection() Projection {
	return Static
}

func (v *Vision) Name() string {
	return "vision"
}

func (v *Vision) ToString() string {
	return ""
}
