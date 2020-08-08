// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_BorderPr struct {
	// Line Style
	StyleAttr ST_BorderStyle
	// Color
	Color *CT_Color
}

func NewCT_BorderPr() *CT_BorderPr {
	ret := &CT_BorderPr{}
	return ret
}

func (m *CT_BorderPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StyleAttr != ST_BorderStyleUnset {
		attr, err := m.StyleAttr.MarshalXMLAttr(xml.Name{Local: "style"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Color != nil {
		secolor := xml.StartElement{Name: xml.Name{Local: "ma:color"}}
		e.EncodeElement(m.Color, secolor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BorderPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "style" {
			m.StyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_BorderPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "color"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "color"}:
				m.Color = NewCT_Color()
				if err := d.DecodeElement(m.Color, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_BorderPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BorderPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BorderPr and its children
func (m *CT_BorderPr) Validate() error {
	return m.ValidateWithPath("CT_BorderPr")
}

// ValidateWithPath validates the CT_BorderPr and its children, prefixing error messages with path
func (m *CT_BorderPr) ValidateWithPath(path string) error {
	if err := m.StyleAttr.ValidateWithPath(path + "/StyleAttr"); err != nil {
		return err
	}
	if m.Color != nil {
		if err := m.Color.ValidateWithPath(path + "/Color"); err != nil {
			return err
		}
	}
	return nil
}
