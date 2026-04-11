package onebot

import (
	"github.com/fahimahamedwork/openagent/pkg/bus"
	"github.com/fahimahamedwork/openagent/pkg/channels"
	"github.com/fahimahamedwork/openagent/pkg/config"
)

func init() {
	channels.RegisterFactory("onebot", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewOneBotChannel(cfg.Channels.OneBot, b)
	})
}
