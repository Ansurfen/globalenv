//go:build windows
// +build windows

// Copyright 2025 The globalenv Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package globalenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQuery(t *testing.T) {
	testset := []struct {
		input    string
		expected string
	}{
		{input: `
HKEY_CURRENT_USER\Environment
    globalEnv    REG_SZ    Hello Global Environment!


`, expected: "Hello Global Environment!"},
	}

	for _, ts := range testset {
		in, err := parseQuery([]byte(ts.input), nil)
		assert.NoError(t, err)
		assert.Equal(t, ts.expected, string(in))
	}
}
