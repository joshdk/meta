package meta

import (
	"fmt"
	u "net/url"
	"testing"
	"time"
)

func TestMustAuthor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         string
		expectedName  string
		expectedEmail string
	}{
		{
			input:         "",
			expectedName:  "",
			expectedEmail: "",
		},
		{
			input:         "John Doe",
			expectedName:  "John Doe",
			expectedEmail: "",
		},
		{
			input:         "John Doe",
			expectedName:  "John Doe",
			expectedEmail: "",
		},
		{
			input:         "jdoe@example.com",
			expectedName:  "",
			expectedEmail: "jdoe@example.com",
		},
		{
			input:         "<jdoe@example.com>",
			expectedName:  "",
			expectedEmail: "jdoe@example.com",
		},
		{
			input:         "Jane Doe <jdoe@example.com>",
			expectedName:  "Jane Doe",
			expectedEmail: "jdoe@example.com",
		},
		{
			input:         "Jane Doe <example@>",
			expectedName:  "Jane Doe <example@>",
			expectedEmail: "",
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			actualName, actualEmail := mustAuthor("", test.input)
			if test.expectedName != actualName {
				t.Fatalf("expected %q but got %q", test.expectedName, actualName)
			}
			if test.expectedEmail != actualEmail {
				t.Fatalf("expected %q but got %q", test.expectedEmail, actualEmail)
			}
		})
	}
}

func TestMustBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "false",
			expected: false,
		},
		{
			input:    "FALSE",
			expected: false,
		},
		{
			input:    "true",
			expected: true,
		},
		{
			input:    "TRUE",
			expected: true,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			actual := mustBool("", test.input)
			if test.expected != actual {
				t.Fatalf("expected %v but got %v", test.expected, actual)
			}
		})
	}
}

func TestMustSHA(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
		panic    bool
	}{
		{
			input:    "",
			expected: "",
		},
		{
			// 7 characters.
			input: "0000000",
			panic: true,
		},
		{
			// 39 characters.
			input: "000000000000000000000000000000000000000",
			panic: true,
		},
		{
			// 40 characters.
			input:    "0000000000000000000000000000000000000000",
			expected: "0000000000000000000000000000000000000000",
		},
		{
			// 41 characters.
			input: "00000000000000000000000000000000000000000",
			panic: true,
		},
		{
			// 40 characters, but one isn't hex.
			input: "000000000000000000_000000000000000000000",
			panic: true,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			defer equalPanic(t, test.panic)
			actual := mustSHA("", test.input)
			equalString(t, test.expected, actual)
		})
	}
}

func TestMustTime(t *testing.T) {
	t.Parallel()

	expected := time.Date(2019, 8, 23, 18, 0, 0, 0, time.UTC)

	tests := []struct {
		input    string
		expected *time.Time
		panic    bool
	}{
		{
			input:    "",
			expected: nil,
		},
		{
			input: "today",
			panic: true,
		},
		{
			// $ date -R
			input:    "Fri, 23 Aug 2019 11:00:00 -0700",
			expected: &expected,
		},
		{
			// $ date -u +%Y-%m-%dT%H:%M:%SZ
			input:    "2019-08-23T18:00:00Z",
			expected: &expected,
		},
		{
			// $ date -Iseconds
			// $ date --iso-8601=seconds
			input:    "2019-08-23T11:00:00-07:00",
			expected: &expected,
		},
		{
			// $ date --iso-8601=seconds (no colon)
			input:    "2019-08-23T11:00:00-0700",
			expected: &expected,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			defer equalPanic(t, test.panic)
			actual := mustTime("", test.input)
			equalTime(t, test.expected, actual)
		})
	}
}

func TestMustURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected *u.URL
		panic    bool
	}{
		{
			input:    "",
			expected: nil,
		},
		{
			input: "http://",
			panic: true,
		},
		{
			input: "http://",
			panic: true,
		},
		{
			input: "example.com",
			panic: true,
		},
		{
			input: "http://localhost:http",
			panic: true,
		},
		{
			input:    "http://localhost",
			expected: &u.URL{Scheme: "http", Host: "localhost"},
		},
		{
			input:    "http://127.0.0.1",
			expected: &u.URL{Scheme: "http", Host: "127.0.0.1"},
		},
		{
			input:    "http://example.com",
			expected: &u.URL{Scheme: "http", Host: "example.com"},
		},
		{
			input:    "https://example.com",
			expected: &u.URL{Scheme: "https", Host: "example.com"},
		},
		{
			input:    "http://example.com:8080/demo",
			expected: &u.URL{Scheme: "http", Host: "example.com:8080", Path: "/demo"},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			defer equalPanic(t, test.panic)
			actual := mustURL("", test.input)
			equalURL(t, test.expected, actual)
		})
	}
}

func equalPanic(t *testing.T, panics bool) {
	t.Helper()

	err := recover()

	switch {
	case err == nil && panics:
		t.Fatal("expected a panic")
	case err != nil && !panics:
		t.Fatal("did not expect a panic")
	}
}

func equalString(t *testing.T, expected, actual string) {
	t.Helper()

	if actual != expected {
		t.Fatalf("expected %q but got %q", expected, actual)
	}
}

func equalTime(t *testing.T, expected, actual *time.Time) {
	t.Helper()

	switch {
	case expected == nil && actual == nil:
		return
	case expected != nil && actual == nil:
		t.Fatalf("expected %v but got nil", expected)
	case expected == nil && actual != nil:
		t.Fatalf("expected nil but got %v", actual)
	case *expected != *actual:
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}

func equalURL(t *testing.T, expected, actual *u.URL) {
	t.Helper()

	switch {
	case expected == nil && actual == nil:
		return
	case expected != nil && actual == nil:
		t.Fatalf("expected %v but got nil", expected)
	case expected == nil && actual != nil:
		t.Fatalf("expected nil but got %v", actual)
	case *expected != *actual:
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}
