package non_db_transcation

const (
	TaskPrefix       string = "task-"
	CommitTaskPrefix string = "commit-"
	ClearTaskPrefix  string = "clear-"
)

type (
	// Event 事件类型
	Event struct {
		Key, Name string
		Value     interface{}
	}
	// EventListener
	EventListener interface {
		onEvent(event *Event)
	}
)
