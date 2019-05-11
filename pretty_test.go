package pretty

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	const magicGoTime string = "Mon Jan 2 15:04:05 MST 2006"

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
	Now = func() time.Time { return now }

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			parsedTime, err := time.Parse(magicGoTime, tc.raw)
			require.NoError(t, err)

			got := NewTime(parsedTime).String()
			assert.Equal(t, tc.want, got)
		})
	}
}
