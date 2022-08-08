package timex

import (
	"time"

	"github.com/araddon/dateparse"
)

func Parse(str string) time.Time {
	t, err := dateparse.ParseAny(str)
	if err != nil {
		t2, _ := dateparse.ParseAny("0000-00-00T10:00:00.705876+08:00")
		return t2
	}
	return t
}
