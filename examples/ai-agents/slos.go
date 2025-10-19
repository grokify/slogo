package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func SLOs() []v1.SLO {
	return []v1.SLO{
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
}
