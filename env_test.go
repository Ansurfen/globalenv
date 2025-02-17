// Copyright 2025 The globalenv Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package globalenv

import (
	"fmt"
	"testing"

	"github.com/djimenez/iconv-go"
	"github.com/stretchr/testify/assert"
)

func TestSetGlobalEnv(t *testing.T) {
	output, err := Set("globalEnv", "Hello Global Environment!")
	assert.NoError(t, err, string(output))

	output2, err := EnvInFile("user").Get("globalEnv")
	assert.NoError(t, err)
	formatOutput, _ := iconv.ConvertString(output2, "GBK", "UTF-8")
	fmt.Println(formatOutput)

	output, err = Unset("globalEnv")
	formatOutput, _ = iconv.ConvertString(string(output), "GBK", "UTF-8")
	assert.NoError(t, err, formatOutput)

	_, err = EnvInFile("user").Get("globalEnv")
	assert.Error(t, err)
}
