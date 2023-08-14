package helper

import (
	"testing"
)

func TestIsFilePathValid(t *testing.T) {
	tests := []struct {
		path   string
		result bool
		name   string
	}{
		{
			path:   "./testData",
			result: true,
			name:   "IsFilePathValidPositive",
		},
		{
			path:   ".//incorect$%^@#@3\\",
			result: false,
			name:   "IsFilePathValidNegative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := IsFilePathValid(tt.path)

			if tt.result != isValid {
				t.Fatalf("Test %s is failed received value %v is not expected %v", tt.name, isValid, tt.result)
			}
		})
	}
}
