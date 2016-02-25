package gokml_gx

import (
	"encoding/xml"
)

/*
 * Container type Atom Author
 */
type AtomAuthor struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom author"`
	Name    string   `xml:"name,omitempty"`
}

/*
 * Folder represents the folder structure within a document or Kml structure
 */

// NewAtomAuthor() creates a new atom authro
func NewAtomAuthor(name string) AtomAuthor {
	f := AtomAuthor{Name: name}
	return f
}
