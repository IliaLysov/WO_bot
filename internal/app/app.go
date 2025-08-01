package app

import (
	"context"
	"fmt"
	"tg_weather/config"
	"tg_weather/pkg/ow"
	"tg_weather/pkg/tg"
)

type Dependencies struct {
	BotAPI *tg.Bot
	OW     *ow.OW
}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependencies

	deps.OW = &ow.OW{c.OW}
	deps.BotAPI, err = tg.NewClient(c.Tg)
	if err != nil {
		return fmt.Errorf("tg.NewClient: %w", err)
	}
	Bot(deps)
	return nil
}
