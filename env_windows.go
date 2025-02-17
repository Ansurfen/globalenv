//go:build windows
// +build windows

// Copyright 2025 The globalenv Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package globalenv

import (
	"strings"
)

// Get retrieves the environment variable value from the user-specific registry.
func Get(key string) (string, error) {
	return EnvInFile("user").Get(key)
}

// Set sets the value of the environment variable in the user-specific registry.
func Set(key, value string) ([]byte, error) {
	return EnvInFile("user").Set(key, value)
}

// Unset removes the environment variable from the user-specific registry.
func Unset(key string) ([]byte, error) {
	return EnvInFile("user").Unset(key)
}

// EnvInFile represents a configuration (user or system) for environment variables.
type EnvInFile string

// Get retrieves the environment variable from the registry based on the configuration (user/system).
func (e EnvInFile) Get(key string) (string, error) {
	reg := regBuilder{path: regPathUser, val: key}
	switch {
	case len(e) == 0 || e == "user":
		return parseQuery(run("reg", reg.Query()...))
	case e == "system":
		reg.path = regPathSys
		return parseQuery(run("reg", reg.Query()...))
	}
	return "", ErrInvalidEnvInFile
}

// Set sets the environment variable in the registry (user/system).
func (e EnvInFile) Set(key, value string) ([]byte, error) {
	switch {
	case len(e) == 0 || e == "user":
		return run("setx", key, value)
	case e == "system":
		return run("setx", key, value, "/M")
	}
	return nil, ErrInvalidEnvInFile
}

// Unset removes the environment variable from the registry (user/system).
func (e EnvInFile) Unset(key string) ([]byte, error) {
	reg := regBuilder{path: regPathUser, val: key}
	switch {
	case len(e) == 0 || e == "user":
		return run("reg", reg.Delete()...)
	case e == "system":
		reg.path = regPathSys
		return run("reg", reg.Delete()...)
	}
	return nil, ErrInvalidEnvInFile
}

// InFile returns an EnvInFile instance based on the specified environment configuration (user/system).
func InFile(name string) EnvInFile {
	return EnvInFile(name)
}

// regPath represents the registry path for either the user or system environment variables.
type regPath int

// String returns the registry path for either user or system environment variables.
func (p regPath) String() string {
	return []string{
		`HKCU\Environment`,
		`HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`,
	}[p]
}

// Constants for user and system registry paths.
const (
	regPathUser regPath = iota
	regPathSys
)

// regBuilder builds the registry command for querying or deleting environment variables.
type regBuilder struct {
	path regPath
	val  string
}

// Query returns the command to query the registry for the environment variable.
func (r regBuilder) Query() []string {
	return []string{"query", r.path.String(), "/v", r.val}
}

// Delete returns the command to delete the environment variable from the registry.
func (r regBuilder) Delete() []string {
	return []string{"delete", r.path.String(), "/F", "/V", r.val}
}

// parseQuery parses the output of the reg query command to extract the value.
func parseQuery(output []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	lines := strings.SplitN(string(output), "\n", 3)
	if len(lines) == 3 {
		lines[2] = strings.TrimSpace(lines[2])
		lines = strings.SplitN(lines[2], "    ", 3)
		if len(lines) == 3 {
			return lines[2], nil
		}
	}
	return "", ErrFailParseValue
}
