package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Server struct {
	EventEmitter
}

func TestEmbedding(t *testing.T) {
	s := new(Server)

	// Init call required when used as subtype
	s.EventEmitter.Init()
	s.On("recv", func(msg string) string {
		return msg
	})
	resp := <-s.Emit("recv", "Testing Embedded!")
	expected := "Testing Embedded!"

	assert.Equal(t, expected, resp.Ret[0].(string))
}

func TestEmitReturnsEventOnChat(t *testing.T) {
	emitter := NewEventEmitter()
	emitter.On("hello", func(name string) string {
		return "Hello! " + name
	})

	e := <-emitter.Emit("hello", "Kofi")

	assert.Equal(t, "hello", e.EventName)
}

func TestListernerCount(t *testing.T) {
	emitter := NewEventEmitter()
	assert.Equal(t, 0, ListenerCount(emitter, "hello"))
}
