// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/unidoc/unioffice"
)

type CT_ParaRPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	RPr    *CT_ParaRPrOriginal
}

func NewCT_ParaRPrChange() *CT_ParaRPrChange {
	ret := &CT_ParaRPrChange{}
	ret.RPr = NewCT_ParaRPrOriginal()
	return ret
}

func (m *CT_ParaRPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
	e.EncodeElement(m.RPr, serPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ParaRPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RPr = NewCT_ParaRPrOriginal()
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_ParaRPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPr"}:
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ParaRPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ParaRPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ParaRPrChange and its children
func (m *CT_ParaRPrChange) Validate() error {
	return m.ValidateWithPath("CT_ParaRPrChange")
}

// ValidateWithPath validates the CT_ParaRPrChange and its children, prefixing error messages with path
func (m *CT_ParaRPrChange) ValidateWithPath(path string) error {
	if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
		return err
	}
	return nil
}
