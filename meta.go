// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

// package meta provides common application metadata for use with go build and
// ldflags. This package is intended to be imported, where variable values can
// be set by using -X arguments, tp the -ldflags argument, when running go
// build. See https://pkg.go.dev/cmd/go and https://pkg.go.dev/cmd/link.
//
// List of variable names:
//   jdk.sh/meta.author
//   jdk.sh/meta.author_url
//   jdk.sh/meta.copyright
//   jdk.sh/meta.date
//   jdk.sh/meta.desc
//   jdk.sh/meta.dev
//   jdk.sh/meta.docs
//   jdk.sh/meta.license
//   jdk.sh/meta.license_url
//   jdk.sh/meta.name
//   jdk.sh/meta.sha
//   jdk.sh/meta.title
//   jdk.sh/meta.url
//   jdk.sh/meta.version
package meta

import (
	u "net/url"
	"runtime"
	"time"
)

// author is the name of the application author. May contain their name, email
// address, or optionally both.
//
// Variable name:
//   jdk.sh/meta.author
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.author=John Doe'"
//   -ldflags "-X 'jdk.sh/meta.author=jdoe@example.com'"
//   -ldflags "-X 'jdk.sh/meta.author=Jane Doe <jdoe@example.com>'"
var author string

var authorParsed, authorEmailParsed = mustAuthor("jdk.sh/meta.author", author)

// Author is the name of the application author.
func Author() string {
	return authorParsed
}

// AuthorEmail is the email address for the application author.
func AuthorEmail() string {
	return authorEmailParsed
}

// author_url is a URL for the application author. Typically links to the
// author's personal homepage or Github profile.
//
// Variable name:
//   jdk.sh/meta.author_url
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.author_url=https://example.com/profile'"
var author_url string

var authorURLParsed = mustURL("jdk.sh/meta.author_url", author_url)

// AuthorURL is the homepage URL for the application author.
func AuthorURL() *u.URL {
	return authorURLParsed
}

// Arch is the architecture target that the application is running on.
func Arch() string {
	return runtime.GOARCH
}

// copyright is the copyright for the application. Typically the name if the
// author or organization, sometimes prefixed with a year or year range.
//
// Variable name:
//   jdk.sh/meta.copyright
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.copyright=John Doe'"
//   -ldflags "-X 'jdk.sh/meta.copyright=2021 Jane Doe'"
//   -ldflags "-X 'jdk.sh/meta.copyright=2019-2021 Jim Doe'"
var copyright string

// Copyright is the copyright for the application.
func Copyright() string {
	return copyright
}

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

// desc is a description for the application. Typically a longer statement
// describing what the application does.
//
// Variable name:
//   jdk.sh/meta.desc
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.desc=A super simple demonstration application'"
var desc string

// Description is the description of the application.
func Description() string {
	return desc
}

// dev is the development status for the application. An application in
// development mode may indicate that it's using experimental or untested
// features, and should be used with caution.
//
// Variable name:
//   jdk.sh/meta.dev
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.dev=true'"
var dev string

var devParsed = mustBool("jdk.sh/meta.dev", dev)

// Development is the development status for the application.
func Development() bool {
	return devParsed
}

// docs is a URL for application documentation. Typically links to a page where
// a user can find technical documentation.
//
// Variable name:
//   jdk.sh/meta.docs
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.docs=https://example.com/demo/README.md'"
var docs string

var docsParsed = mustURL("jdk.sh/meta.docs", docs)

// Docs is the documentation URL for the application.
func Docs() *u.URL {
	return docsParsed
}

// Go is the version of the Go runtime that the application is running on.
func Go() string {
	return runtime.Version()
}

// license is the license identifier for the application. Should not the full
// license body, but one of the identifiers from https://spdx.org/licenses, so
// that the type of license can be easily determined.
//
// Variable name:
//   jdk.sh/meta.license
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.license=Apache-2.0'"
//   -ldflags "-X 'jdk.sh/meta.license=MIT'"
//   -ldflags "-X 'jdk.sh/meta.license=WTFPL'"
var license string

// License is the license identifier for the application.
func License() string {
	return license
}

// license_url is a URL for application license. Typically links to a page
// where the verbatim license body is available.
//
// Variable name:
//   jdk.sh/meta.license_url
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.license_url=https://example.com/demo/LICENSE.txt'"
var license_url string

var licenseURLParsed = mustURL("jdk.sh/meta.license_url", license_url)

// LicenseURL is the license URL for the application.
func LicenseURL() *u.URL {
	return licenseURLParsed
}

// name is the name of the application. Typically named the same as the binary,
// or for display in an error or help message.
//
// Variable name:
//   jdk.sh/meta.name
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.name=demo-app'"
var name string

// Name is the name of the application.
func Name() string {
	return name
}

// OS is the operating system target that the application is running on.
func OS() string {
	return runtime.GOOS
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

// title is the title of the application. Typically a full or non-abbreviated
// form of the application name.
//
// Variable name:
//   jdk.sh/meta.title
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.title=Demo Application'"
var title string

// Title is the title of the application.
func Title() string {
	return title
}

// url is a URL for the application homepage. Typically links to a page where a
// user can learn more about the application.
//
// Variable name:
//   jdk.sh/meta.url
//
// Examples:
//   -ldflags "-X 'jdk.sh/meta.url=https://example.com/demo'"
var url string

var urlParsed = mustURL("jdk.sh/meta.url", url)

// URL is the homepage URL for the application.
func URL() *u.URL {
	return urlParsed
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
