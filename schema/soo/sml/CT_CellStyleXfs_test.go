// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/sml"
)

func TestCT_CellStyleXfsConstructor(t *testing.T) {
	v := sml.NewCT_CellStyleXfs()
	if v == nil {
		t.Errorf("sml.NewCT_CellStyleXfs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellStyleXfs should validate: %s", err)
	}
}

func TestCT_CellStyleXfsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellStyleXfs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellStyleXfs()
	xml.Unmarshal(buf, v2)
}
