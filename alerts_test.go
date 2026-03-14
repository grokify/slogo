package slogo

import (
	"testing"

	"github.com/OpenSLO/go-sdk/pkg/openslo"
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func TestNewAlertCondition(t *testing.T) {
	ac := NewAlertCondition("test-condition", SeverityCritical, 14.0, "1h", "5m")

	if ac.Metadata.Name != "test-condition" {
		t.Errorf("expected name 'test-condition', got %q", ac.Metadata.Name)
	}
	if ac.Kind != openslo.KindAlertCondition {
		t.Errorf("expected kind AlertCondition, got %v", ac.Kind)
	}
	if ac.Spec.Severity != "critical" {
		t.Errorf("expected severity 'critical', got %q", ac.Spec.Severity)
	}
	if ac.Spec.Condition.Kind != v1.AlertConditionKindBurnRate {
		t.Errorf("expected condition kind 'burnrate', got %v", ac.Spec.Condition.Kind)
	}
	if *ac.Spec.Condition.Threshold != 14.0 {
		t.Errorf("expected threshold 14.0, got %v", *ac.Spec.Condition.Threshold)
	}
	if ac.Spec.Condition.LookbackWindow.String() != "1h" {
		t.Errorf("expected lookbackWindow '1h', got %v", ac.Spec.Condition.LookbackWindow)
	}
	if ac.Spec.Condition.AlertAfter == nil || ac.Spec.Condition.AlertAfter.String() != "5m" {
		t.Errorf("expected alertAfter '5m', got %v", ac.Spec.Condition.AlertAfter)
	}

	// Validate
	if err := ac.Validate(); err != nil {
		t.Errorf("validation failed: %v", err)
	}
}

func TestNewAlertNotificationTarget(t *testing.T) {
	target := NewAlertNotificationTarget("pagerduty-critical", "https://events.pagerduty.com/integration/xxx", "PagerDuty critical alerts")

	if target.Metadata.Name != "pagerduty-critical" {
		t.Errorf("expected name 'pagerduty-critical', got %q", target.Metadata.Name)
	}
	if target.Kind != openslo.KindAlertNotificationTarget {
		t.Errorf("expected kind AlertNotificationTarget, got %v", target.Kind)
	}
	if target.Spec.Target != "https://events.pagerduty.com/integration/xxx" {
		t.Errorf("unexpected target: %q", target.Spec.Target)
	}

	// Validate
	if err := target.Validate(); err != nil {
		t.Errorf("validation failed: %v", err)
	}
}

func TestAlertPolicyBuilder(t *testing.T) {
	policy := NewAlertPolicyBuilder("api-availability-alerts").
		WithDescription("Alerts for API availability SLO").
		AlertWhenBreaching(true).
		AlertWhenResolved(true).
		AddConditionInline("fast-burn", SeverityCritical, 14.0, "1h").
		AddNotificationTargetInline("slack-alerts", "https://hooks.slack.com/xxx", "Slack #alerts channel").
		Build()

	if policy.Metadata.Name != "api-availability-alerts" {
		t.Errorf("expected name 'api-availability-alerts', got %q", policy.Metadata.Name)
	}
	if policy.Kind != openslo.KindAlertPolicy {
		t.Errorf("expected kind AlertPolicy, got %v", policy.Kind)
	}
	if policy.Spec.Description != "Alerts for API availability SLO" {
		t.Errorf("unexpected description: %q", policy.Spec.Description)
	}
	if !policy.Spec.AlertWhenBreaching {
		t.Error("expected AlertWhenBreaching to be true")
	}
	if !policy.Spec.AlertWhenResolved {
		t.Error("expected AlertWhenResolved to be true")
	}
	if len(policy.Spec.Conditions) != 1 {
		t.Errorf("expected 1 condition, got %d", len(policy.Spec.Conditions))
	}
	if len(policy.Spec.NotificationTargets) != 1 {
		t.Errorf("expected 1 notification target, got %d", len(policy.Spec.NotificationTargets))
	}

	// Validate
	if err := policy.Validate(); err != nil {
		t.Errorf("validation failed: %v", err)
	}
}

func TestAlertPolicyBuilderWithRefs(t *testing.T) {
	policy := NewAlertPolicyBuilder("api-alerts").
		AddConditionRef("shared-fast-burn").
		AddNotificationTargetRef("shared-pagerduty").
		Build()

	if len(policy.Spec.Conditions) != 1 {
		t.Errorf("expected 1 condition, got %d", len(policy.Spec.Conditions))
	}
	if policy.Spec.Conditions[0].AlertPolicyConditionRef == nil {
		t.Error("expected condition ref to be set")
	}
	if policy.Spec.Conditions[0].AlertPolicyConditionRef.ConditionRef != "shared-fast-burn" {
		t.Errorf("unexpected condition ref: %q", policy.Spec.Conditions[0].AlertPolicyConditionRef.ConditionRef)
	}

	if len(policy.Spec.NotificationTargets) != 1 {
		t.Errorf("expected 1 notification target, got %d", len(policy.Spec.NotificationTargets))
	}
	if policy.Spec.NotificationTargets[0].AlertPolicyNotificationTargetRef == nil {
		t.Error("expected target ref to be set")
	}
	if policy.Spec.NotificationTargets[0].AlertPolicyNotificationTargetRef.TargetRef != "shared-pagerduty" {
		t.Errorf("unexpected target ref: %q", policy.Spec.NotificationTargets[0].AlertPolicyNotificationTargetRef.TargetRef)
	}

	// Validate
	if err := policy.Validate(); err != nil {
		t.Errorf("validation failed: %v", err)
	}
}

func TestStandardBurnRateAlerts(t *testing.T) {
	alerts := StandardBurnRateAlerts()

	if len(alerts) != 2 {
		t.Errorf("expected 2 standard alerts, got %d", len(alerts))
	}

	// Fast burn
	if alerts[0].Metadata.Name != "fast-burn" {
		t.Errorf("expected 'fast-burn', got %q", alerts[0].Metadata.Name)
	}
	if *alerts[0].Spec.Condition.Threshold != 14.0 {
		t.Errorf("expected threshold 14.0, got %v", *alerts[0].Spec.Condition.Threshold)
	}

	// Slow burn
	if alerts[1].Metadata.Name != "slow-burn" {
		t.Errorf("expected 'slow-burn', got %q", alerts[1].Metadata.Name)
	}
	if *alerts[1].Spec.Condition.Threshold != 1.0 {
		t.Errorf("expected threshold 1.0, got %v", *alerts[1].Spec.Condition.Threshold)
	}

	// Validate both
	for _, alert := range alerts {
		if err := alert.Validate(); err != nil {
			t.Errorf("validation failed for %s: %v", alert.Metadata.Name, err)
		}
	}
}

func TestObjectsAddAlerts(t *testing.T) {
	var objs Objects

	// Add alert conditions
	objs.AddAlertConditions(StandardBurnRateAlerts()...)
	if len(objs) != 2 {
		t.Errorf("expected 2 objects, got %d", len(objs))
	}

	// Add notification target
	objs.AddAlertNotificationTargets(
		NewAlertNotificationTarget("slack", "https://hooks.slack.com/xxx", "Slack alerts"),
	)
	if len(objs) != 3 {
		t.Errorf("expected 3 objects, got %d", len(objs))
	}

	// Add alert policy
	policy := NewAlertPolicyBuilder("test-policy").
		AddConditionRef("fast-burn").
		AddNotificationTargetRef("slack").
		Build()
	objs.AddAlertPolicies(policy)
	if len(objs) != 4 {
		t.Errorf("expected 4 objects, got %d", len(objs))
	}

	// Validate all
	if err := objs.Validate(); err != nil {
		t.Errorf("validation failed: %v", err)
	}
}
