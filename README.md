# DateTime Printer

This is a printer for Go's `time.Time` that prints the Date and Time portions of a timestamp differently
depending on how long ago the timestamp occurred.

* When it has been less than a day, only the time portion is printed: "5 minutes ago".
* When it has been more than a day, only the date portion is printed: "2019-10-03".


## Example

```go
package main

import (
	"time"
	
	printer "github.com/carolynvs/datetime-printer"
)

func main() {
now := time.Now()

	p := printer.DateTimePrinter{}

	println(p.Format(now))
	// "now"

	time.Sleep(2*time.Second)
	println(p.Format(now))
	// "2 seconds ago"

	yearAgo := now.AddDate(-1, 0, 0)
	println(p.Format(yearAgo))
	// "2018-05-10"

	p.DateFormat = "2006/01/02"
	println(p.Format(yearAgo))
	// "2018/05/10"
}
```