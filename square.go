package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	s.start.x = s.start.x + int(s.a)
	s.start.y = s.start.y - int(s.a)
	return s.start
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return uint(4) * s.a
}
