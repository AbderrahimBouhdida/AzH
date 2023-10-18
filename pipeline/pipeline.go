package pipeline

import (
	"encoding/json"
	"reflect"
	"sync"
	"time"

	"https://github.com/AbderrahimBouhdida/AzH/internal"
)

type Result[T any] struct {
	Error error
	Ok    T
}

// OrDone provides an explicit cancellation mechanism to ensure the encapsulated and downstream goroutines are cleaned
// up. This frees the caller from depending on the input channel to close in order to free the goroutine, thus
// preventing possible leaks.
func OrDone[D, T any](done <-chan D, in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok {
					return
				} else {
					select {
					case out <- val:
					case <-done:
					}
				}
			}
		}
	}()
	return out
}

// Mux joins multiple channels and returns a channel as single stream of data.
func Mux[D any](done <-chan D, channels ...<-chan any) <-chan any {
	var wg sync.WaitGroup
	out := make(chan interface{})

	muxer := func(channel <-chan any) {
		defer wg.Done()
		for item := range OrDone(done, channel) {
			out <- item
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go muxer(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Demux distributes the stream of data from a single channel across multiple channels to parallelize CPU use and I/O
func Demux[D, T any](done <-chan D, in <-chan T, size int) []<-chan T {
	outputs := make([]chan T, size)

	for i := range outputs {
		outputs[i] = make(chan T)
	}

	closeOutputs := func() {
		for i := range outputs {
			close(outputs[i])
		}
	}

	cases := internal.Map(outputs, func(out chan T) reflect.SelectCase {
		return reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(out),
		}
	})

	go func() {
		defer closeOutputs()
		for item := range OrDone(done, in) {
			// send item to exactly one channel
			for i := range cases {
				cases[i].Send = reflect.ValueOf(item)
			}
			reflect.Select(cases)
		}
	}()

	return internal.Map(outputs, func(out chan T) <-chan T { return out })
}

func ToAny[D, T any](done <-chan D, in <-chan T) <-chan any {
	return Map(done, in, func(t T) any {
		return any(t)
	})
}

func Map[D, T, U any](done <-chan D, in <-chan T, fn func(T) U) <-chan U {
	out := make(chan U)
	go func() {
		defer close(out)
		for item := range OrDone(done, in) {
			out <- fn(item)
		}
	}()
	return out
}

func Filter[D, T any](done <-chan D, in <-chan T, fn func(T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for item := range OrDone(done, in) {
			if fn(item) {
				out <- item
			}
		}
	}()
	return out
}

// Tee copies the stream of data from a single channel to zero or more channels
func Tee[D, T any](done <-chan D, in <-chan T, outputs ...chan T) {
	go func() {
		// Need to close outputs when goroutine exits to ensure we avoid deadlock
		defer func() {
			for i := range outputs {
				close(outputs[i])
			}
		}()

		for item := range OrDone(done, in) {
			for _, out := range outputs {
				select {
				case <-done:
				case out <- item:
				}
			}
		}
	}()
}

func TeeFixed[D, T any](done <-chan D, in <-chan T, size int) []<-chan T {
	out := internal.Map(make([]any, size), func(_ any) chan T {
		return make(chan T)
	})
	Tee(done, in, out...)
	return internal.Map(out, func(c chan T) <-chan T {
		return c
	})
}

func Batch[D, T any](done <-chan D, in <-chan T, maxItems int, maxTimeout time.Duration) <-chan []T {
	out := make(chan []T)

	go func() {
		defer close(out)

		timeout := time.After(maxTimeout)
		var batch []T
		for {
			select {
			case <-done:
				if len(batch) > 0 {
					out <- batch
					batch = nil
				}
				return
			case item, ok := <-in:
				if !ok {
					if len(batch) > 0 {
						out <- batch
						batch = nil
					}
					return
				} else {
					// Add to batch
					batch = append(batch, item)

					// Flush if limit is reached
					if len(batch) >= maxItems {
						out <- batch
						batch = nil
						timeout = time.After(maxTimeout)
					}
				}
			case <-timeout:
				if len(batch) > 0 {
					out <- batch
					batch = nil
				}
				timeout = time.After(maxTimeout)
			}
		}
	}()

	return out
}

func FormatJson[D, T any](done <-chan D, in <-chan T) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for item := range OrDone(done, in) {
			if bytes, err := json.Marshal(item); err != nil {
				panic(err)
			} else {
				out <- string(bytes)
			}
		}
	}()

	return out
}
