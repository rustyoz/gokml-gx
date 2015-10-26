package gokml_gx

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	kml := CreateKml()

	b, err := kml.Marshal()
	if err != nil {
		t.Errorf("Error while marshalling: %q", err.Error())
	}
	fmt.Printf("KML:\n %s\n", b)
}

func TestUnmarshal(t *testing.T) {
	// kml := CreateKml()

	// b, err := kml.Marshal()
	// if err != nil {
	// 	t.Errorf("Error while marshalling: %q", err.Error())
	// }

	new_kml, err := Unmarshal([]byte(kmlExample))
	if err != nil {
		t.Errorf("Error while unmarshalling: %q", err.Error())
	}
	fmt.Printf("KML:\n %+v\n", new_kml)

	// testKml(new_kml, t)
}

func TestJsonMarshal(t *testing.T) {
	kml := CreateKml()

	b, err := kml.JsonMarshal()
	if err != nil {
		t.Errorf("Error while marshalling: %q", err.Error())
	}
	fmt.Printf("KML:\n %s\n", b)
}

func TestJsonUnmarshal(t *testing.T) {
	kml := CreateKml()

	b, err := kml.JsonMarshal()
	if err != nil {
		t.Errorf("Error while marshalling: %q", err.Error())
	}

	new_kml, err := JsonUnmarshal(b)
	if err != nil {
		t.Errorf("Error while unmarshalling: %q", err.Error())
	}
	fmt.Printf("KML:\n %s\n", new_kml)

	testKml(new_kml, t)
}

func testKml(kml Kml, t *testing.T) {
	// KML placemarks
	if len(kml.Placemarks) != 2 {
		t.Errorf("Expecting 2 placemarks, found: %d", len(kml.Placemarks))
	}
	kml_pm1 := kml.Placemarks[0]
	if kml_pm1.Id != "1" {
		t.Errorf("Placemark id should be 1, it was: %s", kml_pm1.Id)
	}
	// KML folders
	if len(kml.Folders) != 2 {
		t.Errorf("Expecting 2 folders, found: %d", len(kml.Folders))
	}
	kml_fd1 := kml.Folders[0]
	if kml_fd1.Id != "1" {
		t.Errorf("folder id should be 1, it was: %s", kml_fd1.Id)
	}
	// KML folder placemarks
	if len(kml_fd1.Placemarks) != 1 {
		t.Errorf("Expecting 1 placemarks, found: %d", len(kml_fd1.Placemarks))
	}
	fd_pm1 := kml_fd1.Placemarks[0]
	if fd_pm1.Id != "1" {
		t.Errorf("Placemark id should be 1, it was: %s", fd_pm1.Id)
	}

	// KML document
	doc := kml.Document
	if doc.Id != "123" {
		t.Errorf("Document id should be 123, it was: %s", doc.Id)
	}
	// Document Style
	if len(doc.DocStyle) != 2 {
		t.Errorf("Expecting 2 styles in the document, found: %d", len(doc.DocStyle))
	}
	doc_style1 := doc.DocStyle[0]
	if doc_style1.Id != "1" {
		t.Errorf("Document style 1 id should be 1, it was: %s", doc_style1.Id)
	}
	// Style icon
	doc_style1_icon := doc_style1.Icon
	if doc_style1_icon.Scale != "1" {
		t.Errorf("Document style icon scale should be 1, it was: %s", doc_style1_icon.Scale)
	}
	// Document placemarks
	if len(doc.Placemarks) != 2 {
		t.Errorf("Expecting 2 placemarks, found: %d", len(doc.Placemarks))
	}
	doc_pm1 := doc.Placemarks[0]
	if doc_pm1.Id != "3" {
		t.Errorf("Placemark id should be 3, it was: %s", doc_pm1.Id)
	}
	// Documents folders
	if len(doc.Folders) != 2 {
		t.Errorf("Expecting 2 folders, found: %d", len(doc.Folders))
	}
	doc_fd1 := doc.Folders[0]
	if doc_fd1.Id != "1" {
		t.Errorf("folder id should be 1, it was: %s", doc_fd1.Id)
	}
	// Document folder placemarks
	if len(doc_fd1.Placemarks) != 1 {
		t.Errorf("Expecting 1 placemarks, found: %d", len(doc_fd1.Placemarks))
	}
	doc_fd_pm1 := doc_fd1.Placemarks[0]
	if doc_fd_pm1.Id != "1" {
		t.Errorf("Placemark id should be 1, it was: %s", doc_fd_pm1.Id)
	}
}

