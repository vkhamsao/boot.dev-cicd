package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		authHeader  http.Header
		expectError bool
	}{
		{name: "empty header", authHeader: http.Header{}, expectError: true},
		{
			name:        "malformed header",
			authHeader:  http.Header{"Authorization": []string{"Bearer"}},
			expectError: true,
		},
		{
			name:        "correct header",
			authHeader:  http.Header{"Authorization": []string{"ApiKey 123"}},
			expectError: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keyString, err := GetAPIKey(tc.authHeader)
			if tc.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error but got %v", err)
				if keyString != "123" {
					t.Errorf("expected key to be 123 but got %v", keyString)
				}
			}
		})
	}
}
