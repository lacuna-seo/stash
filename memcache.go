// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"time"
)

// memcacheStore defines the data stored for the memcache
// client.
type memcacheStore struct {
	client            *memcache.Client
	servers           []string
	defaultExpiration time.Duration
}

// NewMemcache creates a new memcached store and returns a provider.
func NewMemcache(servers []string, defaultExpiration time.Duration) Provider {
	return &memcacheStore{
		client:            memcache.New(servers...),
		servers:           servers,
		defaultExpiration: defaultExpiration,
	}
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (m *memcacheStore) Validate() error {
	if len(m.servers) == 0 {
		return errors.New("no memcache hosts defined in env")
	}
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory Driver name.
func (m *memcacheStore) Driver() string {
	return MemcacheDriver
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (m *memcacheStore) Store() store.StoreInterface {
	return cache.New(store.NewMemcache(m.client, &store.Options{
		Expiration: m.defaultExpiration,
	}))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (m *memcacheStore) Ping() error {
	return m.client.Ping()
}
