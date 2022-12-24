package lbs

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

// Point wrap wkb.Point
//
// https://stackoverflow.com/a/60577841
type Point struct {
	Point *wkb.Point
}

// GormDataType ...
func (p *Point) GormDataType() string {
	return "geometry"
}

func NewPoint(p *geom.Point) *Point {
	return &Point{&wkb.Point{Point: p}}
}

func (p *Point) Value() (driver.Value, error) {
	value, err := p.Point.Value()
	if err != nil {
		return nil, err
	}

	buf, ok := value.([]byte)
	if !ok {
		return nil, fmt.Errorf("did not convert value: expected []byte, but was %T", value)
	}

	mysqlEncoding := make([]byte, 4)
	binary.LittleEndian.PutUint32(mysqlEncoding, 4326)
	mysqlEncoding = append(mysqlEncoding, buf...)

	return mysqlEncoding, err
}

func (p *Point) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	mysqlEncoding, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("did not scan: expected []byte but was %T", src)
	}

	var srid uint32 = binary.LittleEndian.Uint32(mysqlEncoding[0:4])

	err := p.Point.Scan(mysqlEncoding[4:])

	p.Point.SetSRID(int(srid))

	return err
}
