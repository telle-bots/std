package runner

import (
	"context"
	"errors"
)

// waiterCmd represents wait command
type waiterCmd struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// Run waits until context is closed
func (w *waiterCmd) Run() error {
	<-w.ctx.Done()
	if err := w.ctx.Err(); !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop closes context
func (w *waiterCmd) Stop() {
	w.cancel()
}
