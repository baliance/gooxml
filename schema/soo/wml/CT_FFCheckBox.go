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

	"github.com/unidoc/unioffice"
)

type CT_FFCheckBox struct {
	Choice *CT_FFCheckBoxChoice
	// Default Checkbox Form Field State
	Default *CT_OnOff
	// Checkbox Form Field State
	Checked *CT_OnOff
}

func NewCT_FFCheckBox() *CT_FFCheckBox {
	ret := &CT_FFCheckBox{}
	return ret
}

func (m *CT_FFCheckBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.Default != nil {
		sedefault := xml.StartElement{Name: xml.Name{Local: "w:default"}}
		e.EncodeElement(m.Default, sedefault)
	}
	if m.Checked != nil {
		sechecked := xml.StartElement{Name: xml.Name{Local: "w:checked"}}
		e.EncodeElement(m.Checked, sechecked)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FFCheckBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFCheckBox:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "size"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "size"}:
				m.Choice = NewCT_FFCheckBoxChoice()
				if err := d.DecodeElement(&m.Choice.Size, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sizeAuto"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sizeAuto"}:
				m.Choice = NewCT_FFCheckBoxChoice()
				if err := d.DecodeElement(&m.Choice.SizeAuto, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "default"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "default"}:
				m.Default = NewCT_OnOff()
				if err := d.DecodeElement(m.Default, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "checked"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "checked"}:
				m.Checked = NewCT_OnOff()
				if err := d.DecodeElement(m.Checked, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_FFCheckBox %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFCheckBox
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FFCheckBox and its children
func (m *CT_FFCheckBox) Validate() error {
	return m.ValidateWithPath("CT_FFCheckBox")
}

// ValidateWithPath validates the CT_FFCheckBox and its children, prefixing error messages with path
func (m *CT_FFCheckBox) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.Default != nil {
		if err := m.Default.ValidateWithPath(path + "/Default"); err != nil {
			return err
		}
	}
	if m.Checked != nil {
		if err := m.Checked.ValidateWithPath(path + "/Checked"); err != nil {
			return err
		}
	}
	return nil
}
