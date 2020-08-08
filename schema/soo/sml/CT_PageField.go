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
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice"
)

type CT_PageField struct {
	// Field
	FldAttr int32
	// Item Index
	ItemAttr *uint32
	// OLAP Hierarchy Index
	HierAttr *int32
	// Hierarchy Unique Name
	NameAttr *string
	// Hierarchy Display Name
	CapAttr *string
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PageField() *CT_PageField {
	ret := &CT_PageField{}
	return ret
}

func (m *CT_PageField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fld"},
		Value: fmt.Sprintf("%v", m.FldAttr)})
	if m.ItemAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "item"},
			Value: fmt.Sprintf("%v", *m.ItemAttr)})
	}
	if m.HierAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hier"},
			Value: fmt.Sprintf("%v", *m.HierAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.CapAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cap"},
			Value: fmt.Sprintf("%v", *m.CapAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fld" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FldAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "item" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ItemAttr = &pt
			continue
		}
		if attr.Name.Local == "hier" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.HierAttr = &pt
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "cap" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CapAttr = &parsed
			continue
		}
	}
lCT_PageField:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_PageField %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PageField
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PageField and its children
func (m *CT_PageField) Validate() error {
	return m.ValidateWithPath("CT_PageField")
}

// ValidateWithPath validates the CT_PageField and its children, prefixing error messages with path
func (m *CT_PageField) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
