// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/chartDrawing"
)

func TestEG_ObjectChoicesConstructor(t *testing.T) {
	v := chartDrawing.NewEG_ObjectChoices()
	if v == nil {
		t.Errorf("chartDrawing.NewEG_ObjectChoices must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.EG_ObjectChoices should validate: %s", err)
	}
}

func TestEG_ObjectChoicesMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewEG_ObjectChoices()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewEG_ObjectChoices()
	xml.Unmarshal(buf, v2)
}
