package gokml_gx

/*
 * Container type Atom Author
 */
type AtomAuthor struct {
	Name      string `xml:"name,omitempty"`
	Namespace string `xml:"xmlns,attr"`
}

/*
 * Folder represents the folder structure within a document or Kml structure
 */

// NewAtomAuthor() creates a new atom authro
func NewAtomAuthor(name string) AtomAuthor {
	f := AtomAuthor{Name: name}
	return f
}
