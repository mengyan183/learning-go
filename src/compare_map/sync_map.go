package compare_map

import "sync"

// 封装 sync.Map
type SyncMapAdapter struct {
	m sync.Map
}

func (ma *SyncMapAdapter) Set(key interface{}, val interface{}) {
	ma.m.Store(key, val)
}

func (ma *SyncMapAdapter) Get(key interface{}) (interface{}, bool) {
	return ma.m.Load(key)
}

func (ma *SyncMapAdapter) Del(key interface{}) {
	ma.m.Delete(key)
}


func CreateSyncMapAdapter() *SyncMapAdapter{
	return &SyncMapAdapter{}
}