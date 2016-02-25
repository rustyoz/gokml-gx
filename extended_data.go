package gokml_gx

import (
	"encoding/xml"
)

// ExtendedData for placemarks
type ExtendedData struct {
	XMLName xml.Name `xml:"ExtendedData"`
	Datas   []Data   `xml:"Data"`
}

// Data that goes in ExtendedData
type Data struct {
	XMLName xml.Name `xml:"Data"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value"`
}

// NewExtendedData() creates new extended data
func NewExtendedData() ExtendedData {
	return ExtendedData{}
}

// AddData() adds data of {name, value} to the ExtendedData structure
func (e *ExtendedData) AddData(name, value string) {
	d := Data{Name: name, Value: value}
	e.Datas = append(e.Datas, d)
}
