package redmetrics

import (
	"testing"

	"github.com/OpenSLO/go-sdk/pkg/openslo"
)

func TestAlertConditions(t *testing.T) {
	conditions := AlertConditions()
	if len(conditions) != 2 {
		t.Errorf("expected 2 alert conditions, got %d", len(conditions))
	}

	for _, ac := range conditions {
		if err := ac.Validate(); err != nil {
			t.Errorf("validation failed for %s: %v", ac.Metadata.Name, err)
		}
	}
}

func TestAlertNotificationTargets(t *testing.T) {
	targets := AlertNotificationTargets()
	if len(targets) != 2 {
		t.Errorf("expected 2 notification targets, got %d", len(targets))
	}

	for _, target := range targets {
		if err := target.Validate(); err != nil {
			t.Errorf("validation failed for %s: %v", target.Metadata.Name, err)
		}
	}
}

func TestAlertPolicies(t *testing.T) {
	policies := AlertPolicies()
	if len(policies) != 3 {
		t.Errorf("expected 3 alert policies, got %d", len(policies))
	}

	for _, policy := range policies {
		if err := policy.Validate(); err != nil {
			t.Errorf("validation failed for %s: %v", policy.Metadata.Name, err)
		}
	}
}

func TestAllObjects(t *testing.T) {
	objs := AllObjects()

	// Should have: 5 SLOs + 2 conditions + 2 targets + 3 policies = 12
	if len(objs) != 12 {
		t.Errorf("expected 12 objects, got %d", len(objs))
	}

	// Validate all objects
	if err := objs.Validate(); err != nil {
		t.Errorf("validation failed: %v", err)
	}

	// Count by kind
	counts := make(map[openslo.Kind]int)
	for _, obj := range objs {
		counts[obj.GetKind()]++
	}

	if counts[openslo.KindSLO] != 5 {
		t.Errorf("expected 5 SLOs, got %d", counts[openslo.KindSLO])
	}
	if counts[openslo.KindAlertCondition] != 2 {
		t.Errorf("expected 2 AlertConditions, got %d", counts[openslo.KindAlertCondition])
	}
	if counts[openslo.KindAlertNotificationTarget] != 2 {
		t.Errorf("expected 2 AlertNotificationTargets, got %d", counts[openslo.KindAlertNotificationTarget])
	}
	if counts[openslo.KindAlertPolicy] != 3 {
		t.Errorf("expected 3 AlertPolicies, got %d", counts[openslo.KindAlertPolicy])
	}
}