func CreateKml() Kml {
	kml := NewKml("")

	kml_pm1 := NewPlacemark("1", "pm_name1", "desc", "123", "456", "alt", "styleurl")
	kml_pm2 := NewPlacemark("2", "pm_name2", "desc", "123", "456", "alt", "styleurl")
	kml.AddPlacemark(kml_pm1)
	kml.AddPlacemark(kml_pm2)

	kml_fd1 := NewFolder("1", "fd_name1")
	kml_fd2 := NewFolder("2", "fd_name2")
	fd_pm1 := NewPlacemark("1", "pm_name1", "desc", "123", "456", "alt", "styleurl")
	kml_fd1.AddPlacemark(fd_pm1)

	kml.AddFolder(kml_fd1)
	kml.AddFolder(kml_fd2)

	doc := NewDocument("123", "document_name")
	doc.Visibility = 1
	doc.Open = 0
	doc.Address = "123 ABC"
	doc.PhoneNumber = "123-456-7890"
	doc.Description = "a document with stuff"
	doc_pm1 := NewPlacemark("3", "doc_pm_name1", "desc", "123", "456", "alt", "styleurl")
	doc_pm2 := NewPlacemark("4", "doc_pm_name2", "desc", "123", "456", "alt", "styleurl")
	doc.AddPlacemark(doc_pm1)
	doc.AddPlacemark(doc_pm2)

	doc_fd1 := NewFolder("1", "fd_name1")
	doc_fd2 := NewFolder("2", "fd_name2")
	doc_fd_pm1 := NewPlacemark("1", "pm_name1", "desc", "123", "456", "alt", "styleurl")
	doc_fd1.AddPlacemark(doc_fd_pm1)

	doc.AddFolder(doc_fd1)
	doc.AddFolder(doc_fd2)

	style1 := NewStyle("1")
	icon_style1 := NewIconStyle("1", "heading", "href")
	style1.AddIconStyle(icon_style1)
	doc.AddStyle(style1)

	style2 := NewStyle("2")
	icon_style11 := NewIconStyle("1", "heading", "href")
	style2.AddIconStyle(icon_style11)
	doc.AddStyle(style2)

	author := NewAtomAuthor("author name")
	doc.AtomAuthor = author

	kml.AddDocument(doc)

	return kml
}

