package broadcast_test

import (
	"testing"

	"github.com/VagueCoder/Random-Go-Snippets/Message-Broadcaster/broadcast"
	"github.com/stretchr/testify/assert"
)

var (
	ch      broadcast.Channel
	message interface{}
)

func TestNewBroadcaster(t *testing.T) {
	message = "Any"
	t.Run("InitCheck", func(t *testing.T) {
		ch = broadcast.NewBroadcaster(message)
		assert.NotNil(t, ch, "Broadcaster channel type shouldn't be nil.")
	})

	t.Run("Consume", func(t *testing.T) {
		assert.Equal(t, message, ch.Read(), "Broadcaster should produce message same as given.")
	})

	t.Run("UpdateMessage", func(t *testing.T) {
		message = struct {
			Name string
			Age  int
		}{"Vague", 25}
		ch.UpdateMessage(message)
		assert.Equal(t, message, ch.Read(), "Broadcaster should produce message same as given.")
	})

	t.Run("Close", func(t *testing.T) {
		assert.False(t, ch.IsClosed(), "Channel shouldn't be closed already.")
		ch.Close()
		assert.True(t, ch.IsClosed(), "Channel should be closed already.")
	})
}
