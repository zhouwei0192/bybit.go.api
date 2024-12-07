package main

import (
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	_, _ = ws.Connect().SendSubscription([]string{"kline.1.BTCUSDT"})
	select {}
}
