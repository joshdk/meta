// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

package meta

import (
	"fmt"
	"net/mail"
	u "net/url"
	"strconv"
	"time"
)

// mustAuthor validates that the given value contains the author's name and
// potentially email.
func mustAuthor(_, raw string) (string, string) {
	if raw == "" {
		return "", ""
	}

	parsed, err := mail.ParseAddress(raw)
	if err != nil {
		return raw, ""
	}

	return parsed.Name, parsed.Address
}

// mustBool validates that the given value is a properly formatted boolean.
func mustBool(_, raw string) bool {
	if raw == "" {
		return false
	}

	if b, err := strconv.ParseBool(raw); err == nil {
		return b
	}

	return false
}

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

// mustURL validates that the given value is a properly formatted URL.
func mustURL(path, raw string) *u.URL {
	if raw == "" {
		return nil
	}

	// Parse the URL.
	parsed, err := u.Parse(raw)
	if err != nil {
		panic(fmt.Errorf("malformed ldflags value for %s", path))
	}

	// Require that the scheme is http:// or https://.
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		panic(fmt.Errorf("malformed ldflags value for %s", path))
	}

	// Require that the URL contained a host.
	if parsed.Host == "" {
		panic(fmt.Errorf("malformed ldflags value for %s", path))
	}

	return parsed
}
