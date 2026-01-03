package tests

import (
	"regexp"
	"testing"
)

func ValidateEmail(email string) bool {
	var rgxp = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return rgxp.MatchString(email)
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Email Valido",
			input:    "gabriel@test.com",
			expected: true,
		},
		{
			name:     "Faltando arroba",
			input:    "gabrieltest.com",
			expected: false,
		},
		{
			name:     "Email invalido",
			input:    "gabriel",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmail(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateEmail(%s) = %v esperado: %v", tt.input, result, tt.expected)
			}
		})
	}
}
