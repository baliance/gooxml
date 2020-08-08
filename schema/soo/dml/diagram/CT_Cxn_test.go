// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/diagram"
)

func TestCT_CxnConstructor(t *testing.T) {
	v := diagram.NewCT_Cxn()
	if v == nil {
		t.Errorf("diagram.NewCT_Cxn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Cxn should validate: %s", err)
	}
}

func TestCT_CxnMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Cxn()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Cxn()
	xml.Unmarshal(buf, v2)
}
