package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func wrapChannelWithDone(in In, done In, readAfterDone bool) In {
	inWrapper := make(Bi)
	go func() {
	loop:
		for {
			select {
			case <-done:
				break loop
			default:
			}

			var task interface{}
			var ok bool
			select {
			case <-done:
				break loop
			case task, ok = <-in:
				if !ok {
					break loop
				}
			}

			select {
			case <-done:
				break loop
			case inWrapper <- task:
			}
		}

		if readAfterDone {
			for range in {
			}
		}

		close(inWrapper)
	}()
	return inWrapper
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	lastOut := in
	for _, stage := range stages {
		inWrapper := wrapChannelWithDone(lastOut, done, false)
		lastOut = stage(inWrapper)
		lastOut = wrapChannelWithDone(lastOut, done, true)
	}

	return lastOut
}
