package mychannel

import (
	"errors"
	"sync"
)

type Channel struct {
	mu       sync.Mutex
	cond     *sync.Cond
	buffer   []interface{}
	length   int
	capacity int
	closed   bool
}

// SendOnlyChannel is a send-only version of the channel.
type SendOnlyChannel struct {
	ch *Channel
}

// ReceiveOnlyChannel is a receive-only version of the custom channel.
type ReceiveOnlyChannel struct {
	ch         *Channel
	isReadOnly bool
}

type Option func(ch *Channel)

func WithCapacity(capacity int) Option {
	return func(ch *Channel) {
		ch.capacity = capacity
	}
}

// Creates a new channel with optionally passing capacity
func MakeChannel(opts ...Option) *Channel {
	ch := &Channel{
		buffer:   make([]interface{}, 0),
		closed:   false,
		capacity: 1,
	}

	for _, opt := range opts {
		opt(ch)
	}

	ch.cond = sync.NewCond(&ch.mu)
	return ch
}

// Sends a value to the channel
func (c SendOnlyChannel) Send(data interface{}) error {
	return c.ch.Send(data)
}

// Close the channel
func (c SendOnlyChannel) Close() {
	c.ch.Close()
}

// Receives a value from the channel
func (c ReceiveOnlyChannel) Receive() (interface{}, error) {
	if !c.isReadOnly {
		return nil, errors.New("cannot receive from a read-only channel")
	}
	return c.ch.Receive()
}

// Sends data to the channel
func (ch *Channel) Send(data interface{}) error {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if ch.closed {
		return errors.New("channel closed")
	}

	for ch.length == ch.capacity {
		ch.cond.Wait()
	}

	ch.buffer = append(ch.buffer, data)
	ch.length++
	ch.cond.Signal()
	return nil
}

// Receives data from the channel
func (ch *Channel) Receive() (interface{}, error) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	for ch.length == 0 {
		if ch.closed {
			return nil, errors.New("channel closed")
		}
		ch.cond.Wait()
	}

	value := ch.buffer[0]
	ch.buffer = ch.buffer[1:]
	ch.length--
	ch.cond.Signal()
	return value, nil
}

// Close a channel
func (ch *Channel) Close() {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if !ch.closed {
		ch.closed = true
		ch.cond.Broadcast()
	}
}
