package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/grokify/mogo/text/markdown"

	"github.com/grokify/slogo/examples"
	"github.com/grokify/slogo/ontology"
)

func main() {
	// Collect SLOs from all directories
	directories := examples.ExampleSLOsByDirectory()

	// Get label definitions
	labelDefs := ontology.GetLabelDefinitions()

	// Open output file
	f, err := os.Create("examples/METRICS.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Write header
	fmt.Fprintln(f, "# SLO Metrics by Ontology Labels")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "This document shows the distribution of SLOs across different ontology label dimensions.")
	fmt.Fprintln(f)

	// Get sorted label names
	labelOrder := ontology.Labels()

	// For each label type, create a table
	for _, labelName := range labelOrder {
		labelValues, exists := labelDefs[labelName]
		if !exists {
			continue
		}

		fmt.Fprintf(f, "\n## %s\n\n", formatLabelName(labelName))

		// Build table in memory
		var tableBuilder strings.Builder

		// Create table header
		tableBuilder.WriteString("| Directory |")
		for _, value := range labelValues {
			fmt.Fprintf(&tableBuilder, " %s |", formatValue(value))
		}
		tableBuilder.WriteString(" Total |\n")

		// Create separator
		tableBuilder.WriteString("|-----------|")
		for range labelValues {
			tableBuilder.WriteString("------|")
		}
		tableBuilder.WriteString("-------|\n")

		// Get sorted directory names
		dirNames := make([]string, 0, len(directories))
		for dir := range directories {
			dirNames = append(dirNames, dir)
		}
		sort.Strings(dirNames)

		// For each directory, count SLOs by label value
		for _, dir := range dirNames {
			slos := directories[dir]
			counts := make(map[string]int)
			total := 0

			// Count SLOs for each label value
			for _, slo := range slos {
				if labels, exists := slo.Metadata.Labels[labelName]; exists {
					for _, labelValue := range labels {
						counts[labelValue]++
						total++
					}
				}
			}

			// Write row
			fmt.Fprintf(&tableBuilder, "| %s |", dir)
			for _, value := range labelValues {
				count := counts[value]
				if count > 0 {
					fmt.Fprintf(&tableBuilder, " %d |", count)
				} else {
					tableBuilder.WriteString(" - |")
				}
			}
			fmt.Fprintf(&tableBuilder, " %d |\n", total)
		}

		// Align table and write to file
		alignedTable := markdown.TableAlign(tableBuilder.String(), 1)
		fmt.Fprintln(f, alignedTable)
	}

	fmt.Println("Generated examples/METRICS.md")
}

func formatLabelName(name string) string {
	// Convert label-name to Label Name
	parts := strings.Split(name, "-")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, " ")
}

func formatValue(value string) string {
	// Convert value to proper case
	parts := strings.Split(value, "-")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, " ")
}
