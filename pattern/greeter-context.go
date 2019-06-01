package pattern

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

func genGreetingCtx(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch localeCtx, err := localeCtx(ctx); {
	case err != nil:
		return "", err
	case localeCtx == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported localeCtx")
}

func printGreetingCtx(ctx context.Context) error {
	greeting, err := genGreetingCtx(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genFarewellCtx(ctx context.Context) (string, error) {
	switch localeCtx, err := localeCtx(ctx); {
	case err != nil:
		return "", err
	case localeCtx == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported localeCtx")
}

func printFarewellCtx(ctx context.Context) error {
	farewell, err := genFarewellCtx(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func localeCtx(ctx context.Context) (string, error) {
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreetingCtx(ctx); err != nil {
			fmt.Printf("cannot print greeting： %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewellCtx(ctx); err != nil {
			fmt.Printf("cannot print farewell %v\n", err)
			cancel()
		}
	}()

	wg.Wait()
}
