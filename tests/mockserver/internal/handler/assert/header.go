// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package assert

import (
	"fmt"
	"net/http"
)

// HeaderExists generically verifies the given request header name exists with
// a single, non-empty value. If an unexpected value is found, it will return an
// error with a diagnostic text body and which should cause the handler to return
// immediately.
func HeaderExists(req *http.Request, name string) error {
	values := req.Header.Values(name)

	if len(values) == 0 {
		return fmt.Errorf("missing expected header %s", name)
	}

	if len(values) > 1 {
		return fmt.Errorf("expected single header for %s, got %d values", name, len(values))
	}

	got := values[0]

	if got == "" {
		return fmt.Errorf("expected non-empty header %s", name)
	}

	return nil
}

// HeaderValues generically verifies the given request header name has the
// expected values. If an unexpected value is found, it will return an error
// with a diagnostic text body and which should cause the handler to return
// immediately.
func HeaderValues(req *http.Request, name string, expected []string) error {
	got := req.Header.Values(name)

	if len(got) == 0 {
		return fmt.Errorf(
			"missing expected header %s which should be %v",
			name,
			expected,
		)
	}

	if len(got) != len(expected) {
		return fmt.Errorf(
			"expected header %s to be %v, got: %v",
			name,
			expected,
			got,
		)
	}

	for index, expectedValue := range expected {
		if got[index] != expectedValue {
			return fmt.Errorf(
				"expected header %s to be %v, got: %v",
				name,
				expected,
				got,
			)
		}
	}

	return nil
}
