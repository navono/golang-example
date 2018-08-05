package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 在一个函数内进行 取消 有三个方面
// 1. `goroutine` 的父 `goroutine` 想要取消
// 2. 一个 `goroutine` 可能想要取消它的子 `goroutine`
// 3. 在 `goroutine` 中的任意一个阻塞的操作想要获得抢占，这样它就可以被取消

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancle := context.WithTimeout(ctx, 1*time.Second)
	defer cancle()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}

func testContext() {
	var wg sync.WaitGroup
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting： %v\n", err)
			cancle()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell %v\n", err)
			cancle()
		}
	}()

	wg.Wait()
}
