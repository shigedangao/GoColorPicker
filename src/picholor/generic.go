package picholor

type common interface {
	buildJSON()
	handleError() []byte
	toRGB() (Rgb, error)
}
