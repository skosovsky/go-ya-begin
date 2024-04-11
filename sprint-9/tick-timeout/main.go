package main

import (
	"context"
	"fmt"
	"time"
)

func tick(ctx context.Context) {
	ticker := time.NewTicker(300 * time.Millisecond) //nolint:gomnd // it's learning code
	defer ticker.Stop()
	for i := 0; i < 20; i++ {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Print(i, " ") //nolint:forbidigo // it's learning code
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	tick(ctx)
}
