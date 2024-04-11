package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Вводные:
// Функция executeTask может зависнуть.
// В ней не предусмотрен механизм отмены.
// Она не принимает Context или канал с событием отмены как аргумент.

func executeTask() {
	time.Sleep(10 * time.Second) //nolint:gomnd // it's learning code
}

func executeTaskWithTimeout(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	done := make(chan struct{})

	go func() {
		executeTask()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		err := fmt.Errorf("timed out, %w", ctx.Err())
		return err
	}
}

func main() {
	// Задача:
	// Для функции executeTask написать обертку executeTaskWithTimeout.
	// Функция executeTaskWithTimeout принимает аргументом тайм-аут,
	// через который функция executeTask будет отменена.
	// Если executeTask была отменена по тайм-ауту, нужно вернуть ошибку

	ctx := context.Background()
	timeout := 5 * time.Second //nolint:gomnd // it's learning code
	err := executeTaskWithTimeout(ctx, timeout)
	if err != nil {
		log.Println(err)
	}
}
