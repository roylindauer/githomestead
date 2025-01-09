package gitserver

import "testing"

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		s        string
		expected string
		suffix   string
	}{
		{
			s:        "myrepo.git",
			expected: "myrepo",
			suffix:   ".git",
		},
		{
			s:        "myrepo.git.git",
			expected: "myrepo.git",
			suffix:   ".git",
		},
		{
			s:        "myrepo.git",
			expected: "myrepo.git",
			suffix:   ".gi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := TrimSuffix(tt.s, tt.suffix); got != tt.expected {
				t.Errorf("TrimSuffix() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		s              string
		expectedResult string
		expectedError  string
	}{
		{
			s:              "This is a string",
			expectedResult: "this-is-a-string",
			expectedError:  "",
		},
		{
			s:              "This is a !@#$%^&*()_=+string",
			expectedResult: "this-is-a-string",
			expectedError:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := Slugify(tt.s); got != tt.expectedResult {
				t.Errorf("Slugify() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
