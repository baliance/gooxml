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

type CT_ChartFormats struct {
	// Format Count
	CountAttr *uint32
	// PivotChart Format
	ChartFormat []*CT_ChartFormat
}

func NewCT_ChartFormats() *CT_ChartFormats {
	ret := &CT_ChartFormats{}
	return ret
}

func (m *CT_ChartFormats) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sechartFormat := xml.StartElement{Name: xml.Name{Local: "ma:chartFormat"}}
	for _, c := range m.ChartFormat {
		e.EncodeElement(c, sechartFormat)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartFormats) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_ChartFormats:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "chartFormat"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "chartFormat"}:
				tmp := NewCT_ChartFormat()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ChartFormat = append(m.ChartFormat, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_ChartFormats %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartFormats
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartFormats and its children
func (m *CT_ChartFormats) Validate() error {
	return m.ValidateWithPath("CT_ChartFormats")
}

// ValidateWithPath validates the CT_ChartFormats and its children, prefixing error messages with path
func (m *CT_ChartFormats) ValidateWithPath(path string) error {
	for i, v := range m.ChartFormat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ChartFormat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
