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
)

type CT_CellMergeTrackChange struct {
	VMergeAttr     ST_AnnotationVMerge
	VMergeOrigAttr ST_AnnotationVMerge
	AuthorAttr     string
	DateAttr       *time.Time
	// Annotation Identifier
	IdAttr int64
}

func NewCT_CellMergeTrackChange() *CT_CellMergeTrackChange {
	ret := &CT_CellMergeTrackChange{}
	return ret
}

func (m *CT_CellMergeTrackChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.VMergeAttr != ST_AnnotationVMergeUnset {
		attr, err := m.VMergeAttr.MarshalXMLAttr(xml.Name{Local: "w:vMerge"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VMergeOrigAttr != ST_AnnotationVMergeUnset {
		attr, err := m.VMergeOrigAttr.MarshalXMLAttr(xml.Name{Local: "w:vMergeOrig"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellMergeTrackChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "vMerge" {
			m.VMergeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "vMergeOrig" {
			m.VMergeOrigAttr.UnmarshalXMLAttr(attr)
			continue
		}
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
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CellMergeTrackChange: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CellMergeTrackChange and its children
func (m *CT_CellMergeTrackChange) Validate() error {
	return m.ValidateWithPath("CT_CellMergeTrackChange")
}

// ValidateWithPath validates the CT_CellMergeTrackChange and its children, prefixing error messages with path
func (m *CT_CellMergeTrackChange) ValidateWithPath(path string) error {
	if err := m.VMergeAttr.ValidateWithPath(path + "/VMergeAttr"); err != nil {
		return err
	}
	if err := m.VMergeOrigAttr.ValidateWithPath(path + "/VMergeOrigAttr"); err != nil {
		return err
	}
	return nil
}
