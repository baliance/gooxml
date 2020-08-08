// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestCT_ZoomConstructor(t *testing.T) {
	v := wml.NewCT_Zoom()
	if v == nil {
		t.Errorf("wml.NewCT_Zoom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Zoom should validate: %s", err)
	}
}

func TestCT_ZoomMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Zoom()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Zoom()
	xml.Unmarshal(buf, v2)
}
