package gokml_gx

import (
	"encoding/xml"
)

// Style for features
type Style struct {
	XMLName xml.Name  `xml:"Style"`
	Id      string    `xml:"id,attr,omitempty"`
	Line    LineStyle `xml:"LineStyle"`
	Icon    IconStyle `xml:"IconStyle"`
}

// IconStyle
type IconStyle struct {
	XMLName xml.Name `xml:"IconStyle"`
	Scale   string   `xml:"scale,omitempty"`
	Heading string   `xml:"heading,omitempty"`
	Href    string   `xml:"Icon>href,omitempty"`
}

// LineStyle
type LineStyle struct {
	XMLName xml.Name `xml:"LineStyle"`
	Color   string   `xml:"color,omitempty"`
	Width   string   `xml:"width,omitempty"`
}

func NewStyle(id string) Style {
	s := Style{Id: id}
	return s
}

func (s *Style) AddIconStyle(is IconStyle) {
	s.Icon = is
}

func (s *Style) AddLineStyle(ls LineStyle) {
	s.Line = ls
}

func NewIconStyle(scale, heading, href string) IconStyle {
	is := IconStyle{Scale: scale, Heading: heading, Href: href}
	return is
}

func NewLineStyle(color, width string) LineStyle {
	ls := LineStyle{Color: color, Width: width}
	return ls
}
