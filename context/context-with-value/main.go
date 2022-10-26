package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "Basic Xabaras")
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
