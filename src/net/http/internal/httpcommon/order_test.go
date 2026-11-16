// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpcommon

import (
	"slices"
	"testing"
)

func TestOrderedKeys(t *testing.T) {
	h := map[string][]string{
		"Zz":           {"1"},
		"Aa":           {"2"},
		"Mm":           {"3"},
		"Bb":           {"4"},
		"Yy":           {"5"},
		HeaderOrderKey: {"mm", "zz", "aa"},
	}
	got, ok := OrderedKeys(h)
	want := []string{"Mm", "Zz", "Aa", "Bb", "Yy"}
	if !ok || !slices.Equal(got, want) {
		t.Errorf("OrderedKeys = %q, %v; want %q, true", got, ok, want)
	}

	delete(h, HeaderOrderKey)
	if got, ok := OrderedKeys(h); ok || got != nil {
		t.Errorf("OrderedKeys without order = %q, %v; want nil, false", got, ok)
	}
}
