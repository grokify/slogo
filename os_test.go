package slogo

import (
	"os"
	"path/filepath"
	"testing"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"
)

func TestReadWriteYAML(t *testing.T) {
	// Create test SLO
	slo := v1.NewSLO(
		v1.Metadata{
			Name:        "test-slo",
			DisplayName: "Test SLO",
		},
		v1.SLOSpec{
			Description: "Test SLO for unit testing",
			Service:     "test-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name: "test-indicator",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									"query": "sum(rate(requests_total{status=\"200\"}[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									"query": "sum(rate(requests_total[5m]))",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(4, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "High availability",
					Target:          pointer.Pointer(float64(0.999)),
					TimeSliceTarget: pointer.Pointer(float64(0.95)),
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
		},
	)

	// Create temp file
	tmpDir := t.TempDir()
	yamlFile := filepath.Join(tmpDir, "test.yaml")

	// Test WriteYAMLFile
	if err := WriteYAMLFile(yamlFile, slo); err != nil {
		t.Fatalf("WriteYAMLFile() error = %v", err)
	}

	// Test ReadYAMLFile
	objects, err := ReadYAMLFile(yamlFile)
	if err != nil {
		t.Fatalf("ReadYAMLFile() error = %v", err)
	}

	if len(objects) != 1 {
		t.Errorf("ReadYAMLFile() got %d objects, want 1", len(objects))
	}

	// Validate
	if err := objects[0].Validate(); err != nil {
		t.Errorf("objects[0].Validate() error = %v", err)
	}
}

func TestReadWriteJSON(t *testing.T) {
	// Create test SLOs
	slo1 := v1.NewSLO(
		v1.Metadata{
			Name:        "test-slo-1",
			DisplayName: "Test SLO 1",
		},
		v1.SLOSpec{
			Description: "First test SLO",
			Service:     "test-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name: "test-indicator-1",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									"query": "sum(rate(requests_total{status=\"200\"}[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									"query": "sum(rate(requests_total[5m]))",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(4, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "High availability",
					Target:          pointer.Pointer(float64(0.999)),
					TimeSliceTarget: pointer.Pointer(float64(0.95)),
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
		},
	)

	slo2 := v1.NewSLO(
		v1.Metadata{
			Name:        "test-slo-2",
			DisplayName: "Test SLO 2",
		},
		v1.SLOSpec{
			Description: "Second test SLO",
			Service:     "test-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name: "test-indicator-2",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									"query": "sum(rate(requests_total{status=\"200\"}[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									"query": "sum(rate(requests_total[5m]))",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Good availability",
					Target:          pointer.Pointer(float64(0.99)),
					TimeSliceTarget: pointer.Pointer(float64(0.90)),
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
		},
	)

	// Create temp file
	tmpDir := t.TempDir()
	jsonFile := filepath.Join(tmpDir, "test.json")

	// Test WriteJSONFile with multiple objects
	if err := WriteJSONFile(jsonFile, slo1, slo2); err != nil {
		t.Fatalf("WriteJSONFile() error = %v", err)
	}

	// Test ReadJSONFile
	objects, err := ReadJSONFile(jsonFile)
	if err != nil {
		t.Fatalf("ReadJSONFile() error = %v", err)
	}

	if len(objects) != 2 {
		t.Errorf("ReadJSONFile() got %d objects, want 2", len(objects))
	}

	// Validate both objects
	for i, obj := range objects {
		if err := obj.Validate(); err != nil {
			t.Errorf("objects[%d].Validate() error = %v", i, err)
		}
	}
}

func TestReadYAMLFileMultipleObjects(t *testing.T) {
	// Test with actual file containing multiple objects
	testFile := "examples/ai-agents/agent-availability.yaml"
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Skipf("Test file %s does not exist", testFile)
	}

	objects, err := ReadYAMLFile(testFile)
	if err != nil {
		t.Fatalf("ReadYAMLFile() error = %v", err)
	}

	if len(objects) != 2 {
		t.Errorf("ReadYAMLFile() got %d objects, want 2", len(objects))
	}

	// Validate all objects
	for i, obj := range objects {
		if err := obj.Validate(); err != nil {
			t.Errorf("objects[%d].Validate() error = %v", i, err)
		}
	}
}

func TestReadJSONFileMultipleObjects(t *testing.T) {
	// Test with actual file containing multiple objects
	testFile := "examples/ai-agents/agent-availability.json"
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Skipf("Test file %s does not exist", testFile)
	}

	objects, err := ReadJSONFile(testFile)
	if err != nil {
		t.Fatalf("ReadJSONFile() error = %v", err)
	}

	if len(objects) != 2 {
		t.Errorf("ReadJSONFile() got %d objects, want 2", len(objects))
	}

	// Validate all objects
	for i, obj := range objects {
		if err := obj.Validate(); err != nil {
			t.Errorf("objects[%d].Validate() error = %v", i, err)
		}
	}
}
