//go:build !windows
// +build !windows

// Copyright 2025 The globalenv Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package globalenv

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Get retrieves the value of a specific environment variable from a file (e.g., ~/.bashrc)
func Get(key string) (string, error) {
	return EnvInFile("~/.bashrc").Get(key)
}

// Set sets the value of a specific environment variable in a file (e.g., ~/.bashrc) and sources the file.
// Example: ~/.bash_profile ~/.profile ~/.bashrc ~/.zshrc /etc/profile /etc/bashrc /etc/bash.bashrc /etc/environment
func Set(key, value string) ([]byte, error) {
	return EnvInFile("~/.bashrc").Set(key, value)
}

// Unset removes a specific environment variable from the file (e.g., ~/.bashrc)
func Unset(key string) ([]byte, error) {
	return EnvInFile("~/.bashrc").Unset(key)
}

// EnvInFile represents a path to a shell configuration file (e.g., ~/.bashrc)
type EnvInFile string

// Get retrieves the value of a specific environment variable from the file. It reads the file, then queries the key.
func (e EnvInFile) Get(key string) (string, error) {
	path, err := e.pathResolve()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return query(data, key)
}

// Set appends a new environment variable definition (e.g., export KEY=value) to the file
// and sources the file to apply changes.
func (e EnvInFile) Set(key, value string) ([]byte, error) {
	if len(e) == 0 {
		return run("bash", "-c", fmt.Sprintf(`echo 'export %s="%s"' >> ~/.bashrc && source ~/.bashrc`, key, value))
	}
	return run("bash", "-c", fmt.Sprintf(`echo 'export %s="%s"' >> %s && source %s`, key, value, e, e))
}

// Unset removes the environment variable definition from the file by matching the key
// and writing the modified content back to the file.
func (e EnvInFile) Unset(key string) ([]byte, error) {
	path, err := e.pathResolve()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data, err = unset(data, key)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// pathResolve resolves the file path, expanding "~" to the user's home directory if necessary.
func (e EnvInFile) pathResolve() (string, error) {
	if len(e) == 0 {
		return "", nil
	}
	if len(e) > 1 && e[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return strings.Replace(string(e), "~", homeDir, 1), nil
	}
	return string(e), nil
}

// InFile creates an EnvInFile instance with the given file name (e.g., ~/.bashrc).
func InFile(name string) EnvInFile {
	return EnvInFile(name)
}

// query searches for the specified environment variable key in the given content (file content).
// It assumes the format of the variable is `export KEY=value`.
//
// NOTICE: the format as follows isn't support
//
// `export PATH="$PATH:$HOME/scripts"; export EDITOR=nano; export EDITOR=nano`
func query(content []byte, key string) (string, error) {
	pattern := fmt.Sprintf(`(?m)^export\s+%s=([^\n]*)`, key)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	matches := re.FindAllStringSubmatch(string(content), -1)
	if len(matches) == 0 {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	appendkey := fmt.Sprintf("$%s:", key)
	ret := []string{}
	for _, m := range matches {
		if len(m) > 1 {
			s := m[1]
			if s[0] == '"' && s[len(s)-1] == '"' {
				s = s[1 : len(s)-1]
			}
			s = strings.Replace(s, appendkey, "", 1)
			fmt.Println(s)
			strings.Split(s, ";")
			ret = append(ret, strings.Replace(s, appendkey, "", 1))
		}
	}

	return strings.Join(ret, ";"), nil
}

// unset removes the specified key-value pair from the content of the file.
func unset(content []byte, key string) ([]byte, error) {
	pattern := fmt.Sprintf(`(?m)^export\s+%s=.*?(\s*;\s*[^$])?\n`, key)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	newContent := re.ReplaceAllLiteral(content, []byte(""))
	return newContent, nil
}
