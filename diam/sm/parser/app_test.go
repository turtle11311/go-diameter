// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package parser

import (
	"testing"

	"github.com/fiorix/go-diameter/diam"
	"github.com/fiorix/go-diameter/diam/avp"
	"github.com/fiorix/go-diameter/diam/datatype"
	"github.com/fiorix/go-diameter/diam/dict"
)

func TestUnexpectedAVP_BadCode(t *testing.T) {
	a := diam.NewAVP(0, 0, 0, datatype.Unsigned32(0))
	app := &Application{
		AcctApplicationID: []*diam.AVP{a},
	}
	failedAVP, err := app.Parse(dict.Default)
	if err == nil {
		t.Fatal("Unexpected application parsed successfully")
	}
	if failedAVP != a {
		t.Fatalf("Unexpected failed avp. Want %q, have %q", a, failedAVP)
	}
}

func TestUnexpectedAVP_BadData(t *testing.T) {
	a := diam.NewAVP(avp.AcctApplicationID, 0, 0, datatype.OctetString("x"))
	app := &Application{
		AcctApplicationID: []*diam.AVP{a},
	}
	failedAVP, err := app.Parse(dict.Default)
	if err == nil {
		t.Fatal("Unexpected application parsed successfully")
	}
	if failedAVP != a {
		t.Fatalf("Unexpected failed avp. Want %q, have %q", a, failedAVP)
	}
}

func TestUnexpectedAVP_BadGroup(t *testing.T) {
	a := diam.NewAVP(avp.AcctApplicationID, 0, 0, datatype.Unsigned32(0))
	app := &Application{
		VendorSpecificApplicationID: []*diam.AVP{a},
	}
	failedAVP, err := app.Parse(dict.Default)
	if err == nil {
		t.Fatal("Unexpected application parsed successfully")
	}
	if failedAVP != a {
		t.Fatalf("Unexpected failed avp. Want %q, have %q", a, failedAVP)
	}
}
