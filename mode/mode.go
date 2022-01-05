// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// 运行模式
package mode

import (
	"io"
	"os"
)

// EnvGinMode indicates environment name for kn mode.
const EnvKNMode = "KN_MODE"

const (
	// DebugMode indicates kn mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates kn mode is release.
	ReleaseMode = "release"
	// TestMode indicates kn mode is test.
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used by Gin for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
// 		import "github.com/mattn/go-colorable"
// 		kn.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by Gin to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var knMode = debugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(EnvKNMode)
	SetMode(mode)
}

// SetMode sets  mode according to input string.
func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	switch value {
	case DebugMode:
		knMode = debugCode
	case ReleaseMode:
		knMode = releaseCode
	case TestMode:
		knMode = testCode
	default:
		panic("kn mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}

// Mode returns currently  mode.
func Mode() string {
	return modeName
}
func IsDebug() bool {
	return modeName == DebugMode
}
func IsTest() bool {
	return modeName == TestMode
}
func IsRelease() bool {
	return modeName == ReleaseMode
}
