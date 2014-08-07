package channel_test

import (
	"testing"

	"github.com/mix3/go-ikusan/channel"
	"github.com/stretchr/testify/assert"
)

func TestChannel(t *testing.T) {
	assert.Equal(t, []string{}, channel.List())
	channel.Set("channel1", "keyword")
	get, ok := channel.Get("channel1")
	assert.Equal(t, true, ok)
	assert.Equal(t, "keyword", get.ChannelKeyword)
	assert.Equal(t, []string{"channel1"}, channel.List())
	channel.Set("channel2", "keyword")
	assert.Equal(t, []string{"channel1", "channel2"}, channel.List())
	channel.Del("channel1")
	assert.Equal(t, []string{"channel2"}, channel.List())
}
