// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

// package meta provides common application metadata for use with go build and
// ldflags. This package is intended to be imported, where variable values can
// be set by using -X arguments, tp the -ldflags argument, when running go
// build. See https://pkg.go.dev/cmd/go and https://pkg.go.dev/cmd/link.
//
// List of variable names:
//   jdk.sh/meta.date
//   jdk.sh/meta.sha
//   jdk.sh/meta.version
package meta

import "time"

// date is the time that the application was built. Supports several common
// formats.
//
// Variable name:
//   jdk.sh/meta.date
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.date=Wed Sep 15 06:55:44 PDT 2021'"
//   -ldflags "-X 'jdk.sh/meta.date=2021-09-15T13:56:43Z'"
//   -ldflags "-X 'jdk.sh/meta.date=$(date)'"
//   -ldflags "-X 'jdk.sh/meta.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)'"
var date string

var dateParsed = mustTime("jdk.sh/meta.date", date)

// Date is the time at which the application was built.
func Date() *time.Time {
	return dateParsed
}

// sha is the git SHA that was used to build the application. A 40 character
// "long" SHA should be provided.
//
// Variable name:
//   jdk.sh/meta.sha
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.sha=bb2fecbb4a287ea4c1f9887ca86dd0eb7ff28ec6'"
//   -ldflags "-X 'jdk.sh/meta.sha=$(git rev-parse HEAD)'"
var sha string

var shaParsed = mustSHA("jdk.sh/meta.sha", sha)

// SHA is the git SHA used to build the application.
func SHA() string {
	return shaParsed
}

// ShortSHA is the git "short" SHA used to build the application.
func ShortSHA() string {
	if shaParsed == "" {
		return ""
	}

	return shaParsed[:7]
}

// version is the version slug. The value can be used to point back to
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

// Version is the version slug for the application.
func Version() string {
	return version
}
