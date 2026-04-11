package pico

import (
	"github.com/fahimahamedwork/openagent/pkg/bus"
	"github.com/fahimahamedwork/openagent/pkg/channels"
	"github.com/fahimahamedwork/openagent/pkg/config"
)

func init() {
	channels.RegisterFactory("pico", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewPicoChannel(cfg.Channels.Pico, b)
	})
	channels.RegisterFactory("pico_client", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewPicoClientChannel(cfg.Channels.PicoClient, b)
	})
}
