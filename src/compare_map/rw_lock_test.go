package compare_map

import (
	"fmt"
	"sync"
	"testing"
)

const (
	NUM_OF_READER int = 40
	READ_TIMES    int = 100000
)

var cache map[string]string

func init() {
	cache = map[string]string{}
	cache["a"] = "aa"
	cache["b"] = "bb"
}

func LockFree()  {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func ReadLock() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	m := new(sync.RWMutex)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				m.RLock()
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
				m.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkLockFree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LockFree()
	}
}

func BenchmarkLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadLock()
	}
}
