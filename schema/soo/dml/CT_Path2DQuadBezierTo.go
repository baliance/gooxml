// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Path2DQuadBezierTo struct {
	Pt []*CT_AdjPoint2D
}

func NewCT_Path2DQuadBezierTo() *CT_Path2DQuadBezierTo {
	ret := &CT_Path2DQuadBezierTo{}
	return ret
}

func (m *CT_Path2DQuadBezierTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sept := xml.StartElement{Name: xml.Name{Local: "a:pt"}}
	for _, c := range m.Pt {
		e.EncodeElement(c, sept)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Path2DQuadBezierTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Path2DQuadBezierTo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pt"}:
				tmp := NewCT_AdjPoint2D()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pt = append(m.Pt, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Path2DQuadBezierTo %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path2DQuadBezierTo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Path2DQuadBezierTo and its children
func (m *CT_Path2DQuadBezierTo) Validate() error {
	return m.ValidateWithPath("CT_Path2DQuadBezierTo")
}

// ValidateWithPath validates the CT_Path2DQuadBezierTo and its children, prefixing error messages with path
func (m *CT_Path2DQuadBezierTo) ValidateWithPath(path string) error {
	for i, v := range m.Pt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
