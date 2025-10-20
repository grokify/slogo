package examples

import (
	"path/filepath"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

	"github.com/grokify/slogo"
	aiagents "github.com/grokify/slogo/examples/ai-agents"
	budgetingmethod "github.com/grokify/slogo/examples/budgeting-method"
	redmetrics "github.com/grokify/slogo/examples/red-metrics"
	saascrm "github.com/grokify/slogo/examples/saas-crm"
	lowtraffic "github.com/grokify/slogo/examples/treat-low-traffic-as-equally-important"
	usemetrics "github.com/grokify/slogo/examples/use-metrics"
)

const (
	SetAIAgents        = "ai-agents"
	SetBudgetingMethod = "budgeting-method"
	SetREDMetrics      = "red-metrics"
	SetSAASCRM         = "saas-crm"
	SetLowTraffic      = "treat-low-traffic-as-equally-important"
	SetUSEMetrics      = "use-metrics"
)

func ExampleSLOsByDirectory() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		SetAIAgents:        aiagents.SLOs(),
		SetBudgetingMethod: budgetingmethod.SLOs(),
		SetREDMetrics:      redmetrics.SLOs(),
		SetSAASCRM:         saascrm.SLOs(),
		SetLowTraffic:      lowtraffic.SLOs(),
		SetUSEMetrics:      usemetrics.SLOs(),
	}
}

func SLOsByDirAndFileSlug() map[string]map[string][]v1.SLO {
	return map[string]map[string][]v1.SLO{
		SetAIAgents:        aiagents.SLOsBySetSlug(),
		SetBudgetingMethod: budgetingmethod.SLOsBySetSlug(),
		SetREDMetrics:      redmetrics.SLOsBySetSlug(),
		SetSAASCRM:         saascrm.SLOsBySetSlug(),
		SetLowTraffic:      lowtraffic.SLOsBySetSlug(),
		SetUSEMetrics:      usemetrics.SLOsBySetSlug(),
	}
}

func WriteExampleJSONAndYAMLFiles(examplesDir string) error {
	m := SLOsByDirAndFileSlug()
	for setSlug, mapFileSlugSLOs := range m {
		if setSlug == SetBudgetingMethod || setSlug == SetLowTraffic {
			continue
		}
		for fileSlug, slos := range mapFileSlugSLOs {
			fpJSON := filepath.Join(examplesDir, setSlug, fileSlug+".json")
			fpYAML := filepath.Join(examplesDir, setSlug, fileSlug+".yaml")
			objs := slogo.Objects{}
			objs.AddSLOs(slos...)
			// sloObjs := slogo.Objects(slos)
			if err := objs.WriteFileJSON(fpJSON); err != nil {
				return err
			}
			if err := objs.WriteFileYAML(fpYAML); err != nil {
				return err
			}
		}
	}
	return nil
}
