package enum

type TaskStatus int

const (
	StatusCreated TaskStatus = iota
	StatusInProgress
	StatusDone
	StatusDropped
)

var statusToName = map[TaskStatus]string{
	StatusCreated:    "created",
	StatusInProgress: "in_progress",
	StatusDone:       "done",
	StatusDropped:    "dropped",
}

var nameToStatus = map[string]TaskStatus{
	"created":     StatusCreated,
	"in_progress": StatusInProgress,
	"done":        StatusDone,
	"dropped":     StatusDropped,
}

func CheckTaskStatus(name string) bool {
	_, ok := nameToStatus[name]
	return ok
}

func (s TaskStatus) String() string {
	return statusToName[s]
}
