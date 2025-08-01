package main

import (
	"context"
	"log"
	"tg_weather/config"
	"tg_weather/internal/app"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(ctx, c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot stopped")
}
