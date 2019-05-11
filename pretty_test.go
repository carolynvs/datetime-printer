package printer_test

import (
	"testing"
	"time"

	printer "github.com/carolynvs/datetime-printer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const magicGoTime string = "Mon Jan 2 15:04:05 MST 2006"

func TestTime_String(t *testing.T) {
	testcases := []struct {
		name string
		raw  string
		want string
	}{
		{name: "day ago", raw: "Thu May 9 15:04:05 CST 2019", want: "2019-05-09"},
		{name: "5 hours ago", raw: "Fri May 10 10:04:05 CST 2019", want: "5 hours ago"},
		{name: "4 minutes ago", raw: "Fri May 10 15:00:05 CST 2019", want: "4 minutes ago"},
		{name: "1 second ago", raw: "Fri May 10 15:04:04 CST 2019", want: "1 second ago"},
		{name: "now", raw: "Fri May 10 15:04:05 CST 2019", want: "now"},
	}

	p := printer.DateTimePrinter{}
	now, _ := time.Parse(magicGoTime, "Fri May 10 15:04:05 CST 2019")
	p.Now = func() time.Time { return now }

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			parsedTime, err := time.Parse(magicGoTime, tc.raw)
			require.NoError(t, err)

			got := p.Format(parsedTime)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTime_SetDateFormat(t *testing.T) {
	parsedTime, err := time.Parse(magicGoTime, "Thu May 9 15:04:05 CST 2019")
	require.NoError(t, err)

	p := printer.DateTimePrinter{}
	p.DateFormat = "2006/1/2"

	got := p.Format(parsedTime)
	want := "2019/5/9"
	assert.Equal(t, want, got)
}

func TestReadme(t *testing.T) {
	now := time.Now()

	p := printer.DateTimePrinter{}

	println(p.Format(now))
	// "now"

	time.Sleep(2 * time.Second)
	println(p.Format(now))
	// "2 seconds ago"

	yearAgo := now.AddDate(-1, 0, 0)
	println(p.Format(yearAgo))
	// "2018-05-10"

	p.DateFormat = "2006/01/02"
	println(p.Format(yearAgo))
	// "2018/05/10"
}
