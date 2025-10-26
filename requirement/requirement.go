package requirement

import "github.com/grokify/mogo/time/duration"

const (
	ActionTypeHTTP = "HTTP"
	ActionTypeSQL  = "SQL"
)

type Requirement struct {
	Name             string        `json:"name"`
	DisplayName      string        `json:"displayName"`
	Goal             float64       `json:"goal"`
	Comparison       string        `json:"comparison"` // GTE, LTE, GT, LT
	TimeTarget       duration.Spec `json:"timeTarget"`
	EvaluationWindow duration.Spec `json:"evaluationWindow"`
	Severity         string        `json:"severity"`
	Owner            string        `json:"owner"`
	Description      string        `json:"description"`
	Action           ActionSpec    `json:"action"`
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
