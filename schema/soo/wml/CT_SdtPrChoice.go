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

type CT_SdtPrChoice struct {
	Equation     *CT_Empty
	ComboBox     *CT_SdtComboBox
	Date         *CT_SdtDate
	DocPartObj   *CT_SdtDocPart
	DocPartList  *CT_SdtDocPart
	DropDownList *CT_SdtDropDownList
	Picture      *CT_Empty
	RichText     *CT_Empty
	Text         *CT_SdtText
	Citation     *CT_Empty
	Group        *CT_Empty
	Bibliography *CT_Empty
}

func NewCT_SdtPrChoice() *CT_SdtPrChoice {
	ret := &CT_SdtPrChoice{}
	return ret
}

func (m *CT_SdtPrChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Equation != nil {
		seequation := xml.StartElement{Name: xml.Name{Local: "w:equation"}}
		e.EncodeElement(m.Equation, seequation)
	}
	if m.ComboBox != nil {
		secomboBox := xml.StartElement{Name: xml.Name{Local: "w:comboBox"}}
		e.EncodeElement(m.ComboBox, secomboBox)
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "w:date"}}
		e.EncodeElement(m.Date, sedate)
	}
	if m.DocPartObj != nil {
		sedocPartObj := xml.StartElement{Name: xml.Name{Local: "w:docPartObj"}}
		e.EncodeElement(m.DocPartObj, sedocPartObj)
	}
	if m.DocPartList != nil {
		sedocPartList := xml.StartElement{Name: xml.Name{Local: "w:docPartList"}}
		e.EncodeElement(m.DocPartList, sedocPartList)
	}
	if m.DropDownList != nil {
		sedropDownList := xml.StartElement{Name: xml.Name{Local: "w:dropDownList"}}
		e.EncodeElement(m.DropDownList, sedropDownList)
	}
	if m.Picture != nil {
		sepicture := xml.StartElement{Name: xml.Name{Local: "w:picture"}}
		e.EncodeElement(m.Picture, sepicture)
	}
	if m.RichText != nil {
		serichText := xml.StartElement{Name: xml.Name{Local: "w:richText"}}
		e.EncodeElement(m.RichText, serichText)
	}
	if m.Text != nil {
		setext := xml.StartElement{Name: xml.Name{Local: "w:text"}}
		e.EncodeElement(m.Text, setext)
	}
	if m.Citation != nil {
		secitation := xml.StartElement{Name: xml.Name{Local: "w:citation"}}
		e.EncodeElement(m.Citation, secitation)
	}
	if m.Group != nil {
		segroup := xml.StartElement{Name: xml.Name{Local: "w:group"}}
		e.EncodeElement(m.Group, segroup)
	}
	if m.Bibliography != nil {
		sebibliography := xml.StartElement{Name: xml.Name{Local: "w:bibliography"}}
		e.EncodeElement(m.Bibliography, sebibliography)
	}
	return nil
}

func (m *CT_SdtPrChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtPrChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "equation"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "equation"}:
				m.Equation = NewCT_Empty()
				if err := d.DecodeElement(m.Equation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "comboBox"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "comboBox"}:
				m.ComboBox = NewCT_SdtComboBox()
				if err := d.DecodeElement(m.ComboBox, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "date"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "date"}:
				m.Date = NewCT_SdtDate()
				if err := d.DecodeElement(m.Date, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartObj"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartObj"}:
				m.DocPartObj = NewCT_SdtDocPart()
				if err := d.DecodeElement(m.DocPartObj, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartList"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartList"}:
				m.DocPartList = NewCT_SdtDocPart()
				if err := d.DecodeElement(m.DocPartList, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dropDownList"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dropDownList"}:
				m.DropDownList = NewCT_SdtDropDownList()
				if err := d.DecodeElement(m.DropDownList, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "picture"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "picture"}:
				m.Picture = NewCT_Empty()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "richText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "richText"}:
				m.RichText = NewCT_Empty()
				if err := d.DecodeElement(m.RichText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "text"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "text"}:
				m.Text = NewCT_SdtText()
				if err := d.DecodeElement(m.Text, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "citation"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "citation"}:
				m.Citation = NewCT_Empty()
				if err := d.DecodeElement(m.Citation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "group"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "group"}:
				m.Group = NewCT_Empty()
				if err := d.DecodeElement(m.Group, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bibliography"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bibliography"}:
				m.Bibliography = NewCT_Empty()
				if err := d.DecodeElement(m.Bibliography, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SdtPrChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtPrChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtPrChoice and its children
func (m *CT_SdtPrChoice) Validate() error {
	return m.ValidateWithPath("CT_SdtPrChoice")
}

// ValidateWithPath validates the CT_SdtPrChoice and its children, prefixing error messages with path
func (m *CT_SdtPrChoice) ValidateWithPath(path string) error {
	if m.Equation != nil {
		if err := m.Equation.ValidateWithPath(path + "/Equation"); err != nil {
			return err
		}
	}
	if m.ComboBox != nil {
		if err := m.ComboBox.ValidateWithPath(path + "/ComboBox"); err != nil {
			return err
		}
	}
	if m.Date != nil {
		if err := m.Date.ValidateWithPath(path + "/Date"); err != nil {
			return err
		}
	}
	if m.DocPartObj != nil {
		if err := m.DocPartObj.ValidateWithPath(path + "/DocPartObj"); err != nil {
			return err
		}
	}
	if m.DocPartList != nil {
		if err := m.DocPartList.ValidateWithPath(path + "/DocPartList"); err != nil {
			return err
		}
	}
	if m.DropDownList != nil {
		if err := m.DropDownList.ValidateWithPath(path + "/DropDownList"); err != nil {
			return err
		}
	}
	if m.Picture != nil {
		if err := m.Picture.ValidateWithPath(path + "/Picture"); err != nil {
			return err
		}
	}
	if m.RichText != nil {
		if err := m.RichText.ValidateWithPath(path + "/RichText"); err != nil {
			return err
		}
	}
	if m.Text != nil {
		if err := m.Text.ValidateWithPath(path + "/Text"); err != nil {
			return err
		}
	}
	if m.Citation != nil {
		if err := m.Citation.ValidateWithPath(path + "/Citation"); err != nil {
			return err
		}
	}
	if m.Group != nil {
		if err := m.Group.ValidateWithPath(path + "/Group"); err != nil {
			return err
		}
	}
	if m.Bibliography != nil {
		if err := m.Bibliography.ValidateWithPath(path + "/Bibliography"); err != nil {
			return err
		}
	}
	return nil
}
