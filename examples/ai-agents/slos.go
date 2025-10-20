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

func SLOsBySetSlug() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		"agent-availability": {
			ExampleAgentAvailabilitySLO(),
			ExamplePerUserAgentAvailabilitySLO()},
		"cost-efficiency": {
			ExampleTokenUsagePerTaskSLO(),
			ExamplePerUserCostSLO(),
			ExampleCostPerSuccessfulTaskSLO(),
			ExampleCacheHitRateSLO()},
		"response-quality": {
			ExampleAgentResponseQualitySLO(),
			ExamplePerUserResponseQualitySLO(),
			ExampleAgentAccuracySLO()},
		"response-time": {
			ExampleAgentResponseTimeSLO(),
			ExamplePerUserResponseTimeSLO(),
			ExampleAgentFirstTokenLatencySLO()},
		"task-completion": {
			ExampleTaskCompletionRateSLO(),
			ExamplePerUserTaskCompletionSLO(),
			ExampleTaskAbandonmentRateSLO(),
			ExampleMultiStepTaskSuccessSLO()},
		"user-engagement": {
			ExampleDailyActiveUsersSLO(),
			ExampleUserRetentionSLO(),
			ExampleSessionDurationSLO(),
			ExampleConversationTurnsSLO()},
	}
}
