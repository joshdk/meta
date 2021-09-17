// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

package meta

import (
	"fmt"
	"time"
)

// mustSHA validates that the given value is a properly formatted git SHA.
func mustSHA(path, raw string) string {
	if raw == "" {
		return ""
	}

	// Git SHAs are 40 characters long.
	const gitSHALength = 40
	if len(raw) != gitSHALength {
		panic(fmt.Errorf("malformed ldflags value for %s", path))
	}

	// Git SHAs are made of only lowercase hex characters.
	for _, rune := range raw {
		switch {
		case '0' <= rune && rune <= '9':
		case 'a' <= rune && rune <= 'f':
		default:
			panic(fmt.Errorf("malformed ldflags value for %s", path))
		}
	}

	return raw
}

// mustTime validates that the given value is a properly formatted timestamp.
func mustTime(path, raw string) *time.Time {
	if raw == "" {
		return nil
	}

	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		time.UnixDate,
	}

	// Try each layout until one parses.
	for _, spec := range layouts {
		if t, err := time.Parse(spec, raw); err == nil {
			return &t
		}
	}

	panic(fmt.Errorf("malformed ldflags value for %s", path))
}
