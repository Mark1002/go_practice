package context

import (
	"context"
	"fmt"
)

func Execute() {
	ctx := context.Background()
	ctx = addValue(ctx)
	readValue(ctx)
}

type CustomKey string

var key CustomKey = "key"

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, key, "test-value")
}

func readValue(ctx context.Context) {
	val := ctx.Value(key)
	fmt.Println(val)
	newCtx := context.WithValue(ctx, key, "override-value")
	fmt.Println(newCtx.Value(key))
	fmt.Println(val)
}
