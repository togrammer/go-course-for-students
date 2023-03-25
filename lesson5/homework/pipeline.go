package executor

import (
	"context"
)

type (
	In  = chan any
	Out = In
)

type Stage func(in In) (out Out)

func ExecutePipeline(ctx context.Context, in In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = stage(out)
	}

	go func() {
		<-ctx.Done()
		close(out)
	}()

	return out
}
