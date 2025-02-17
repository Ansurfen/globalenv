// Copyright 2025 The globalenv Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package globalenv

import (
	"errors"
	"os/exec"
)

func run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

var (
	ErrInvalidEnvInFile = errors.New("invalid EnvInFile")
	ErrFailParseValue   = errors.New("fail to parse value")
)
