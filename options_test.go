// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"github.com/eko/gocache/v2/store"
	"time"
)

func (t *StashTestSuite) TestOptions_ToStore() {
	o := Options{Expiration: time.Hour * 1}
	got := o.toStore()
	want := store.Options{Expiration: time.Hour * 1}
	if got == nil {
		t.Fail("shouldn't be nil")
		return
	}
	t.Equal(want, *got)
}

func (t *StashTestSuite) TestInvalidateOptions_ToStore() {
	o := InvalidateOptions{Tags: []string{"test"}}
	got := o.toStore()
	want := store.InvalidateOptions{Tags: []string{"test"}}
	t.Equal(want, got)
}