var kmlExample = `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns:atom="http://www.w3.org/2005/Atom" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns="http://www.opengis.net/kml/2.2">
<Document>
<open>1</open>
<visibility>1</visibility>
<name>
<![CDATA[254-260 Main St E]]>
</name>
<atom:author>
<atom:name>
<![CDATA[Created by Google My Tracks on Android]]>
</atom:name>
</atom:author>
<Style id="track">
<LineStyle>
<color>7f0000ff</color>
<width>4</width>
</LineStyle>
<IconStyle>
<scale>1.3</scale>
<Icon>
<href>http://earth.google.com/images/kml-icons/track-directional/track-0.png</href>
</Icon>
</IconStyle>
</Style>
<Style id="start">
<IconStyle>
<scale>1.3</scale>
<Icon>
<href>http://maps.google.com/mapfiles/kml/paddle/grn-circle.png</href>
</Icon>
<hotSpot x="32" y="1" xunits="pixels" yunits="pixels"/>
</IconStyle>
</Style>
<Style id="end">
<IconStyle>
<scale>1.3</scale>
<Icon>
<href>http://maps.google.com/mapfiles/kml/paddle/red-circle.png</href>
</Icon>
<hotSpot x="32" y="1" xunits="pixels" yunits="pixels"/>
</IconStyle>
</Style>
<Style id="statistics">
<IconStyle>
<scale>1.3</scale>
<Icon>
<href>http://maps.google.com/mapfiles/kml/pushpin/ylw-pushpin.png</href>
</Icon>
<hotSpot x="20" y="2" xunits="pixels" yunits="pixels"/>
</IconStyle>
</Style>
<Style id="waypoint">
<IconStyle>
<scale>1.3</scale>
<Icon>
<href>http://maps.google.com/mapfiles/kml/pushpin/blue-pushpin.png</href>
</Icon>
<hotSpot x="20" y="2" xunits="pixels" yunits="pixels"/>
</IconStyle>
</Style>
<Schema id="schema">
<gx:SimpleArrayField name="power" type="int">
<displayName>
<![CDATA[Power (W)]]>
</displayName>
</gx:SimpleArrayField>
<gx:SimpleArrayField name="cadence" type="int">
<displayName>
<![CDATA[Cadence (rpm)]]>
</displayName>
</gx:SimpleArrayField>
<gx:SimpleArrayField name="heart_rate" type="int">
<displayName>
<![CDATA[Heart rate (bpm)]]>
</displayName>
</gx:SimpleArrayField>
</Schema>
<Placemark>
<name>
<![CDATA[254-260 Main St E (Start)]]>
</name>
<description>
<![CDATA[]]>
</description>
<TimeStamp>
<when>2015-10-20T23:53:59.609Z</when>
</TimeStamp>
<styleUrl>#start</styleUrl>
<Point>
<coordinates>-111.891225,40.763656,1405.0</coordinates>
</Point>
</Placemark>
<Placemark id="tour">
<name>
<![CDATA[254-260 Main St E]]>
</name>
<description>
<![CDATA[]]>
</description>
<styleUrl>#track</styleUrl>
<ExtendedData>
<Data name="type">
<value>
<![CDATA[biking]]>
</value>
</Data>
</ExtendedData>
<gx:MultiTrack>
<altitudeMode>absolute</altitudeMode>
<gx:interpolate>1</gx:interpolate>
<gx:Track>
<when>2015-10-20T23:53:59.609Z</when>
<gx:coord>-111.891225 40.763656 1405.0</gx:coord>
<when>2015-10-20T23:55:35.984Z</when>
<gx:coord>-111.891225 40.763656 1405.0</gx:coord>
<when>2015-10-20T23:55:36.639Z</when>
<gx:coord>-111.891052 40.763578 1412.0</gx:coord>
<ExtendedData>
<SchemaData schemaUrl="#schema">
<gx:SimpleArrayData name="speed">
<gx:value>6.6754622</gx:value>
<gx:value>6.6754622</gx:value>
<gx:value>3.740762</gx:value>
</gx:SimpleArrayData>
<gx:SimpleArrayData name="bearing">
<gx:value>187.0</gx:value>
<gx:value>187.0</gx:value>
<gx:value>36.0</gx:value>
</gx:SimpleArrayData>
<gx:SimpleArrayData name="accuracy">
<gx:value>26.0</gx:value>
<gx:value>26.0</gx:value>
<gx:value>26.0</gx:value>
</gx:SimpleArrayData>
</SchemaData>
</ExtendedData>
</gx:Track>
</gx:MultiTrack>
</Placemark>
<Placemark>
<name>
<![CDATA[254-260 Main St E (End)]]>
</name>
<description>
<![CDATA[Created by Google My Tracks on Android

Name: 254-260 Main St E
Activity type: biking
Description: -
Total distance: 7.21 km (4.5 mi)
Total time: 27:31
Moving time: 25:03
Average speed: 15.72 km/h (9.8 mi/h)
Average moving speed: 17.27 km/h (10.7 mi/h)
Max speed: 41.72 km/h (25.9 mi/h)
Average pace: 3:49 min/km (6:09 min/mi)
Average moving pace: 3:28 min/km (5:35 min/mi)
Fastest pace: 1:26 min/km (2:19 min/mi)
Max elevation: 1441 m (4727 ft)
Min elevation: 1276 m (4186 ft)
Elevation gain: 155 m (508 ft)
Max grade: 19 %
Min grade: -37 %
Recorded: 10/20/2015 5:53 PM
]]>
</description>
<TimeStamp>
<when>2015-10-21T00:21:31.587Z</when>
</TimeStamp>
<styleUrl>#end</styleUrl>
<Point>
<coordinates>-111.891052 40.763578 1412.0</coordinates>
</Point>
</Placemark>
</Document>
</kml>
`
