package aiagents

import (
	"testing"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func TestAIAgentsSLOs(t *testing.T) {
	sloTests := []v1.SLO{
		// Availability metrics
		ExampleAgentAvailabilitySLO(),
		ExamplePerUserAgentAvailabilitySLO(),
		// Response quality metrics
		ExampleAgentResponseQualitySLO(),
		ExamplePerUserResponseQualitySLO(),
		ExampleAgentAccuracySLO(),
		// Response time metrics
		ExampleAgentResponseTimeSLO(),
		ExamplePerUserResponseTimeSLO(),
		ExampleAgentFirstTokenLatencySLO(),
		// Task completion metrics
		ExampleTaskCompletionRateSLO(),
		ExamplePerUserTaskCompletionSLO(),
		ExampleTaskAbandonmentRateSLO(),
		ExampleMultiStepTaskSuccessSLO(),
		// User engagement metrics
		ExampleDailyActiveUsersSLO(),
		ExampleUserRetentionSLO(),
		ExampleSessionDurationSLO(),
		ExampleConversationTurnsSLO(),
		// Cost efficiency metrics
		ExampleTokenUsagePerTaskSLO(),
		ExamplePerUserCostSLO(),
		ExampleCostPerSuccessfulTaskSLO(),
		ExampleCacheHitRateSLO(),
	}
	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}
