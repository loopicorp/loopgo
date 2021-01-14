package events

import (
	"reflect"
)

type Response struct {
	EventName string
	Ret       []interface{}
}

type EventEmitter struct {
	events map[string][]reflect.Value
}

func NewEventEmitter() *EventEmitter {
	e := new(EventEmitter)
	e.Init()

	return e
}

func (e *EventEmitter) Init() {
	e.events = make(map[string][]reflect.Value)
}

// Listeners returns an array of listeners from the givent event
func (e *EventEmitter) Listeners(event string) []reflect.Value {
	return e.events[event]
}

// AddListener adds event listener on the givent event
func (e *EventEmitter) AddListener(event string, listener interface{}) {
	if _, exists := e.events[event]; !exists {
		e.events[event] = []reflect.Value{}
	}
	if listnr, ok := listener.(reflect.Value); ok {
		e.events[event] = append(e.events[event], listnr)
	} else {
		listnr := reflect.ValueOf(listener)
		e.events[event] = append(e.events[event], listnr)
	}
}

// On is alias to Add listener
func (e *EventEmitter) On(event string, listener interface{}) {
	e.AddListener(event, listener)
}

// RemoveListeners removes listeners for a givent event
func (e *EventEmitter) RemoveListeners(event string) {
	delete(e.events, event)
}

// ListenerCount returns the number of listeners for a given event.
func ListenerCount(emitter *EventEmitter, event string) int {
	if emitter == nil {
		panic("ListenerCount cannot be determined as the given emitter is nil")
	}
	if _, exists := emitter.events[event]; !exists {
		return 0
	}
	return len(emitter.events[event])
}

// Emit emits the given event. Puts all arguments following the event name
// into the Event's `args` member. Returns a channel if listeners were
// called, nil otherwise.
func (e *EventEmitter) Emit(event string, args ...interface{}) <-chan *Response {
	listeners, exits := e.events[event]
	if !exits {
		return nil
	}

	var callArgs []reflect.Value
	c := make(chan *Response)
	for _, arg := range args {
		callArgs = append(callArgs, reflect.ValueOf(arg))
	}
	for _, listener := range listeners {
		go func(listener reflect.Value) {
			retVals := listener.Call(callArgs)
			var ret []interface{}
			for _, r := range retVals {
				ret = append(ret, r.Interface())
			}
			c <- &Response{EventName: event, Ret: ret}
		}(listener)
	}

	return c
}
