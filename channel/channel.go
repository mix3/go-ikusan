package channel

import (
	"sync"
	"time"
)

type channelMap struct {
	Map map[string]channel
	sync.RWMutex
}

type channel struct {
	ChannelKeyword string
	JoinAt         time.Time
}

var joinChannels = &channelMap{Map: make(map[string]channel)}

func Get(name string) (channel, bool) {
	joinChannels.Lock()
	defer joinChannels.Unlock()
	c, ok := joinChannels.Map[name]
	return c, ok
}

func Set(name, key string) {
	joinChannels.Lock()
	defer joinChannels.Unlock()
	joinChannels.Map[name] = channel{
		ChannelKeyword: key,
		JoinAt:         time.Now(),
	}
}

func Del(key string) {
	joinChannels.Lock()
	defer joinChannels.Unlock()
	delete(joinChannels.Map, key)
}

func List() []string {
	list := []string{}
	for k, _ := range joinChannels.Map {
		list = append(list, k)
	}
	return list
}
