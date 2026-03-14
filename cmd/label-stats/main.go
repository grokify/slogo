package main

import (
	"encoding/json"
	"fmt"
	"os"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

	aiagents "github.com/grokify/slogo/examples/ai-agents"
	budgetingmethod "github.com/grokify/slogo/examples/budgeting-method"
	redmetrics "github.com/grokify/slogo/examples/red-metrics"
	saascrm "github.com/grokify/slogo/examples/saas-crm"
	usemetrics "github.com/grokify/slogo/examples/use-metrics"
	"github.com/grokify/slogo/ontology"
)

func main() {
	reports := make(map[string]*ontology.MetricReport)

	// Analyze each directory
	reports["examples/budgeting-method"] = ontology.AnalyzeLabels("examples/budgeting-method", getBudgetingMethodSLOs())
	reports["examples/red-metrics"] = ontology.AnalyzeLabels("examples/red-metrics", getREDMetricsSLOs())
	reports["examples/use-metrics"] = ontology.AnalyzeLabels("examples/use-metrics", getUSEMetricsSLOs())
	reports["examples/ai-agents"] = ontology.AnalyzeLabels("examples/ai-agents", getAIAgentsSLOs())
	reports["examples/saas-crm"] = ontology.AnalyzeLabels("examples/saas-crm", getSaaSCRMSLOs())

	// Output as JSON
	output, err := json.MarshalIndent(reports, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

func getBudgetingMethodSLOs() []v1.SLO {
	return []v1.SLO{
		budgetingmethod.ExampleOccurrencesSLO(),
		budgetingmethod.ExampleAvailbilitySLO(),
		budgetingmethod.ExampleRatioTimeSlicesSLO(),
	}
}

func getREDMetricsSLOs() []v1.SLO {
	return []v1.SLO{
		redmetrics.ExampleRateSLO(),
		redmetrics.ExampleErrorRateSLO(),
		redmetrics.ExampleDurationSLO(),
		redmetrics.ExampleDurationP99SLO(),
		redmetrics.ExampleAvailabilitySLO(),
	}
}

func getUSEMetricsSLOs() []v1.SLO {
	return []v1.SLO{
		usemetrics.ExampleCPUUtilizationSLO(),
		usemetrics.ExampleMemoryUtilizationSLO(),
		usemetrics.ExampleDiskUtilizationSLO(),
		usemetrics.ExampleCPUSaturationSLO(),
		usemetrics.ExampleMemorySaturationSLO(),
		usemetrics.ExampleDiskSaturationSLO(),
		usemetrics.ExampleNetworkSaturationSLO(),
		usemetrics.ExampleDiskErrorsSLO(),
		usemetrics.ExampleNetworkErrorsSLO(),
		usemetrics.ExampleMemoryErrorsSLO(),
		usemetrics.ExampleCPUThrottlingErrorsSLO(),
	}
}

func getAIAgentsSLOs() []v1.SLO {
	return []v1.SLO{
		aiagents.ExampleAgentAvailabilitySLO(),
		aiagents.ExamplePerUserAgentAvailabilitySLO(),
		aiagents.ExampleAgentResponseQualitySLO(),
		aiagents.ExamplePerUserResponseQualitySLO(),
		aiagents.ExampleAgentAccuracySLO(),
		aiagents.ExampleAgentResponseTimeSLO(),
		aiagents.ExamplePerUserResponseTimeSLO(),
		aiagents.ExampleAgentFirstTokenLatencySLO(),
		aiagents.ExampleTaskCompletionRateSLO(),
		aiagents.ExamplePerUserTaskCompletionSLO(),
		aiagents.ExampleTaskAbandonmentRateSLO(),
		aiagents.ExampleMultiStepTaskSuccessSLO(),
		aiagents.ExampleDailyActiveUsersSLO(),
		aiagents.ExampleUserRetentionSLO(),
		aiagents.ExampleSessionDurationSLO(),
		aiagents.ExampleConversationTurnsSLO(),
		aiagents.ExampleTokenUsagePerTaskSLO(),
		aiagents.ExamplePerUserCostSLO(),
		aiagents.ExampleCostPerSuccessfulTaskSLO(),
		aiagents.ExampleCacheHitRateSLO(),
	}
}

func getSaaSCRMSLOs() []v1.SLO {
	return []v1.SLO{
		saascrm.ExampleUserActivationSLO(),
		saascrm.ExampleTimeToFirstValueSLO(),
		saascrm.ExampleOnboardingCompletionSLO(),
		saascrm.ExampleDailyActiveUsersSLO(),
		saascrm.ExampleMonthlyActiveUsersSLO(),
		saascrm.ExampleDAUMAURatioSLO(),
		saascrm.ExampleWeeklyActiveUsersSLO(),
		saascrm.ExampleDAUWAURatioSLO(),
		saascrm.ExamplePowerUserRatioSLO(),
		saascrm.ExampleContactManagementUsageSLO(),
		saascrm.ExampleDealPipelineUsageSLO(),
		saascrm.ExampleEmailIntegrationUsageSLO(),
		saascrm.ExampleReportingUsageSLO(),
		saascrm.ExampleMobileAppUsageSLO(),
		saascrm.ExampleDealsCreatedSLO(),
		saascrm.ExampleDealWinRateSLO(),
		saascrm.ExampleSalesCycleLengthSLO(),
		saascrm.ExampleContactCreationRateSLO(),
		saascrm.ExampleEmailsSentPerUserSLO(),
		saascrm.ExampleDay7RetentionSLO(),
		saascrm.ExampleDay30RetentionSLO(),
		saascrm.ExampleChurnRateSLO(),
		saascrm.ExampleResurrectionRateSLO(),
		saascrm.ExampleCohortRetentionSLO(),
	}
}
