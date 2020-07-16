package alltask

import (
	"fmt"
	"sync"
	"testing"
)

func AllResponse() string {
	numOfRoutine := 10
	ch := make(chan string, numOfRoutine)
	for i := 0; i < numOfRoutine; i++ {
		go func(i int) {
			ch <- fmt.Sprintf("The result is from %d", i)
		}(i)
	}
	totalResult := ""
	for i := 0; i < numOfRoutine; i++ {
		// 阻塞式获取channel中的所有消息
		totalResult += <-ch + "\n"
	}
	// 关闭channel
	close(ch)
	return totalResult
}

func AllResponseWithRange() string {
	numOfRoutine := 10
	var wg sync.WaitGroup
	ch := make(chan string, numOfRoutine)
	for i := 0; i < numOfRoutine; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- fmt.Sprintf("The result is from %d", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// 关闭channel
	close(ch)
	totalResult := ""
	for v := range ch{
		totalResult += v + "\n"
	}
	return totalResult
}

func TestGetAllResponse(t *testing.T) {
	//t.Log(AllResponse())
	t.Log(AllResponseWithRange())
}
