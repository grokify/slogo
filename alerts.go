package slogo

import (
	"github.com/OpenSLO/go-sdk/pkg/openslo"
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"
)

// AlertSeverity represents standard alert severity levels.
type AlertSeverity string

const (
	SeverityCritical AlertSeverity = "critical"
	SeverityHigh     AlertSeverity = "high"
	SeverityMedium   AlertSeverity = "medium"
	SeverityLow      AlertSeverity = "low"
	SeverityInfo     AlertSeverity = "info"
)

// NewAlertCondition creates a new AlertCondition with burn rate configuration.
// lookbackWindow and alertAfter are duration strings like "1h", "5m", "1d".
func NewAlertCondition(name string, severity AlertSeverity, threshold float64, lookbackWindow, alertAfter string) v1.AlertCondition {
	lookback, err := v1.ParseDurationShorthand(lookbackWindow)
	if err != nil {
		panic("invalid lookbackWindow: " + err.Error())
	}

	spec := v1.AlertConditionSpec{
		Severity: string(severity),
		Condition: v1.AlertConditionType{
			Kind:           v1.AlertConditionKindBurnRate,
			Operator:       v1.OperatorGTE,
			Threshold:      pointer.Pointer(threshold),
			LookbackWindow: lookback,
		},
	}
	if alertAfter != "" {
		aa, err := v1.ParseDurationShorthand(alertAfter)
		if err != nil {
			panic("invalid alertAfter: " + err.Error())
		}
		spec.Condition.AlertAfter = &aa
	}
	return v1.NewAlertCondition(v1.Metadata{Name: name}, spec)
}

// NewAlertNotificationTarget creates a new AlertNotificationTarget.
func NewAlertNotificationTarget(name, target, description string) v1.AlertNotificationTarget {
	return v1.NewAlertNotificationTarget(
		v1.Metadata{Name: name},
		v1.AlertNotificationTargetSpec{
			Target:      target,
			Description: description,
		},
	)
}

// AlertPolicyBuilder provides a fluent interface for building AlertPolicy objects.
type AlertPolicyBuilder struct {
	name        string
	description string
	labels      v1.Labels

	alertWhenNoData    bool
	alertWhenBreaching bool
	alertWhenResolved  bool

	conditions          []v1.AlertPolicyCondition
	notificationTargets []v1.AlertPolicyNotificationTarget
}

// NewAlertPolicyBuilder creates a new AlertPolicyBuilder.
func NewAlertPolicyBuilder(name string) *AlertPolicyBuilder {
	return &AlertPolicyBuilder{
		name:               name,
		alertWhenBreaching: true, // Default to alerting when breaching
	}
}

// WithDescription sets the alert policy description.
func (b *AlertPolicyBuilder) WithDescription(desc string) *AlertPolicyBuilder {
	b.description = desc
	return b
}

// WithLabels sets the alert policy labels.
func (b *AlertPolicyBuilder) WithLabels(labels v1.Labels) *AlertPolicyBuilder {
	b.labels = labels
	return b
}

// AlertWhenNoData configures alerting when no data is available.
func (b *AlertPolicyBuilder) AlertWhenNoData(v bool) *AlertPolicyBuilder {
	b.alertWhenNoData = v
	return b
}

// AlertWhenBreaching configures alerting when SLO is breaching.
func (b *AlertPolicyBuilder) AlertWhenBreaching(v bool) *AlertPolicyBuilder {
	b.alertWhenBreaching = v
	return b
}

// AlertWhenResolved configures alerting when SLO breach is resolved.
func (b *AlertPolicyBuilder) AlertWhenResolved(v bool) *AlertPolicyBuilder {
	b.alertWhenResolved = v
	return b
}

// AddConditionRef adds a reference to an external AlertCondition.
func (b *AlertPolicyBuilder) AddConditionRef(ref string) *AlertPolicyBuilder {
	b.conditions = append(b.conditions, v1.AlertPolicyCondition{
		AlertPolicyConditionRef: &v1.AlertPolicyConditionRef{
			ConditionRef: ref,
		},
	})
	return b
}

