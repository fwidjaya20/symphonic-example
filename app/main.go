package main

import "fmt"

type Event interface {
	Signature() string
}

type Listener interface {
	Handle()
}

type Foo struct{}

func (e *Foo) Signature() string {
	return "foo"
}

type Bar struct{}

func (e *Bar) Signature() string {
	return "bar"
}

type FooListenerOne struct{}

func (l *FooListenerOne) Handle() {
	fmt.Println("Foo Listener One")
}

type FooListenerTwo struct{}

func (l *FooListenerTwo) Handle() {
	fmt.Println("Foo Listener Two")
}

type BarListenerOne struct{}

func (l *BarListenerOne) Handle() {
	fmt.Println("Bar Listener One")
}

func main() {
	handler := map[Event][]Listener{
		&Foo{}: {
			&FooListenerOne{},
			&FooListenerTwo{},
		},
		&Bar{}: {
			&BarListenerOne{},
		},
	}

	evt := &Bar{}

	fmt.Println("Find This", handler, evt)

	for _, it := range handler[evt] {
		it.Handle()
	}
}
