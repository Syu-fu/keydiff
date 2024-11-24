package cli

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"keydiff/interfaces/input"
	"keydiff/interfaces/output"
	"keydiff/usecases"

	"github.com/urfave/cli/v2"
)

const RequiredArguments = 2

func RunCLI(args []string) error {
	app := &cli.App{
		Name:  "keydiff",
		Usage: "Find differences between two CSV files using a specified key column",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key",
				Usage: "Comma-separated list of column indices to use as the key (default: 0)",
				Value: "0",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != RequiredArguments {
				return errors.New("exactly two arguments are required: <original.csv> <modified.csv>")
			}

			originalFilePath := c.Args().Get(0)
			modifiedFilePath := c.Args().Get(1)

			keyIndicesStr := c.String("key")
			keyIndices := parseKeyIndices(keyIndicesStr)

			reader := input.CSVReader{}
			original, err := reader.Read(originalFilePath)
			if err != nil {
				return fmt.Errorf("error reading original file: %w", err)
			}

			modified, err := reader.Read(modifiedFilePath)
			if err != nil {
				return fmt.Errorf("error reading modified file: %w", err)
			}

			calculator := usecases.DiffCalculator{}
			results := calculator.CalculateDiff(original, modified, keyIndices)

			formatter := output.TextFormatter{}
			output := formatter.Format(results)

			fmt.Println(output)
			return nil
		},
	}

	err := app.Run(args)
	if err != nil {
		log.Println("Failed to execute CLI:", err)
		return fmt.Errorf("failed to execute CLI: %w", err)
	}

	return nil
}

func parseKeyIndices(keyIndicesStr string) []int {
	parts := strings.Split(keyIndicesStr, ",")

	var indices []int

	for _, part := range parts {
		var index int

		_, err := fmt.Sscanf(part, "%d", &index)
		if err != nil {
			log.Fatalf("Invalid key index: %v", err)
		}

		indices = append(indices, index)
	}

	return indices
}
