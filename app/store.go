package main

import (
	"sync"
	"time"
)

type kvStore struct {
	mu          sync.RWMutex
	keyValStore map[string]kValue
}

type kValue struct {
	value  string
	expiry time.Time
}

func createStore() *kvStore {
	return &kvStore{
		keyValStore: make(map[string]kValue),
	}
}

func (kv *kvStore) Set(key string, value string, expiry string) (err error) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	var expire time.Time

	if expiry == "" {
		expire = time.Time{}
	} else {
		expiryDuration, _ := time.ParseDuration(expiry + "ms")
		expire = time.Now().Add(expiryDuration)
	}

	kv.keyValStore[key] = kValue{value: value, expiry: expire}

	return nil
}

func (kv *kvStore) Get(key string) (value string) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	kVal := kv.keyValStore[key]

	if kVal.expiry.IsZero() || kVal.expiry.After(time.Now()) {
		return kVal.value
	}

	delete(kv.keyValStore, key)

	return NullBulkString

}
