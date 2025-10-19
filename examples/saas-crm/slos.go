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
