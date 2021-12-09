// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalUintptrTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uintptr": pfmt.Uintptr(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uintptr":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uintp": pfmt.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uintp":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uintp": pfmt.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uintp":42
		}`,
	},
}

func TestMarshalUintptr(t *testing.T) {
	testMarshal(t, MarshalUintptrTests)
}
