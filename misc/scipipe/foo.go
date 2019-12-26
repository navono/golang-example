package scipipe

import (
	"bytes"

	. "github.com/scipipe/scipipe"
)

type Fooer struct {
	*Process
	name string
}

func NewFooer(wf *Workflow, name string) *Fooer {
	// Initiate task from a "shell like" pattern, though here we
	// just specify the out-port foo
	innerFoo := NewProc(wf, "fooer", "{o:foo}")
	// Set the output formatter to a static string
	innerFoo.SetOut("foo", "foo.txt")
	// Create the custom execute function, with pure Go code
	innerFoo.CustomExecute = func(task *Task) {
		task.OutIP("foo").Write([]byte("foo\n"))
	}
	// Connect the ports of the outer task to the inner, generic one
	fooer := &Fooer{
		innerFoo,
		name,
	}
	return fooer
}

func (p *Fooer) OutFoo() *OutPort { return p.Out("foo") }

// Foo2Barer

type Foo2Barer struct {
	*Process
	name string
}

func NewFoo2Barer(wf *Workflow, name string) *Foo2Barer {
	// Initiate task from a "shell like" pattern, though here we
	// just specify the in-port foo and the out-port bar
	innerProc := NewProc(wf, "foo2bar", "{i:foo}{o:bar}")
	// Set the output formatter to extend the path on the "bar"" in-port
	innerProc.SetOut("bar", "{i:foo}.bar.txt")
	// Create the custom execute function, with pure Go code
	innerProc.CustomExecute = func(task *Task) {
		task.OutIP("bar").Write(bytes.Replace(task.InIP("foo").Read(), []byte("foo"), []byte("bar"), 1))
	}

	// Connect the ports of the outer task to the inner, generic one
	return &Foo2Barer{
		innerProc,
		name,
	}
}

func (p *Foo2Barer) InFoo() *InPort   { return p.In("foo") }
func (p *Foo2Barer) OutBar() *OutPort { return p.Out("bar") }
