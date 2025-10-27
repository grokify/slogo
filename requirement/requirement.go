package requirement

import (
	"github.com/grokify/mogo/cmp/cmputil"
	"github.com/grokify/mogo/time/duration"
)

type Severity string

const (
	ActionTypeHTTP = "HTTP"
	ActionTypeSQL  = "SQL"

	SeverityCritical      Severity = "critical"
	SeverityHigh          Severity = "high"
	SeverityMedium        Severity = "medium"
	SeverityLow           Severity = "low"
	SeverityInformational Severity = "informational"
	SeverityUnknown       Severity = "unknown"
)

type Requirement struct {
	Name             string              `json:"name"`
	DisplayName      string              `json:"displayName"`
	Description      string              `json:"description"`
	Goal             float64             `json:"goal"`
	Comparison       cmputil.Operator    `json:"comparison"` // GTE, LTE, GT, LT
	TimeTarget       duration.Spec       `json:"timeTarget"`
	EvaluationWindow duration.Spec       `json:"evaluationWindow"`
	Severity         Severity            `json:"severity"`
	Owner            string              `json:"owner"`
	Action           ActionSpec          `json:"action"`
	Labels           map[string][]string `json:"labels"`
}

// ActionSpec describes the target action or process for the SLORequirement.
type ActionSpec struct {
	Type        string   `json:"type"`        // e.g. "HTTP", "SQL", "Business", "Job"
	Method      string   `json:"method"`      // e.g. "GET", "POST", "SELECT", "INSERT"
	Target      string   `json:"target"`      // e.g. "/api/v1/users" or "SELECT * FROM users"
	Service     string   `json:"service"`     // optional: logical service name
	Environment string   `json:"environment"` // optional: "prod", "staging"
	Tags        []string `json:"tags"`        // optional: arbitrary labels
}
