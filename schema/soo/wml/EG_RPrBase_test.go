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

func TestEG_RPrBaseConstructor(t *testing.T) {
	v := wml.NewEG_RPrBase()
	if v == nil {
		t.Errorf("wml.NewEG_RPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RPrBase should validate: %s", err)
	}
}

func TestEG_RPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RPrBase()
	xml.Unmarshal(buf, v2)
}
