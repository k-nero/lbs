package lbs

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

// Point wrap wkb.Point
type Point wkb.Point

// GormDataType ...
func (g Point) GormDataType() string {
	return "geometry"
}
func NewPoint(p *geom.Point) *Point {
	return (*Point)(&wkb.Point{Point: p})
}
