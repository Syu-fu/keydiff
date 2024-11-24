package output

import (
	"fmt"
	"strings"

	"keydiff/entities"
)

type TextFormatter struct{}

func (f *TextFormatter) Format(results []entities.DiffResult) string {
	var builder strings.Builder

	for _, result := range results {
		builder.WriteString(fmt.Sprintf("Key: %s\n", result.Key))

		if result.Original != nil {
			builder.WriteString(fmt.Sprintf("-: %s\n", strings.Join(result.Original, ", ")))
		}

		if result.Modified != nil {
			builder.WriteString(fmt.Sprintf("+: %s\n", strings.Join(result.Modified, ", ")))
		}

		builder.WriteString("\n")
	}

	return builder.String()
}
