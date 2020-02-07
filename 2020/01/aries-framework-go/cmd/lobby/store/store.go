/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package store

import "errors"

var ErrDataNotFound = errors.New("data not found")

type Store interface {
	Get(k string) ([]byte, error)

	Put(k string, v []byte) error

	Remove(k string) error
}
