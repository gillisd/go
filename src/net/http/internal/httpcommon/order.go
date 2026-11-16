// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpcommon

import (
	"slices"
	"strings"
)

// HeaderOrderKey is a magic header map key whose values list header
// names, in lowercase, in the order they should be written.
// It must match net/http.HeaderOrderKey.
const HeaderOrderKey = "Header-Order:"

// OrderedKeys returns the keys of h, excluding HeaderOrderKey.
//
// If h has a HeaderOrderKey entry, keys listed in it come first, in
// list order, followed by unlisted keys sorted lexicographically;
// this matches how net/http orders HTTP/1 headers. Keys are matched
// against the list after lowercasing.
//
// If h has no HeaderOrderKey entry, OrderedKeys returns nil, false.
func OrderedKeys(h map[string][]string) (keys []string, ok bool) {
	list, ok := h[HeaderOrderKey]
	if !ok {
		return nil, false
	}
	order := make(map[string]int, len(list))
	for i, v := range list {
		order[v] = i
	}
	keys = make([]string, 0, len(h))
	for k := range h {
		if k != HeaderOrderKey {
			keys = append(keys, k)
		}
	}
	slices.SortFunc(keys, func(a, b string) int {
		al, _ := LowerHeader(a)
		bl, _ := LowerHeader(b)
		ia, aok := order[al]
		ib, bok := order[bl]
		switch {
		case aok && bok:
			return ia - ib
		case aok:
			return -1
		case bok:
			return 1
		}
		return strings.Compare(a, b)
	})
	return keys, true
}
