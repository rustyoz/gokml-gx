package gokml_gx

import (
	"encoding/xml"
	"strings"
)

/*
 * Path
 */
type Path struct {
	XMLName     xml.Name     `xml:"Placemark"`
	Id          string       `xml:"id,attr,omitempty"`
	Name        string       `xml:"name,omitempty"`
	Description string       `xml:"description,omitempty"`
	StyleUrl    string       `xml:"styleUrl,omitempty"`
	Coordinates string       `xml:"LineString>coordinates"`
	Extended    ExtendedData `xml:"ExtendedData"`
	Style       Style        `xml:",omitempty"`
	TimeStamp   string       `xml:"TimeStamp,omitempty"`
}

// NewPath() creates a new Path.  All parameters are strings.
func NewPath(id, name, desc, lat, lon, alt, styleurl string) Path {
	p := Path{Id: id, Name: name, Description: desc}
	if styleurl != "" {
		p.StyleUrl = styleurl
	}
	p.SetCoordinates(lat, lon, alt)
	return p
}

// SetCoordinates takes latitude, longitude, and altitude
// and sets the Coordinates value of the Path.
func (p *Path) SetCoordinates(lat, lon, alt string) {
	s := []string{lon, lat, alt}
	if len(s[2]) > 0 {
		p.Coordinates = strings.Join(s, ",")
		return
	}
	p.Coordinates = strings.Join(s[0:2], ",")
}

// AddCoordinates takes latitude, longitude, and altitude
// and adds the Coordinate value to the path
func (p *Path) AddCoordinates(lat, lon, alt string) {
	s := []string{lon, lat, alt}
	if len(s[2]) > 0 {
		p.Coordinates = strings.Join(s, ",")
		return
	}
	p.Coordinates = p.Coordinates + " " + strings.Join(s[0:2], ",")
}

// AddExtendedData() adds an ExtendedData structure to the Path
func (p *Path) AddExtendedData(e ExtendedData) {
	p.Extended = e
}

// Path.Marshal()
func (p *Path) Marshal() ([]byte, error) {
	return xml.MarshalIndent(p, "", "	")
}
