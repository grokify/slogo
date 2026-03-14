package redmetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

	"github.com/grokify/slogo"
)

// AlertConditions returns standard burn rate alert conditions for SLOs.
// These implement the Google SRE multi-window, multi-burn-rate alerting approach.
func AlertConditions() []v1.AlertCondition {
	return slogo.StandardBurnRateAlerts()
}

// AlertNotificationTargets returns example notification targets.
func AlertNotificationTargets() []v1.AlertNotificationTarget {
	return []v1.AlertNotificationTarget{
		slogo.NewAlertNotificationTarget(
			"pagerduty-critical",
			"https://events.pagerduty.com/integration/XXXXXXX/enqueue",
			"PagerDuty for critical alerts - pages on-call engineer",
		),
		slogo.NewAlertNotificationTarget(
			"slack-alerts",
			"https://hooks.slack.com/services/XXXXXXXXX",
			"Slack #sre-alerts channel for visibility",
		),
	}
}

// AlertPolicies returns alert policies for the RED metrics SLOs.
func AlertPolicies() []v1.AlertPolicy {
	return []v1.AlertPolicy{
		// Availability SLO alert policy
		slogo.NewAlertPolicyBuilder("availability-alerts").
			WithDescription("Alert policy for service availability SLO").
			AlertWhenBreaching(true).
			AlertWhenResolved(true).
			AddConditionRef("fast-burn").
			AddNotificationTargetRef("pagerduty-critical").
			AddNotificationTargetRef("slack-alerts").
			Build(),

		// Latency SLO alert policy
		slogo.NewAlertPolicyBuilder("latency-alerts").
			WithDescription("Alert policy for service latency SLO").
			AlertWhenBreaching(true).
			AlertWhenResolved(true).
			AddConditionRef("slow-burn").
			AddNotificationTargetRef("slack-alerts").
			Build(),

		// Error rate SLO alert policy with inline condition
		slogo.NewAlertPolicyBuilder("error-rate-alerts").
			WithDescription("Alert policy for error rate SLO").
			AlertWhenBreaching(true).
			AddConditionInline("error-fast-burn", slogo.SeverityCritical, 10.0, "1h").
			AddNotificationTargetInline(
				"email-oncall",
				"oncall@example.com",
				"Email notification to on-call rotation",
			).
			Build(),
	}
}

// AllObjects returns all OpenSLO objects including SLOs, alert conditions,
// notification targets, and alert policies.
func AllObjects() slogo.Objects {
	var objs slogo.Objects

	// Add SLOs
	objs.AddSLOs(SLOs()...)

	// Add shared alert conditions
	objs.AddAlertConditions(AlertConditions()...)

	// Add notification targets
	objs.AddAlertNotificationTargets(AlertNotificationTargets()...)

	// Add alert policies
	objs.AddAlertPolicies(AlertPolicies()...)

	return objs
}
