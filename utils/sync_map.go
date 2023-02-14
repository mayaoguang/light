package utils

import "sync"

type DataType interface {
	int | int64 | int32 | string
}
type GSyncMap[K comparable, V DataType] struct {
	Lock sync.RWMutex
	Data map[K]V
}

func (slf *GSyncMap[K, DataType]) Set(key K, value DataType) {
	slf.Lock.Lock()
	defer slf.Lock.Unlock()
	slf.Data[key] = value
}

func (slf *GSyncMap[K, DataType]) Get(key K) (r DataType) {
	slf.Lock.RLock()
	defer slf.Lock.RUnlock()
	d, ok := slf.Data[key]
	if ok {
		return d
	}
	return *new(DataType)
}
