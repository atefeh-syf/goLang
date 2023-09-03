package examples

import (
	"context"
	"time"
)

func InternalCancellationExample() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//Timeout
	//ctx, cancel = context.WithTimeout(ctx, time.Second * 3)

	//Deadline
	//ctx, cancel = context.WithDeadline(ctx, time.Date(2022, time.February, 13, 23, 30, 0, 0, time.UTC))

	defer cancel()
	go func() {
		for {
			if time.Now().Second()%15 == 0 {
				cancel()
			}
		}
	}()

	go process1(ctx, 0)
	process2(ctx, 0)
}

func process1(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			println("Task canceled", ctx.Err().Error())
			return
		default:
			num++
			println("Processing ", num)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func process2(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			println("Task canceled", ctx.Err().Error())
			return
		default:
			num--
			println("Processing ", num)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
