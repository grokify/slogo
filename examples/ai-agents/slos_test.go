package aiagents

import (
	_ "embed"
	"testing"

	"github.com/grokify/slogo"
)

func TestAIAgentsSLOs(t *testing.T) {
	sloTests := SLOs()

	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}

//go:embed agent-availability.json
var agentAvailabilityJSON []byte

//go:embed agent-availability.yaml
var agentAvailabilityYAML []byte

func TestReadExamples(t *testing.T) {
	var readExamplesTests = []struct {
		v        []byte
		filename string
		count    int
	}{
		{agentAvailabilityJSON, "agent-availability.json", 2},
		{agentAvailabilityYAML, "agent-availability.yaml", 2},
	}

	for _, tt := range readExamplesTests {
		objs, err := slogo.ReadFile(tt.filename)
		if err != nil {
			t.Errorf("slogo.Read(\"%s\") error [%s]", tt.filename, err.Error())
		}
		if len(objs) != tt.count {
			t.Errorf("slogo.Read(\"%s\") mismatch on count want [%d] got [%d]", tt.filename, tt.count, len(objs))
		}
	}
}
