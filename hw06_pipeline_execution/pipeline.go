package hw06pipelineexecution

//import "fmt"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// add safe stop for stage by done channel, with out channel full read
func stageWrapper(in In, done In, stage Stage) Out {
	inWrapper := make(Bi)
	out := stage(inWrapper)
	outWrapper := make(Bi)

	go func() {
		// inputWrapper
		go func() {
inputLoop:
			for {
				select {
				case <-done:
					break inputLoop
				default:
				}

				var task interface{}
				var ok bool
				select {
				case <-done:
					break inputLoop
				case task, ok = <-in:
					if !ok {
						break inputLoop
					}
				}

				select {
				case <-done:
					break inputLoop
				case inWrapper <- task:
				}
			}
			close(inWrapper)
		}()

		// outputWrapper
		go func() {
outputLoop:
			for {
				select {
				case <-done:
					break outputLoop
				default:
				}

				var result interface{}
				var ok bool
				select {
				case <-done:
					break outputLoop
				case result, ok = <-out:
					if !ok {
						break outputLoop
					}
				}

				select {
				case <-done:
					break outputLoop
				case outWrapper <- result:
				}
			}

			for _ = range out {
			}
			close(outWrapper)
		}()
	}()
	return outWrapper
}


func ExecutePipeline(in In, done In, stages ...Stage) Out {
	lastOut := in
	for _, stage := range stages {
		lastOut = stageWrapper(lastOut, done, stage)
	}

	resultWrapper := make(Bi) 
	go func() {
loop:
		for {
			select {
			case <-done:
				break loop
			default:
			}

			var result interface{}
			var ok bool
			select {
			case <-done:
				break loop
			case result, ok = <-lastOut:
				if !ok {
					break loop
				}
			}

			select {
			case <-done:
				break loop
			case resultWrapper <- result:
			}
		}

		close(resultWrapper)
	}()

	return resultWrapper
}
