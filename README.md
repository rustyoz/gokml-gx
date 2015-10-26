GoKML GX
========

A KML library with [Google Extensions](https://developers.google.com/kml/documentation/kmlreference#kmlextensions) (http://www.google.com/kml/ext/2.2), that besides being able to encode/decode to XML.

This library is based on the good work done by [dahenson](https://github.com/dahenson/gokml)

Overview
--------

Go KML GX was written to be able to read Google My Tracks data. Google My Tracks uses Google's KML extensions in it's documents. This implementation has a focus on the extension elements that is used by My Tracks.

This library can also take in or produce a KMZ (zipped KML) file in. It can unzip the KMZ file and produce a set of Go structs from the KML file inside the archive, or in reverse create a KMZ file.

Status
------

This is currently a work in progress. Not all parts of Google's KML extensions are implemented, and not all of My Tracks elements are implemented.
