package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLineSuccess(t *testing.T) {
	p1 := NewPoint(0, 1)
	p2 := NewPoint(2, 3)

	l := NewLine(p1, p2)
	if l == nil {
		t.Error(fmt.Errorf("NewLine return nil but should have returned valid *Line"))
	}

	assert.Equal(t, l.p1.X, 0)
	assert.Equal(t, l.p1.Y, 1)
	assert.Equal(t, l.p2.X, 2)
	assert.Equal(t, l.p2.Y, 3)
}

func TestNewLineFailP1(t *testing.T) {
	var p1 *Point
	p2 := NewPoint(2, 3)

	l := NewLine(p1, p2)
	if l != nil {
		t.Error(fmt.Errorf("NewLine return valid *Line but should return nil (point#1 was nil)"))
	}
}

func TestNewLineFailP2(t *testing.T) {
	p1 := NewPoint(0, 1)
	var p2 *Point

	l := NewLine(p1, p2)
	if l != nil {
		t.Error(fmt.Errorf("NewLine return valid *Line but should return nil (point#2 was nil)"))
	}
}

func TestNewLineFailEqual(t *testing.T) {
	p := NewPoint(0, 1)

	l := NewLine(p, p)
	if l != nil {
		t.Error(fmt.Errorf("NewLine return valid *Line but should return nil (points euqal)"))
	}
}

func TestLineIsVertical(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(1, 5)

	l := NewLine(p1, p2)

	assert.Equal(t, true, l.IsVertical())
	assert.Equal(t, false, l.IsHorizontal())
}

func TestLineIsHorizontal(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(5, 2)

	l := NewLine(p1, p2)

	assert.Equal(t, false, l.IsVertical())
	assert.Equal(t, true, l.IsHorizontal())
}

func TestLineXYMinMax(t *testing.T) {
	p1 := NewPoint(2, 3)
	p2 := NewPoint(0, 1)

	l := NewLine(p1, p2)

	assert.Equal(t, 0, l.XMin())
	assert.Equal(t, 2, l.XMax())
	assert.Equal(t, 1, l.YMin())
	assert.Equal(t, 3, l.YMax())
}

func TestLineGetIntersectionParallelVertical(t *testing.T) {
	p11 := NewPoint(1, 2)
	p12 := NewPoint(1, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(3, 2)
	p22 := NewPoint(3, 5)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionParallelHorizontal(t *testing.T) {
	p11 := NewPoint(2, 1)
	p12 := NewPoint(5, 1)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(2, 3)
	p22 := NewPoint(5, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalLL(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(0, 1)
	p22 := NewPoint(2, 1)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalML(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(0, 3)
	p22 := NewPoint(2, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalHL(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(0, 6)
	p22 := NewPoint(2, 6)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalLR(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(4, 1)
	p22 := NewPoint(7, 1)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalMR(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(4, 3)
	p22 := NewPoint(7, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalHR(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(4, 6)
	p22 := NewPoint(7, 6)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalLC(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 1)
	p22 := NewPoint(4, 1)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionOrthogonalMCR(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 3)
	p22 := NewPoint(3, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCM(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 3)
	p22 := NewPoint(4, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCL(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(3, 3)
	p22 := NewPoint(4, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCX1R(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 2)
	p22 := NewPoint(3, 2)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCX1M(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 2)
	p22 := NewPoint(4, 2)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCX1L(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(3, 2)
	p22 := NewPoint(4, 2)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCX2R(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 5)
	p22 := NewPoint(3, 5)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCX2M(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 5)
	p22 := NewPoint(4, 5)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalMCX2L(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(3, 5)
	p22 := NewPoint(4, 5)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	// TODO: Check exact intersection
}

func TestLineGetIntersectionOrthogonalHC(t *testing.T) {
	p11 := NewPoint(3, 2)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(1, 6)
	p22 := NewPoint(4, 6)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.Nil(t, intersection)
}

func TestLineGetIntersectionCustom1(t *testing.T) {
	p11 := NewPoint(8, 5)
	p12 := NewPoint(3, 5)
	l1 := NewLine(p11, p12)

	p21 := NewPoint(6, 7)
	p22 := NewPoint(6, 3)
	l2 := NewLine(p21, p22)

	intersection := l1.GetIntersection(l2)
	assert.NotNil(t, intersection)
	assert.Equal(t, NewPoint(6, 5), intersection)
}
