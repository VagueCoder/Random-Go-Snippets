// Package broadcast is a service package that holds functions, methods and types \
// related to broadcasting a message to unlimited consumer goroutines.
package broadcast

import (
	"context"
)

// Channel is a custom type to hold channel and context variables.
type Channel struct {

	// Only unexported fields
	ch         chan interface{}
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// runProducer is creates a separate goroutine that continuously produces given message.
// Unexported. No return values.
func (c *Channel) runProducer(msg interface{}) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				close(c.ch)
				return
			default:
				c.ch <- msg
			}
		}
	}(c.ctx)
}

// close closes the goroutine. i,e., calls the cancellation function.
// Unexported. No return values.
func (c *Channel) close() {
	c.cancelFunc()
	_, _ = <-c.ch, <-c.ch
}

// NewBroadcaster initiates new broadcaster and produces first message continuously.
// Exported. Returns custom Channel type variable with channel, context and a cancellation function.
func NewBroadcaster(msg interface{}) Channel {
	ctx, cancel := context.WithCancel(context.Background())
	ch := Channel{
		ch:         make(chan interface{}, 1),
		ctx:        ctx,
		cancelFunc: cancel,
	}
	ch.runProducer(msg)
	return ch
}

// UpdateMessage kills the earlier producer channel and replaces with producer with new message.
// Exported. No return values.
func (c *Channel) UpdateMessage(msg interface{}) {
	c.close()
	*c = NewBroadcaster(msg)
}

// Read reads message from channel.
// Exported. Returns message.
func (c *Channel) Read() interface{} {
	return <-c.ch
}

// Close closes the producer goroutine.
// Exported. No return values.
func (c *Channel) Close() {
	c.close()
}

// IsClosed checks whether the channel is closed.
// Exported. Returns true if channel is closed, else false.
func (c *Channel) IsClosed() bool {
	return <-c.ch == nil
}
