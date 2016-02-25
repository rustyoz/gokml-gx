package gokml_gx

import (
	"encoding/xml"
)

type Schema struct {
	XMLName xml.Name `xml:"Schema"`
	Id      string   `xml:"id,attr,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
}

type GxSimpleArrayField struct {
	XMLName     xml.Name `xml:"http://www.google.com/kml/ext/2.2 SimpleArrayField"`
	Name        string   `xml:"name,attr,omitempty"`
	Type        string   `xml:"type,attr,omitempty"`
	DisplayName string   `xml:"displayName,omitempty"`
}

func NewSchema(id, name string) Schema {
	s := Schema{Id: id, Name: name}
	return s
}
