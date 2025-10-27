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
	Name             string              `json:"name" yaml:"name"`
	DisplayName      string              `json:"displayName" yaml:"displayName"`
	Description      string              `json:"description" yaml:"description"`
	Goal             float64             `json:"goal" yaml:"goal"`
	Comparison       cmputil.Operator    `json:"comparison" yaml:"comparison"` // GTE, LTE, GT, LT
	TimeTarget       duration.Spec       `json:"timeTarget" yaml:"timeTarget"`
	EvaluationWindow duration.Spec       `json:"evaluationWindow" yaml:"evaluationWindow"`
	Severity         Severity            `json:"severity" yaml:"severity"`
	Owner            string              `json:"owner" yaml:"owner"`
	Action           ActionSpec          `json:"action" yaml:"action"`
	Labels           map[string][]string `json:"labels" yaml:"labels"`
}

// ActionSpec describes the target action or process for the SLORequirement.
type ActionSpec struct {
	Type        string   `json:"type" yaml:"type"`               // e.g. "HTTP", "SQL", "Business", "Job"
	Method      string   `json:"method" yaml:"method"`           // e.g. "GET", "POST", "SELECT", "INSERT"
	Target      string   `json:"target" yaml:"target"`           // e.g. "/api/v1/users" or "SELECT * FROM users"
	Service     string   `json:"service" yaml:"service"`         // optional: logical service name
	Environment string   `json:"environment" yaml:"environment"` // optional: "prod", "staging"
	Tags        []string `json:"tags" yaml:"tags"`               // optional: arbitrary labels
}
