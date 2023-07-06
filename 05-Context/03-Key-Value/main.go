package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Hotel")
}

func bookHotel(ctx context.Context, name string) { //por convenção as variaveis de ctx vão em primeiro nos parametros da funcao
	token := ctx.Value("token")
	fmt.Println(token)
}
