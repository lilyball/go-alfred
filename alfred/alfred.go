package alfred

import (
	"encoding/xml"
)

const (
	TypeFile = "file"
)

const (
	IconTypePath     = "path"
	IconTypeFileIcon = "fileicon"
	IconTypeFileType = "filetype"
)

type Icon struct {
	Type  string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Item struct {
	// Attributes
	Uid          string `xml:"uid,attr,omitempty"`
	Arg          string `xml:"arg,attr,omitempty"`
	Valid        bool   `xml:"-"`
	Autocomplete string `xml:"autocomplete,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`

	// Sub-elements
	Title    string `xml:"title"`
	Subtitle string `xml:"subtitle,omitempty"`
	Icon     Icon `xml:"-"`
}

type item struct {
	Item
	XMLValid string `xml:"valid,attr"`
	Icon     *Icon  `xml:"icon,omitempty"`
}

func Encode(items []Item) ([]byte, error) {
	xmlitems := make([]item, len(items))
	for i, item := range items {
		xmlitems[i].Item = item
		if item.Valid {
			xmlitems[i].XMLValid = "YES"
		} else {
			xmlitems[i].XMLValid = "NO"
		}
		if item.Icon.Type != "" {
			xmlitems[i].Icon = &item.Icon
			if xmlitems[i].Icon.Type == IconTypePath {
				xmlitems[i].Icon.Type = ""
			}
		}
	}

	bytes, err := xml.MarshalIndent(xmlitems, "    ", "    ")
	if err == nil {
		output := append([]byte(xml.Header), []byte("<items>\n")...)
		if bytes[0] == '\n' {
			bytes = bytes[1:]
		}
		output = append(output, bytes...)
		output = append(output, []byte("\n</items>\n")...)
		bytes = output
	}
	return bytes, err
}
