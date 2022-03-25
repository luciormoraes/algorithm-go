package searchinsertposition

import "testing"

type checkPositions struct {
	name             string
	arrayToCheck     []int
	target, expected int
}

func TestSearch(t *testing.T) {
	tests := []checkPositions{
		{
			name:         "Positions must be 4",
			arrayToCheck: []int{-1, 0, 3, 5, 9, 12},
			target:       9,
			expected:     4,
		},
		{
			name:         "Position must be 2",
			arrayToCheck: []int{-1, 0, 3, 5, 9, 12},
			target:       0,
			expected:     1,
		},
		{
			name:         "Preparation time for default case",
			arrayToCheck: []int{-1, 0, 3, 5, 9, 12},
			target:       13,
			expected:     6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.arrayToCheck, tt.target); got != tt.expected {
				t.Errorf("PreparationTime(%v, %d) = %d; want %d", tt.arrayToCheck, tt.target, got, tt.expected)
			}
		})

	}
}
