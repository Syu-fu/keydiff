package input

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVReader struct{}

func (r *CSVReader) Read(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file %s: %w", filePath, err)
	}

	return records, nil
}
