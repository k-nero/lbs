package lbs

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

// Point wrap wkb.Point
type Point struct {
	*wkb.Point
}

// GormDataType ...
func (p *Point) GormDataType() string {
	return "geometry"
}

func NewPoint(p *geom.Point) *Point {
	return &Point{&wkb.Point{Point: p}}
}
