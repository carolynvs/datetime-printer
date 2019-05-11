package pretty

import (
	"math"
	"sort"
	"time"

	humanize "github.com/dustin/go-humanize"
)

// Now is an overridable Now that you can swap for testing
var Now = time.Now

const defaultDateFormat = "2006-01-02"

type Time struct {
	value          time.Time
	dateFormat     string
	timeMagnitudes func() []humanize.RelTimeMagnitude
}

func NewTime(value time.Time) Time {
	return Time{
		value:      value,
		dateFormat: defaultDateFormat,
	}
}

var prettyFormats = []humanize.RelTimeMagnitude{
	{time.Second, "now", time.Second},
	{2 * time.Second, "1 second %s", 1},
	{time.Minute, "%d seconds %s", time.Second},
	{2 * time.Minute, "1 minute %s", 1},
	{time.Hour, "%d minutes %s", time.Minute},
	{2 * time.Hour, "1 hour %s", 1},
	{humanize.Day, "%d hours %s", time.Hour},
	{math.MaxInt64, "", 1},
}

func (t Time) DefaultDateFormat() string {
	return defaultDateFormat
}

func (t Time) SetDateFormat(format string) {
	t.dateFormat = format
}

func (t Time) DefaultTimeMagnitudes() []humanize.RelTimeMagnitude {
	return []humanize.RelTimeMagnitude{
		{humanize.Week, t.dateFormat, humanize.Day},
	}
}

func (t Time) SetTimeMagnitudes(f func() []humanize.RelTimeMagnitude) {
	t.timeMagnitudes = f
}

func (t Time) TimeMagnitudes() []humanize.RelTimeMagnitude {
	if t.timeMagnitudes != nil {
		return t.timeMagnitudes()
	}
	return t.DefaultTimeMagnitudes()
}

func (t Time) String() string {
	now := Now()
	relativeResult := humanize.CustomRelTime(t.value, now, "ago", "from now", prettyFormats)
	if relativeResult != "" {
		return relativeResult
	}

	return t.FormatTimeByMagnitude(now)
}

func (t Time) FormatTimeByMagnitude(relativeTo time.Time) string {
	magnitudes := t.TimeMagnitudes()
	diff := relativeTo.Sub(t.value)

	if t.value.After(relativeTo) {
		diff = t.value.Sub(relativeTo)
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].D > diff
	})

	if n >= len(magnitudes) {
		n = len(magnitudes) - 1
	}
	mag := magnitudes[n]
	return t.value.Format(mag.Format)
}
