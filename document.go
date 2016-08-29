package gokml_gx

import (
	"encoding/xml"
)

/*
 * Container type Document
 */
type Document struct {
	XMLName     xml.Name    `xml:"Document"`
	Id          string      `xml:"id,attr,omitempty"`
	Name        string      `xml:"name,omitempty"`
	Visibility  int         `xml:"visibility,omitempty"`
	Open        int         `xml:"open,omitempty"`
	AtomAuthor  AtomAuthor  `xml:"http://www.w3.org/2005/Atom author,omitempty"`
	Address     string      `xml:"address,omitempty"`
	PhoneNumber string      `xml:"phoneNumber,omitempty"`
	Description string      `xml:"description,omitempty"`
	Schema      []Schema    `xml:"Schema"`
	DocStyle    []Style     `xml:"Style"`
	Placemarks  []Placemark `xml:"Placemark"`
	Folders     []Folder    `xml:"Folder"`
	Elements    []interface{}
}

// NewDocument() creates a new document
func NewDocument(id, name string) Document {
	d := Document{Id: id, Name: name}
	return d
}

// AddPlacemark() adds a placemark to the document
func (d *Document) AddPlacemark(p Placemark) {
	d.Placemarks = append(d.Placemarks, p)
}

// AddSchema() adds a schema to the document
func (d *Document) AddSchema(s Schema) {
	d.Schema = append(d.Schema, s)
}

// AddStyle() adds a style to the document
func (d *Document) AddStyle(s Style) {
	d.DocStyle = append(d.DocStyle, s)
}

// AddFolder() adds a folder to the document
func (d *Document) AddFolder(f Folder) {
	d.Folders = append(d.Folders, f)
}

// Document.Marshal() returns the marshalled xml structure
func (d *Document) Marshal() ([]byte, error) {
	return xml.MarshalIndent(d, "", "	")
}

// AddElement() adds a path to the document
func (d *Document) AddElement(e Element) {
	d.Elements = append(d.Elements, e)
}
