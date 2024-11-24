package usecases

import (
	"strings"

	"keydiff/entities"
)

type DiffCalculator struct{}

func (dc *DiffCalculator) CalculateDiff(original, modified [][]string, keyIndices []int) []entities.DiffResult {
	results := []entities.DiffResult{}

	originalMap := make(map[string][]string)
	modifiedMap := make(map[string][]string)

	generateKey := func(row []string, indices []int) string {
		var keyParts []string

		for _, index := range indices {
			if index < len(row) {
				keyParts = append(keyParts, row[index])
			}
		}

		return strings.Join(keyParts, ":")
	}

	for _, row := range original {
		key := generateKey(row, keyIndices)
		if key != "" {
			originalMap[key] = row
		}
	}

	for _, row := range modified {
		key := generateKey(row, keyIndices)
		if key != "" {
			modifiedMap[key] = row
		}
	}

	for key, originalRow := range originalMap {
		if modifiedRow, exists := modifiedMap[key]; exists {
			if !equalRows(originalRow, modifiedRow) {
				results = append(results, entities.DiffResult{
					Key:      key,
					Original: originalRow,
					Modified: modifiedRow,
				})
			}
		} else {
			results = append(results, entities.DiffResult{
				Key:      key,
				Original: originalRow,
				Modified: nil,
			})
		}
	}

	for key, modifiedRow := range modifiedMap {
		if _, exists := originalMap[key]; !exists {
			results = append(results, entities.DiffResult{
				Key:      key,
				Original: nil,
				Modified: modifiedRow,
			})
		}
	}

	return results
}

func equalRows(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
