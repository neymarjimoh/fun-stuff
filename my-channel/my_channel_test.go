package mychannel

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	ch1 := MakeChannel(WithCapacity(2)) // passing capacity
	ch2 := MakeChannel()                // no capacity, defaults to 1

	fmt.Printf("ch1 => %v, ch2 => %v\n", ch1.capacity, ch2.capacity)

	ch1.Send(78)
	value, _ := ch1.Receive()
	fmt.Printf("received ch1 value: %v\n", value)
	ch1.Close()
}
