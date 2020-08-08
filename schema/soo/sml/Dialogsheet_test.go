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

func TestDialogsheetConstructor(t *testing.T) {
	v := sml.NewDialogsheet()
	if v == nil {
		t.Errorf("sml.NewDialogsheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Dialogsheet should validate: %s", err)
	}
}

func TestDialogsheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewDialogsheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewDialogsheet()
	xml.Unmarshal(buf, v2)
}
