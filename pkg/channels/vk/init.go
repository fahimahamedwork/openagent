package vk

import (
	"github.com/fahimahamedwork/openagent/pkg/bus"
	"github.com/fahimahamedwork/openagent/pkg/channels"
	"github.com/fahimahamedwork/openagent/pkg/config"
)

func init() {
	channels.RegisterFactory("vk", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewVKChannel(cfg, b)
	})
}
