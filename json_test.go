// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

package meta

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	u "net/url"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// info is a placeholder struct that only exists to store values from each
// public function in this package.
type info struct {
	Arch              string
	Author            string
	AuthorEmail       string
	AuthorURL         *u.URL
	Copyright         string
	Date              *time.Time
	Description       string
	Development       bool
	Docs              *u.URL
	Go                string
	License           string
	LicenseURL        *u.URL
	Name              string
	Note              string
	OS                string
	SHA               string
	ShortSHA          string
	Source            *u.URL
	Title             string
	URL               *u.URL
	Version           string
	VersionBuild      string
	VersionMajor      string
	VersionMinor      string
	VersionPatch      string
	VersionPreRelease string
}

// TestJSON serializes the placeholder info struct as JSON to stdout.
func TestJSON(t *testing.T) {
	t.Parallel()

	// Store a value from each public function in this package.
	info := info{
		Arch:              Arch(),
		Author:            Author(),
		AuthorEmail:       AuthorEmail(),
		AuthorURL:         AuthorURL(),
		Copyright:         Copyright(),
		Date:              Date(),
		Description:       Description(),
		Development:       Development(),
		Docs:              Docs(),
		Go:                Go(),
		License:           License(),
		LicenseURL:        LicenseURL(),
		Name:              Name(),
		Note:              Note(),
		OS:                OS(),
		SHA:               SHA(),
		ShortSHA:          ShortSHA(),
		Source:            Source(),
		Title:             Title(),
		URL:               URL(),
		Version:           Version(),
		VersionBuild:      VersionBuild(),
		VersionMajor:      VersionMajor(),
		VersionMinor:      VersionMinor(),
		VersionPatch:      VersionPatch(),
		VersionPreRelease: VersionPreRelease(),
	}

	if err := json.NewEncoder(os.Stdout).Encode(info); err != nil {
		t.Fatal(err)
	}
}

// execTestJSON executes the specially crafted test TestJSON, by constructing a
// go test command line along with a custom set of ldflags. This causes the
// executed TestJSON test to react in a manner identical to a normal main()
// function. Including potential panic if an invalid ldflags value is passed.
// The TestJSON function marshals a single-line JSON string that contains a
// value for every public function in the package. This JSON string is searched
// for in the exec output, unmarshaled, and used in test assertions.
func execTestJSON(t *testing.T, args map[string]string) (*info, bool) {
	t.Helper()

	// Dynamically construct a -ldflags command line argument list.
	var ldflags string
	for key, value := range args {
		ldflags += fmt.Sprintf(`-X '%s=%s' `, key, value)
	}

	// Construct a go test command line.
	output, err := exec.Command("go", "test", "-ldflags", ldflags, "-test.run", "^TestJSON$", "-test.v", ".").Output() // nolint:lll
	if err != nil {
		// Error was not an exit error with the test program. Indicates an
		// error with the test running setup itself.
		if _, ok := err.(*exec.ExitError); !ok { // nolint:errorlint
			t.Fatal(err)
		}

		// Program has a runtime error, but it was not a panic.
		if !strings.HasPrefix(string(output), "panic:") {
			t.Fatal(err)
		}

		// We failed...successfully!
		return nil, true
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(output))
	for scanner.Scan() {
		// Attempt to parse each line until we find one that is actually JSON.
		var info info
		if err := json.Unmarshal(scanner.Bytes(), &info); err == nil {
			return &info, false
		}
	}

	// None of the lines contained any JSON.
	t.Fatal("no JSON could be parsed")

	return nil, true
}
