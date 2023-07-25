// Copyright 2021 The golang.design Initiative Authors, 2023 Gabriele V.
// All rights reserved. Use of this source code is governed
// by a MIT license that can be found in the LICENSE file.
//
// Original code written by Changkun Ou <changkun.de>
// Modified to handle Cacoo web MIME type by Gabriele V. <gabriele@nulab.com>

package cacooclip

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/vibridi/cacooclip/internal/chromium"
)

const (
	CacooWebMIMEType = "cacoo/shape"
)

var (
	errUnavailable = errors.New("unavailable or no data from clipboard")
	errNoCacooMime = errors.New("no content with Cacoo MIME type in clipboard")
)

var (
	// Due to the limitation on operating systems (such as darwin),
	// concurrent read can even cause panic, use a global lock to
	// guarantee one read at a time.
	lock = sync.Mutex{}
)

func Read() (string, error) {
	lock.Lock()
	defer lock.Unlock()

	b, err := read()
	if err != nil {
		return "", fmt.Errorf("read clipboard err: %w", err)
	}
	wcmap, err := chromium.ReadWebCustomMIMEData(b)
	if err != nil {
		return "", fmt.Errorf("decode err: %w", err)
	}
	content, ok := wcmap[CacooWebMIMEType]
	if !ok {
		return "", errNoCacooMime
	}
	return content, nil
}

func Write(buf []byte) <-chan struct{} {
	lock.Lock()
	defer lock.Unlock()

	changed, err := write(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write to clipboard err: %v\n", err)
		return nil
	}
	return changed
}
