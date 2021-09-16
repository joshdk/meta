// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

// package meta provides common application metadata for use with go build and
// ldflags. This package is intended to be imported, where variable values can
// be set by using -X arguments, tp the -ldflags argument, when running go
// build. See https://pkg.go.dev/cmd/go and https://pkg.go.dev/cmd/link.
//
// List of variable names:
//   jdk.sh/meta.version
package meta

// version is an arbitrary version slug. The value can be used to point back to
// a specific tag or release.
//
// Variable name:
//   jdk.sh/meta.version
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.version=development'"
//   -ldflags "-X 'jdk.sh/meta.version=v1.0.0'"
//   -ldflags "-X 'jdk.sh/meta.version=$(git describe)'"
var version string

// Version is an arbitrary version slug. The value can be used to point back to
// a specific tag or release.
func Version() string {
	return version
}
