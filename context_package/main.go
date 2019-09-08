package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	// set timeout for it
	// cancel after one second.
	// it's equal to line 20
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	// cancel after one second
	// time.AfterFunc(time.Second, cancel)

	// cancel when something in happend in stdin
	// go func() {
	// 	s := bufio.NewScanner(os.Stdin)
	// 	s.Scan()
	// 	cancel()
	// }()

	sleepAndTalk(ctx, 5*time.Second, "Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
