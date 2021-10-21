package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f1(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("f1: %w", ctx.Err())
	case <-time.After(time.Millisecond):
		return fmt.Errorf("f1 err in 1ms")
	}
}
func f2(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("f2 : %w", ctx.Err())
	case <-time.After(time.Hour):
		return nil
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := f1(ctx); err != nil {
			fmt.Println(err)
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		if err := f2(ctx); err != nil {
			fmt.Println(err)
			cancel()
		}
	}()
	wg.Wait()
}
