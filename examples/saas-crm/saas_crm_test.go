package saascrm

import (
	"testing"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func TestSaaSCRMSLOs(t *testing.T) {
	sloTests := []v1.SLO{
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
	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}
