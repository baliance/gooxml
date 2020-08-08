// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/ofc/math"
)

func TestCT_BarPrConstructor(t *testing.T) {
	v := math.NewCT_BarPr()
	if v == nil {
		t.Errorf("math.NewCT_BarPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BarPr should validate: %s", err)
	}
}

func TestCT_BarPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BarPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BarPr()
	xml.Unmarshal(buf, v2)
}
