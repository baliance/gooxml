// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_BuiltInUnit struct {
	ValAttr ST_BuiltInUnit
}

func NewCT_BuiltInUnit() *CT_BuiltInUnit {
	ret := &CT_BuiltInUnit{}
	return ret
}

func (m *CT_BuiltInUnit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_BuiltInUnitUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BuiltInUnit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_BuiltInUnit: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_BuiltInUnit and its children
func (m *CT_BuiltInUnit) Validate() error {
	return m.ValidateWithPath("CT_BuiltInUnit")
}

// ValidateWithPath validates the CT_BuiltInUnit and its children, prefixing error messages with path
func (m *CT_BuiltInUnit) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
