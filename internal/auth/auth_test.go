package auth

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GetAPIKey(t *testing.T) {

	tests := map[string]struct {
		headers  map[string][]string
		expected string
		isErr    bool
	}{
		"no auth header": {
			headers:  map[string][]string{},
			expected: "",
			isErr:    true,
		},
		"malformed auth header": {
			headers:  map[string][]string{"Authorization": {"malformed"}},
			expected: "",
			isErr:    true,
		},
		"valid auth header": {
			headers:  map[string][]string{"Authorization": {"ApiKey abc123"}},
			expected: "abc123",
			isErr:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			apiKey, err := GetAPIKey(test.headers)
			isError := err != nil
			if isError != test.isErr {
				t.Fatalf("unexpected error value, is errror expected: [%v], got [%#v]", test.isErr, err)
			}

			diff := cmp.Diff(test.expected, apiKey)
			if diff != "" {
				t.Fatalf("unexpected api key (-want +got):\n%s", diff)
			}
		})
	}
}
