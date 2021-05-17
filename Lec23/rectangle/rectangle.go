package rectangle

type Rectangle struct {
	A, B int
}

func New(newA, newB int) *Rectangle {
	return &Rectangle{newA, newB}
}
