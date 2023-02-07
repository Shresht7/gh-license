package helpers

import (
	"testing"
)

func TestDetermineOwnerAndRepo(t *testing.T) {
	// Test cases
	testCases := []struct {
		arg      string
		expected string
	}{
		{"", "Shresht7/gh-license"},
		{"Shresht7/gh-license", "Shresht7/gh-license"},
		{"gh-license", "Shresht7/gh-license"},
	}

	// Run tests
	for _, tc := range testCases {
		owner, repo, err := DetermineOwnerAndRepo(tc.arg)
		if err != nil {
			t.Error(err)
		}
		if owner+"/"+repo != tc.expected {
			t.Errorf("Expected: %s, Got: %s", tc.expected, owner+"/"+repo)
		}
	}
}
