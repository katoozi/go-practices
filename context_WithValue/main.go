package main

import (
	"context"
	"fmt"
)

type aKey string

func searchKey(ctx context.Context, key aKey) {
	v := ctx.Value(key)
	if v != nil {
		fmt.Println("Found value:", v)
		return
	}
	fmt.Println("key not found")
}

func main() {
	myKey := aKey("mySecretKey")
	ctx := context.WithValue(context.Background(), myKey, "mySecretValue")

	searchKey(ctx, myKey)

	searchKey(ctx, aKey("notThere"))
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}
