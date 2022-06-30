package common

import "sync"

var locks map[string]*sync.Mutex
var methodLock *sync.Mutex

func InitLocks() {
	locks = make(map[string]*sync.Mutex)
	methodLock = new(sync.Mutex)
}

func ReleaseLock(key string) {
	locks[key].Unlock()
}

func TryLock(key string) {
	methodLock.Lock()
	if _, ok := locks[key]; !ok {
		locks[key] = new(sync.Mutex)
	}
	methodLock.Unlock()
	locks[key].Lock()
}
