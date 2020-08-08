// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestHandoutMasterConstructor(t *testing.T) {
	v := pml.NewHandoutMaster()
	if v == nil {
		t.Errorf("pml.NewHandoutMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.HandoutMaster should validate: %s", err)
	}
}

func TestHandoutMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewHandoutMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewHandoutMaster()
	xml.Unmarshal(buf, v2)
}
