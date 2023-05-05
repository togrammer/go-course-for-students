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
	var out = in
	for _, stage := range stages {
		out = stageWithContext(ctx, out, stage)
	}

	return out
}

func stageWithContext(ctx context.Context, in In, stage Stage) Out {
	out := make(Out)
	go func() {
		defer close(out)
		for val := range in {
			select {
			case <-ctx.Done():
				return
			default:
				out <- val
			}
		}
	}()

	return stage(out)
}
