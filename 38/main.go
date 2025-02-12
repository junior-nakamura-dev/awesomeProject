package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	bookCtx(ctx)
}

func bookCtx(ctx context.Context) {
	token := ctx.Value("key").(string)
	fmt.Println(token)
}
