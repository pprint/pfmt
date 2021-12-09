// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalIntpsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int = 42, 77
			return map[string]json.Marshaler{"int pointer slice": pfmt.Intps(&f, &f2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"int pointer slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil int pointers": pfmt.Intps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil int pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int pointers": pfmt.Intps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without int pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int = 42, 77
			return map[string]json.Marshaler{"slice of any int pointers": pfmt.Anys(&f, &f2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any int pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int = 42, 77
			return map[string]json.Marshaler{"slice of reflects of int pointers": pfmt.Reflects(&f, &f2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of reflects of int pointers":[42,77]
		}`,
	},
}

func TestMarshalIntps(t *testing.T) {
	testMarshal(t, MarshalIntpsTests)
}