package job

type (
	Task interface {
		out
		Execute() *Result
		Clear()
	}

	TaskContext struct {
		ctx *taskContextVal
	}

	Result struct {
		Error error
		Data  interface{}
	}

	out interface {
		Cancel()
		getCtx() *taskContextVal
		setCtx(ctx *taskContextVal)
	}

	taskContextVal struct {
		name   string
		index  int
		cancel chan<- interface{}
	}
)

func (tc *TaskContext) getCtx() *taskContextVal {
	return tc.ctx
}

func (tc *TaskContext) setCtx(ctx *taskContextVal) {
	tc.ctx = ctx
}

func (tc *TaskContext) Cancel() {
	tc.ctx.cancel <- true
}
