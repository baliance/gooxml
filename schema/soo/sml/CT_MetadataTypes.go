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

type CT_MetadataTypes struct {
	// Metadata Type Count
	CountAttr *uint32
	// Metadata Type Information
	MetadataType []*CT_MetadataType
}

func NewCT_MetadataTypes() *CT_MetadataTypes {
	ret := &CT_MetadataTypes{}
	return ret
}

func (m *CT_MetadataTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	semetadataType := xml.StartElement{Name: xml.Name{Local: "ma:metadataType"}}
	for _, c := range m.MetadataType {
		e.EncodeElement(c, semetadataType)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MetadataTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_MetadataTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "metadataType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "metadataType"}:
				tmp := NewCT_MetadataType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.MetadataType = append(m.MetadataType, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_MetadataTypes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MetadataTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MetadataTypes and its children
func (m *CT_MetadataTypes) Validate() error {
	return m.ValidateWithPath("CT_MetadataTypes")
}

// ValidateWithPath validates the CT_MetadataTypes and its children, prefixing error messages with path
func (m *CT_MetadataTypes) ValidateWithPath(path string) error {
	for i, v := range m.MetadataType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/MetadataType[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
