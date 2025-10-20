package budgetingmethod

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func SLOs() []v1.SLO {
	return []v1.SLO{
		ExampleAvailbilitySLO(),
		ExampleOccurencesSLO(),
		ExampleRatioTimeSlicesSLO(),
	}
}

func SLOsBySetSlug() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		"occurences-slo": {
			ExampleOccurencesSLO()},
		"ratio-timeslices": {
			ExampleRatioTimeSlicesSLO()},
		"timeslices-slo": {
			ExampleAvailbilitySLO()},
	}
}
