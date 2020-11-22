package wireanalyzer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Path struct {
	points []*Point
}

func NewPathFromWire(w Wire) (*Path, error) {
	path := Path{}

	x := 0
	y := 0
	path.AddPoint(NewPoint(x, y))

	re := regexp.MustCompile(`[uUrRdDlL][1-9][0-9]*`)

	for _, v := range w {
		if !re.MatchString(v) {
			return nil, fmt.Errorf("unexpected wire segment: %s", v)
		}

		vDir := strings.ToLower(string(v[0]))
		vVal, err := strconv.Atoi(v[1:])
		if err != nil {
			return nil, fmt.Errorf("wire segement distance not an integer: %s", err)
		}

		switch vDir {
		case "u":
			y += vVal
		case "r":
			x += vVal
		case "d":
			y -= vVal
		case "l":
			x -= vVal
		}

		path.AddPoint(NewPoint(x, y))
	}

	return &path, nil
}

func (u *Path) AddPoint(point *Point) {
	u.points = append(u.points, point)
}

func (u *Path) GetLines() []*Line {
	lines := []*Line{}

	for i := 1; i < len(u.points); i++ {
		lines = append(lines, NewLine(u.points[i-1], u.points[i]))
	}

	return lines
}

func (u *Path) GetIntersections(v *Path) ([]*Point, error) {
	intersections := []*Point{}

	for _, l1 := range u.GetLines() {
		for _, l2 := range v.GetLines() {
			intersection := l1.GetIntersection(l2)
			if intersection != nil && (intersection.X != 0 && intersection.Y != 0) {
				intersections = append(intersections, intersection)
			}
		}
	}

	return intersections, nil
}

func (u *Path) StepsToIntersection(p *Point) int {
	steps := 0

	for _, l := range u.GetLines() {
		// fmt.Printf("[DEBUG] Path::StepsToIntersection line=%s\n", spew.Sdump(l))
		var intersection *Point
		if l.IsVertical() {
			intersection = l.GetIntersection(NewLine(NewPoint(p.X-1, p.Y), p))
		} else {
			intersection = l.GetIntersection(NewLine(NewPoint(p.X, p.Y-1), p))
		}
		if intersection != nil {
			la, _ := l.Split(p)
			steps += la.Length()
			break
		}

		steps += l.Length()
	}

	return steps
}
