// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.

package meta

import (
	"fmt"
	"testing"
	"time"
)

func TestMeta(t *testing.T) {
	t.Parallel()

	expectedDate := time.Date(2019, 8, 23, 18, 0, 0, 0, time.UTC)

	tests := []struct {
		flags    map[string]string
		assertfn func(*testing.T, *info)
		panics   bool
	}{
		{
			// Validate that the test program does not panic when no
			// definitions are given.
		},
		{
			// Value for jdk.sh/meta.date that is valid.
			flags: map[string]string{
				"jdk.sh/meta.date": "Fri, 23 Aug 2019 11:00:00 -0700",
			},
			assertfn: func(t *testing.T, actual *info) {
				equalTime(t, &expectedDate, actual.Date)
			},
		},
		{
			// Value for jdk.sh/meta.date that causes a panic.
			flags: map[string]string{
				"jdk.sh/meta.date": "tomorrow",
			},
			panics: true,
		},
	}

	for index, test := range tests {
		test := test

		t.Run(fmt.Sprint(index), func(t *testing.T) {
			t.Parallel()

			actual, failed := execTestJSON(t, test.flags)
			switch {
			case failed && !test.panics:
				t.Error("expected test success but got failure")
			case !failed && test.panics:
				t.Error("expected test failure but got success")
			case !failed && !test.panics && test.assertfn != nil:
				test.assertfn(t, actual)
			}
		})
	}
}
