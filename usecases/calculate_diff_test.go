package usecases

import (
	"reflect"
	"testing"

	"keydiff/entities"
)

// nolint:funlen
func TestCalculateDiff(t *testing.T) {
	dc := DiffCalculator{}

	tests := []struct {
		name           string
		original       [][]string
		modified       [][]string
		keyIndices     []int
		expectedResult []entities.DiffResult
	}{
		{
			name: "No differences with single key",
			original: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			keyIndices:     []int{0},
			expectedResult: []entities.DiffResult{},
		},
		{
			name: "Row removed in modified with single key",
			original: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
			},
			keyIndices: []int{0},
			expectedResult: []entities.DiffResult{
				{
					Key:      "banana",
					Original: []string{"banana", "1111", "yen"},
					Modified: nil,
				},
			},
		},
		{
			name: "Row content differs with single key",
			original: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1141", "yen"},
			},
			keyIndices: []int{0},
			expectedResult: []entities.DiffResult{
				{
					Key:      "banana",
					Original: []string{"banana", "1111", "yen"},
					Modified: []string{"banana", "1141", "yen"},
				},
			},
		},
		{
			name: "No differences with composite key",
			original: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			keyIndices:     []int{0, 2},
			expectedResult: []entities.DiffResult{},
		},
		{
			name: "Row added in modified with composite key",
			original: [][]string{
				{"apple", "1521", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			keyIndices: []int{0, 2},
			expectedResult: []entities.DiffResult{
				{
					Key:      "banana:yen",
					Original: nil,
					Modified: []string{"banana", "1111", "yen"},
				},
			},
		},
		{
			name: "Row content differs with composite key",
			original: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1111", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
				{"banana", "1141", "yen"},
			},
			keyIndices: []int{0, 2},
			expectedResult: []entities.DiffResult{
				{
					Key:      "banana:yen",
					Original: []string{"banana", "1111", "yen"},
					Modified: []string{"banana", "1141", "yen"},
				},
			},
		},
		{
			name:           "Both files empty with composite key",
			original:       [][]string{},
			modified:       [][]string{},
			keyIndices:     []int{0, 2},
			expectedResult: []entities.DiffResult{},
		},
		{
			name: "Invalid key index in composite key",
			original: [][]string{
				{"apple", "1521", "yen"},
			},
			modified: [][]string{
				{"apple", "1521", "yen"},
			},
			keyIndices:     []int{3},
			expectedResult: []entities.DiffResult{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := dc.CalculateDiff(tt.original, tt.modified, tt.keyIndices)

			if !reflect.DeepEqual(results, tt.expectedResult) {
				t.Errorf("got %v, want %v", results, tt.expectedResult)
			}
		})
	}
}
