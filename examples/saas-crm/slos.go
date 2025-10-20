package saascrm

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func SLOs() []v1.SLO {
	return []v1.SLO{
		// User activation metrics
		ExampleUserActivationSLO(),
		ExampleTimeToFirstValueSLO(),
		ExampleOnboardingCompletionSLO(),
		// User engagement metrics (MAU, DAU, DAU/MAU ratio)
		ExampleDailyActiveUsersSLO(),
		ExampleMonthlyActiveUsersSLO(),
		ExampleDAUMAURatioSLO(),
		ExampleWeeklyActiveUsersSLO(),
		ExampleDAUWAURatioSLO(),
		ExamplePowerUserRatioSLO(),
		// Feature adoption metrics
		ExampleContactManagementUsageSLO(),
		ExampleDealPipelineUsageSLO(),
		ExampleEmailIntegrationUsageSLO(),
		ExampleReportingUsageSLO(),
		ExampleMobileAppUsageSLO(),
		// Business outcome metrics
		ExampleDealsCreatedSLO(),
		ExampleDealWinRateSLO(),
		ExampleSalesCycleLengthSLO(),
		ExampleContactCreationRateSLO(),
		ExampleEmailsSentPerUserSLO(),
		// User retention metrics
		ExampleDay7RetentionSLO(),
		ExampleDay30RetentionSLO(),
		ExampleChurnRateSLO(),
		ExampleResurrectionRateSLO(),
		ExampleCohortRetentionSLO(),
	}
}

func SLOsBySetSlug() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		"business-outcomes": {
			ExampleDealsCreatedSLO(),
			ExampleDealWinRateSLO(),
			ExampleSalesCycleLengthSLO(),
			ExampleContactCreationRateSLO(),
			ExampleEmailsSentPerUserSLO()},
		"feature-adoption": {
			ExampleContactManagementUsageSLO(),
			ExampleDealPipelineUsageSLO(),
			ExampleEmailIntegrationUsageSLO(),
			ExampleReportingUsageSLO(),
			ExampleMobileAppUsageSLO()},
		"user-activation": {
			ExampleUserActivationSLO(),
			ExampleTimeToFirstValueSLO(),
			ExampleOnboardingCompletionSLO()},
		"user-engagement": {
			ExampleDailyActiveUsersSLO(),
			ExampleMonthlyActiveUsersSLO(),
			ExampleDAUMAURatioSLO(),
			ExampleWeeklyActiveUsersSLO(),
			ExampleDAUWAURatioSLO(),
			ExamplePowerUserRatioSLO()},
		"user-retention": {
			ExampleDay7RetentionSLO(),
			ExampleDay30RetentionSLO(),
			ExampleChurnRateSLO(),
			ExampleResurrectionRateSLO(),
			ExampleCohortRetentionSLO()},
	}
}
