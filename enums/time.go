package enums

const (
	DDMMYY = "2006-01-02"
)

type TimeDuration int

const (
	YEAR TimeDuration = iota + 1
	MONTH
	DAY
	HOUR
	MINUTE
	SECOND
	MILLISECOND
)
