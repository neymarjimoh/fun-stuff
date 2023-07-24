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

func TestUnidirectionalChannel(t *testing.T) {
	ch := MakeChannel(WithCapacity(3))
	sendOnlyCh := SendOnlyChannel{ch: ch}
	receiveOnlyCh := ReceiveOnlyChannel{ch: ch, isReadOnly: true}

	for i := 1; i <= 5; i++ {
		sendOnlyCh.Send(i)
	}

	val, err := receiveOnlyCh.Receive()
	if err != nil {
		println("Error:", err)
	} else {
		println("Received:", val.(int))
	}

	sendOnlyCh.Close()
}
