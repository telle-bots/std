package runner

import (
	"context"
	"sync"
)

// runner represents command runner
type runner struct {
	sync.Mutex

	wg       sync.WaitGroup
	err      error
	running  bool
	commands []Command
}

// NewRunner provides new runner
func NewRunner(ctx context.Context, cancel context.CancelFunc) Runner {
	return &runner{
		commands: []Command{
			&waiterCmd{
				ctx:    ctx,
				cancel: cancel,
			},
		},
	}
}

// Add adds command to runner
func (r *runner) Add(cmd Command) {
	r.Lock()
	if r.running {
		r.wg.Add(1)
		go r.run(cmd)
	} else {
		r.commands = append(r.commands, cmd)
	}
	r.Unlock()
}

// Start calls all commands and waits for any to stop or context to close
func (r *runner) Start() error {
	r.Lock()
	if !r.running {
		r.running = true
		for _, cmd := range r.commands {
			r.wg.Add(1)
			go r.run(cmd)
		}
	}
	r.Unlock()

	r.wg.Wait()
	return r.err
}

// run calls command and stops all other commands on exit
func (r *runner) run(cmd Command) {
	err := cmd.Run()

	r.Lock()
	if r.err == nil {
		r.err = err
		for _, otherCmd := range r.commands {
			otherCmd.Stop()
		}
		r.commands = nil
	}
	r.Unlock()

	r.wg.Done()
}
