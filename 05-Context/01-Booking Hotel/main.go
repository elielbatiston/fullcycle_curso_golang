package main

import (
	"context"
	"fmt"
	"time"
)

//WithValue: quando passamos o valor
//WithCancel: posso cancelar a qualquer momento independente de tempo
//WithDeadline: Você só pode cancelar depois do tempo ter estourado o limite
//WithTimeout: Contagem regressiva

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	//case <-time.After(5 * time.Second):
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
