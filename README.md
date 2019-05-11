# DateTime Printer

This is a printer for Go's `time.Time` that prints the Date and Time portions of a timestamp differently
depending on how long ago the timestamp occurred.

* When it has been less than a day, only the time portion is printed: "5 minutes ago".
* When it has been more than a day, only the date portion is printed: "2019-10-03".
