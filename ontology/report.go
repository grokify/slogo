package ontology

import (
	"fmt"
	"io"
	"sort"
	"strings"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/gocharts/v2/data/histogram"
	"github.com/grokify/mogo/text/markdown"
)

// LabelDistribution represents the distribution of values for a specific label
type LabelDistribution struct {
	LabelName    string               `json:"label_name"`
	Histogram    *histogram.Histogram `json:"histogram"`
	TotalMetrics int                  `json:"total_metrics"`
}

// MetricReport represents the complete label analysis for a directory
type MetricReport struct {
	Directory          string                        `json:"directory"`
	TotalSLOs          int                           `json:"total_slos"`
	LabelDistributions map[string]*LabelDistribution `json:"label_distributions"`
}

// AnalyzeLabels analyzes a slice of SLOs and generates a label distribution report
func AnalyzeLabels(directory string, slos []v1.SLO) *MetricReport {
	report := &MetricReport{
		Directory:          directory,
		TotalSLOs:          len(slos),
		LabelDistributions: make(map[string]*LabelDistribution),
	}

	// Initialize histograms for all label types with all possible values
	labelDefs := GetLabelDefinitions()

	// Initialize all label distributions with zero counts
	for labelName, possibleValues := range labelDefs {
		hist := histogram.Histogram{
			Name:  labelName,
			Items: make(map[string]int),
		}
		// Initialize all possible values with 0
		for _, value := range possibleValues {
			hist.Items[value] = 0
		}
		report.LabelDistributions[labelName] = &LabelDistribution{
			LabelName:    labelName,
			Histogram:    &hist,
			TotalMetrics: 0,
		}
	}

	// Count actual label values from SLOs
	for _, slo := range slos {
		for labelKey, labelValues := range slo.Metadata.Labels {
			if dist, exists := report.LabelDistributions[labelKey]; exists {
				for _, labelValue := range labelValues {
					dist.Histogram.Add(labelValue, 1)
					dist.TotalMetrics++
				}
			}
		}
	}

	return report
}

// PrintReport prints a Markdown-formatted table of a single report
func PrintReport(w io.Writer, report *MetricReport) {
	fmt.Fprintln(w)
	fmt.Fprintf(w, "## %s\n\n", report.Directory)
	fmt.Fprintf(w, "**Total SLOs:** %d\n\n", report.TotalSLOs)

	// Sort label names
	labelNames := make([]string, 0, len(report.LabelDistributions))
	for labelName := range report.LabelDistributions {
		labelNames = append(labelNames, labelName)
	}
	sort.Strings(labelNames)

	for _, labelName := range labelNames {
		dist := report.LabelDistributions[labelName]
		fmt.Fprintf(w, "### %s\n\n", dist.LabelName)

		// Build table as a string
		var tableBuilder strings.Builder
		tableBuilder.WriteString("| Value | Count |\n")
		tableBuilder.WriteString("|-------|-------|\n")

		// Sort bin names
		binNames := dist.Histogram.BinNames()

		// Create a map of display values to original bin names and counts
		type valueCount struct {
			displayValue string
			binName      string
			count        int
		}
		var values []valueCount

		for _, binName := range binNames {
			count, _ := dist.Histogram.BinValue(binName)
			displayValue := binName

			// Prefix category values with framework name
			if labelName == LabelCategory {
				displayValue = getCategoryWithFramework(binName)
			}

			values = append(values, valueCount{
				displayValue: displayValue,
				binName:      binName,
				count:        count,
			})
		}

		// Sort by display value
		sort.Slice(values, func(i, j int) bool {
			return values[i].displayValue < values[j].displayValue
		})

		for _, v := range values {
			tableBuilder.WriteString(fmt.Sprintf("| %s | %d |\n", v.displayValue, v.count))
		}

		// Align the table and write it
		alignedTable := markdown.TableAlign(tableBuilder.String(), 1)
		fmt.Fprint(w, alignedTable)
		fmt.Fprintln(w)
	}
}

// getCategoryWithFramework returns the category name prefixed with its associated framework
func getCategoryWithFramework(category string) string {
	switch category {
	// RED framework categories
	case CategoryAvailability, CategoryLatency, CategoryThroughput:
		return "RED: " + category
	// USE framework categories
	case CategoryResource:
		return "USE: " + category
	// Custom/Business categories
	case CategoryQuality, CategoryEngagement, CategoryConversion, CategoryCost, CategorySecurity, CategoryCompliance:
		return "Custom: " + category
	default:
		return category
	}
}

// PrintReports prints human-readable table format of all reports
func PrintReports(w io.Writer, reports map[string]*MetricReport) {
	// Sort directory names for consistent output
	dirs := make([]string, 0, len(reports))
	for dir := range reports {
		dirs = append(dirs, dir)
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		PrintReport(w, reports[dir])
	}
}
