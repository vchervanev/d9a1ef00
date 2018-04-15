package model

import (
	"io"
)

var id = []byte("id")
var place = []byte("place")
var country = []byte("country")
var city = []byte("city")
var distance = []byte("distance")

var attributeNames = [][]byte{id, place, country, city, distance}
var attributeTypes = []byte{'n', 's', 's', 's', 'n'}

type locationFactory struct {
	ModelDefinition
}

func (location *locationFactory) Allocate(src [][]byte) Instance {
	return &Location{id: 0, place: src[1], country: src[2], city: src[3], distance: 0}
}

var LocationFactory = &locationFactory{ModelDefinition{attributeNames, attributeTypes}}

type Location struct {
	id       int
	place    []byte
	country  []byte
	city     []byte
	distance int
}

func (loc *Location) Model() Model {
	return LocationFactory
}

func (loc *Location) Id() int {
	return loc.id
}
func (loc *Location) Write(writer io.Writer, index byte) {
	switch index {
	case 0:
		writer.Write([]byte(string(loc.id)))
	case 1:
		writer.Write(loc.place)
	case 2:
		writer.Write(loc.country)
	case 3:
		writer.Write(loc.city)
	case 4:
		writer.Write([]byte(string(loc.distance)))
	}
	panic("critical app configuration error")
}

func (loc *Location) Update([][]byte) {

}
