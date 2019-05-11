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

	now, _ := time.Parse(magicGoTime, "Fri May 10 15:04:05 CST 2019")
	printer.Now = func() time.Time { return now }

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			parsedTime, err := time.Parse(magicGoTime, tc.raw)
			require.NoError(t, err)

			p := printer.New(parsedTime)
			got := p.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTime_SetDateFormat(t *testing.T) {
	parsedTime, err := time.Parse(magicGoTime, "Thu May 9 15:04:05 CST 2019")
	require.NoError(t, err)

	p := printer.New(parsedTime)
	p.SetDateFormat("2006/1/2")

	got := p.String()
	want := "2019/5/9"
	assert.Equal(t, want, got)
}
