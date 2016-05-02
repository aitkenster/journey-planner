package main

import "testing"

func TestValidateCoordinates(t *testing.T) {
	tests := []struct {
		coords   Coordinates
		expected bool
	}{
		{
			coords:   Coordinates{51.5034070, -0.1275920},
			expected: true,
		},
		{
			// invalid latitude
			coords:   Coordinates{91.5034070, -0.1275920},
			expected: false,
		},
		{
			// invalid longitude
			coords:   Coordinates{-1.5034070, -370.1275920},
			expected: false,
		},
	}

	for _, test := range tests {
		got := test.coords.AreValid()
		if got != test.expected {
			t.Errorf("expected %v for coords %v, got %v", test.expected, test.coords, got)
		}
	}
}
