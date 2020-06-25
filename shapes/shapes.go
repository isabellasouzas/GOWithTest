package shapes


type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0

}

type Circle struct {
	Radius float64 
}

func 
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
