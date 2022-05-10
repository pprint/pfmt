// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalInt8(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int8": pfmt.Int8(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int8":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int8": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int8":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int8": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int8":42
		}`,
		},
	}

	testMarshal(t, tests)
}
