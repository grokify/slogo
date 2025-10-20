package ontology

import (
	"fmt"
	"io"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/gocharts/v2/data/histogram"
	"github.com/grokify/mogo/type/maputil"
)

func WriteSLOSetsReport(w io.Writer, slos map[string][]v1.SLO) error {
	// Get label definitions
	labelDefs := GetLabelDefinitions()

	// Write header
	fmt.Fprintln(w, "# SLO Metrics by Ontology Labels")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "This document shows the distribution of SLOs across different ontology label dimensions.")
	fmt.Fprintln(w)

	hsetsLabelSetValue := histogram.NewHistogramSets("")
	hsetsLabelSetValue.Order = Labels()

	// Get sorted label names
	labelOrder := Labels()
	// For each label type, create a table
	for _, labelName := range labelOrder {
		_, exists := labelDefs[labelName]
		if !exists {
			continue
		}

		// Get sorted set names
		setNames := maputil.Keys(slos)
		// For each set, count SLOs by label value
		for i, setName := range setNames {
			if i == 0 { // add zero columns for 0 count reporting
				if labelValues, ok := labelDefs[labelName]; ok {
					for _, labelValue := range labelValues {
						hsetsLabelSetValue.Add(labelName, setName, labelValue, 0, true)
					}
				}
			}
			slos := slos[setName]
			// Count SLOs for each label value
			for _, slo := range slos {
				if labels, exists := slo.Metadata.Labels[labelName]; exists {
					for _, labelValue := range labels {
						hsetsLabelSetValue.Add(labelName, setName, labelValue, 1, true)
					}
				}
			}
		}
	}

	hsetsLabelSetValue.UpdateSetOrders(labelDefs)

	return hsetsLabelSetValue.Markdown(w, "## ", "set name", true, &histogram.SetTablePivotOpts{
		ColTotalRight:  true,
		RowTotalBottom: true,
	})
}
