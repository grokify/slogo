package requirement

import (
	"strings"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/mogo/strconv/strconvutil"
)

type Requirements []Requirement

func (reqs Requirements) Table() *table.Table {
	tbl := table.NewTable("")
	tbl.Columns = []string{
		"Owner",
		"Name",
		"Rate Comparison",
		"Rate",
		"Duration Comparison",
		"Duration Value",
		"Duration Unit",
		"Severity",
		"Summary",
	}
	tbl.FormatMap = map[int]string{
		3: table.FormatInt,
	}
	for _, req := range reqs {
		row := []string{
			req.Owner,
			req.DisplayName,
			strings.ToUpper(string(req.Target.RateGoal.Comparison)),
			strconvutil.Ftoa(req.Target.RateGoal.Value, 2),
			req.Target.DurationComparisonString(),
			req.Target.DurationValueString(),
			req.Target.DurationUnitString(),
			string(req.Severity),
			req.Summary,
		}
		tbl.Rows = append(tbl.Rows, row)
	}
	return &tbl
}

/*
“[RateGoal]% of [events] must [meet DurationGoal condition] [over Window].”
*/
