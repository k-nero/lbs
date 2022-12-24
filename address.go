package lbs

import (
	"encoding/json"
	"errors"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	ErrGeoNotPoint = errors.New("geometry is not a point")
)

type AddressEntity struct {
	//Country or region
	Country string
	//State or province
	State   string
	City    string
	ZipCode string
	Line1   string
	Line2   string
	Line3   string
	Geo     *Point
}

func NewAddressEntityFromPb(s *Address) (*AddressEntity, error) {
	ret := &AddressEntity{
		Country: s.Country,
		State:   s.State,
		City:    s.City,
		ZipCode: s.ZipCode,
		Line1:   s.Line1,
		Line2:   s.Line2,
		Line3:   s.Line3,
	}
	if s.Geo != nil {
		b, err := s.Geo.MarshalJSON()
		if err != nil {
			return nil, err
		}
		var geometry geom.T
		if err := geojson.Unmarshal(b, &geometry); err != nil {
			return nil, err
		}
		point, ok := geometry.(*geom.Point)
		if !ok {
			return nil, ErrGeoNotPoint
		}
		ret.Geo = NewPoint(point)
	}
	return ret, nil
}

func (s *AddressEntity) ToPb() (*Address, error) {
	ret := &Address{
		Country: s.Country,
		State:   s.State,
		City:    s.City,
		ZipCode: s.ZipCode,
		Line1:   s.Line1,
		Line2:   s.Line2,
		Line3:   s.Line3,
	}
	if s.Geo != nil {
		b, err := geojson.Marshal(s.Geo.Point.Point)
		if err != nil {
			return nil, err
		}
		v := map[string]interface{}{}
		err = json.Unmarshal(b, &v)
		if err != nil {
			return nil, err
		}
		s, err := structpb.NewStruct(v)
		if err != nil {
			return nil, err
		}
		ret.Geo = s
	}
	return ret, nil
}
