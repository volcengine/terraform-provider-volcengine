package common

import (
	"context"

	"golang.org/x/sync/semaphore"
)

var syncSemaphore *semaphore.Weighted

func InitSyncLimit() {
	syncSemaphore = semaphore.NewWeighted(10)
}

func Acquire() {
	_ = syncSemaphore.Acquire(context.Background(), 1)
}

func Release() {
	syncSemaphore.Release(1)
}
