package wireanalyzer

type Line struct {
	p1 *Point
	p2 *Point
}

func NewLine(p1 *Point, p2 *Point) *Line {
	if p1 == nil || p2 == nil || p1.Equal(p2) {
		return nil
	}

	return &Line{
		p1: p1,
		p2: p2,
	}
}

func (u *Line) IsVertical() bool {
	if u.p1.X == u.p2.X {
		return true
	}

	return false
}

func (u *Line) IsHorizontal() bool {
	return !u.IsVertical()
}

func (u *Line) XMin() int {
	if u.p1.X < u.p2.X {
		return u.p1.X
	}

	return u.p2.X
}

func (u *Line) XMax() int {
	if u.p1.X > u.p2.X {
		return u.p1.X
	}

	return u.p2.X
}

func (u *Line) YMin() int {
	if u.p1.Y < u.p2.Y {
		return u.p1.Y
	}

	return u.p2.Y
}

func (u *Line) YMax() int {
	if u.p1.Y > u.p2.Y {
		return u.p1.Y
	}

	return u.p2.Y
}

func (u *Line) Length() int {
	if u.IsVertical() {
		return u.YMax() - u.YMin()
	}

	return u.XMax() - u.XMin()
}

func (u *Line) Split(p *Point) (*Line, *Line) {
	if u.IsVertical() {
		return NewLine(u.p1, NewPoint(u.p1.X, p.Y)), NewLine(NewPoint(u.p1.X, p.Y), u.p2)
	}

	return NewLine(u.p1, NewPoint(p.X, u.p1.Y)), NewLine(NewPoint(p.X, u.p1.Y), u.p2)
}

func (u *Line) GetIntersection(v *Line) *Point {
	if u.IsVertical() {
		if v.IsVertical() {
			return nil // Parallel
		} else if v.XMin() > u.XMin() || v.XMax() < u.XMin() {
			return nil // Orthogonal but misaligned on X
		} else if v.YMin() < u.YMin() || v.YMin() > u.YMax() {
			return nil // Orthogonal but misaligned on Y
		} else {
			return NewPoint(u.XMin(), v.YMin())
		}
	} else {
		if v.IsHorizontal() {
			return nil // Parallel
		} else if v.YMin() > u.YMin() || v.YMax() < u.YMin() {
			return nil // Orthogonal but misaligned on Y
		} else if v.XMin() > u.XMax() || v.XMax() < u.XMin() {
			return nil // Orthogonal but misaligned on X
		} else {
			return NewPoint(v.XMin(), u.YMin())
		}
	}
}
