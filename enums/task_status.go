package enums

type TaskStatus string

const (
	Open       TaskStatus = `open`
	OnProgress            = `on_progress`
	Finish                = `finish`
	Closed                = `closed`
)
