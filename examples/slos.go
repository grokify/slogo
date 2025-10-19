package examples

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

	aiagents "github.com/grokify/slogo/examples/ai-agents"
	budgetingmethod "github.com/grokify/slogo/examples/budgeting-method"
	redmetrics "github.com/grokify/slogo/examples/red-metrics"
	saascrm "github.com/grokify/slogo/examples/saas-crm"
	treatlowtraffic "github.com/grokify/slogo/examples/treat-low-traffic-as-equally-important"
	usemetrics "github.com/grokify/slogo/examples/use-metrics"
)

func ExampleSLOsByDirectory() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		"ai-agents":                              aiagents.SLOs(),
		"budgeting-method":                       budgetingmethod.SLOs(),
		"red-metrics":                            redmetrics.SLOs(),
		"saas-crm":                               saascrm.SLOs(),
		"treat-low-traffic-as-equally-important": treatlowtraffic.SLOs(),
		"use-metrics":                            usemetrics.SLOs(),
	}
}
