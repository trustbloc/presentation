/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mem

import (
	"errors"
	"github.com/trustbloc/presentation/2020/01/aries-framework-go/cmd/lobby/store"
	"sync"
)

type memStore struct {
	db map[string][]byte
	sync.RWMutex
}

func New() store.Store{
	return &memStore{
		db:      make(map[string][]byte),
		RWMutex: sync.RWMutex{},
	}
}

func (s *memStore) Get(k string) ([]byte, error) {
	if k == "" {
		return nil, errors.New("key is mandatory")
	}

	s.RLock()
	data, ok := s.db[k]
	s.RUnlock()

	if !ok {
		return nil, store.ErrDataNotFound
	}

	return data, nil
}

func (s *memStore) Put(k string, v []byte) error {
	if k == "" || v == nil {
		return errors.New("key and value are mandatory")
	}

	s.Lock()
	s.db[k] = v
	s.Unlock()

	return nil
}

func (s *memStore) Remove(k string) error {
	return nil
}