// AddConditionInline adds an inline AlertCondition.
// lookbackWindow is a duration string like "1h", "5m", "1d".
func (b *AlertPolicyBuilder) AddConditionInline(name string, severity AlertSeverity, threshold float64, lookbackWindow string) *AlertPolicyBuilder {
	lookback, err := v1.ParseDurationShorthand(lookbackWindow)
	if err != nil {
		panic("invalid lookbackWindow: " + err.Error())
	}

	b.conditions = append(b.conditions, v1.AlertPolicyCondition{
		AlertPolicyConditionInline: &v1.AlertPolicyConditionInline{
			Kind:     openslo.KindAlertCondition,
			Metadata: v1.Metadata{Name: name},
			Spec: v1.AlertConditionSpec{
				Severity: string(severity),
				Condition: v1.AlertConditionType{
					Kind:           v1.AlertConditionKindBurnRate,
					Operator:       v1.OperatorGTE,
					Threshold:      pointer.Pointer(threshold),
					LookbackWindow: lookback,
				},
			},
		},
	})
	return b
}

// AddNotificationTargetRef adds a reference to an external AlertNotificationTarget.
func (b *AlertPolicyBuilder) AddNotificationTargetRef(ref string) *AlertPolicyBuilder {
	b.notificationTargets = append(b.notificationTargets, v1.AlertPolicyNotificationTarget{
		AlertPolicyNotificationTargetRef: &v1.AlertPolicyNotificationTargetRef{
			TargetRef: ref,
		},
	})
	return b
}

// AddNotificationTargetInline adds an inline AlertNotificationTarget.
func (b *AlertPolicyBuilder) AddNotificationTargetInline(name, target, description string) *AlertPolicyBuilder {
	b.notificationTargets = append(b.notificationTargets, v1.AlertPolicyNotificationTarget{
		AlertPolicyNotificationTargetInline: &v1.AlertPolicyNotificationTargetInline{
			Kind:     openslo.KindAlertNotificationTarget,
			Metadata: v1.Metadata{Name: name},
			Spec: v1.AlertNotificationTargetSpec{
				Target:      target,
				Description: description,
			},
		},
	})
	return b
}

// Build creates the AlertPolicy.
func (b *AlertPolicyBuilder) Build() v1.AlertPolicy {
	return v1.NewAlertPolicy(
		v1.Metadata{
			Name:   b.name,
			Labels: b.labels,
		},
		v1.AlertPolicySpec{
			Description:         b.description,
			AlertWhenNoData:     b.alertWhenNoData,
			AlertWhenBreaching:  b.alertWhenBreaching,
			AlertWhenResolved:   b.alertWhenResolved,
			Conditions:          b.conditions,
			NotificationTargets: b.notificationTargets,
		},
	)
}

// StandardBurnRateAlerts returns standard multi-window burn rate alert conditions.
// This implements the Google SRE recommended approach with fast and slow burn rates.
//
// Fast burn (14x): 1h lookback, alerts quickly for severe issues
// Slow burn (1x): 6h lookback, catches gradual degradation
func StandardBurnRateAlerts() []v1.AlertCondition {
	return []v1.AlertCondition{
		// Fast burn: 14x burn rate over 1 hour (will exhaust 30-day budget in ~2 days)
		NewAlertCondition("fast-burn", SeverityCritical, 14.0, "1h", "5m"),
		// Slow burn: 1x burn rate over 6 hours (on track to exhaust budget)
		NewAlertCondition("slow-burn", SeverityHigh, 1.0, "6h", "30m"),
	}
}

// AddAlertPolicies adds AlertPolicy objects to the Objects slice.
func (objs *Objects) AddAlertPolicies(policies ...v1.AlertPolicy) {
	for _, p := range policies {
		*objs = append(*objs, p)
	}
}

// AddAlertConditions adds AlertCondition objects to the Objects slice.
func (objs *Objects) AddAlertConditions(conditions ...v1.AlertCondition) {
	for _, c := range conditions {
		*objs = append(*objs, c)
	}
}

// AddAlertNotificationTargets adds AlertNotificationTarget objects to the Objects slice.
func (objs *Objects) AddAlertNotificationTargets(targets ...v1.AlertNotificationTarget) {
	for _, t := range targets {
		*objs = append(*objs, t)
	}
}
