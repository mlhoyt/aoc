package main

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (u *Point) Equal(v *Point) bool {
	if u.X == v.X && u.Y == v.Y {
		return true
	}

	return false
}
