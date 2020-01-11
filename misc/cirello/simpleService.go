package cirello

import (
	"context"
	"fmt"
	"time"
)

type Simpleservice int

func (s *Simpleservice) String() string {
	return fmt.Sprintf("simple service %d", int(*s))
}

func (s *Simpleservice) Serve(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("do something...")
			time.Sleep(2000 * time.Millisecond)
		}
	}
}
