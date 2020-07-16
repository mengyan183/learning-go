package pool

import (
	"errors"
	"time"
)

// 空的struct
type EmptyCustomerStruct struct {
}

// 定义pool
type EmptyCustomerStructPool struct {
	bufferChan chan *EmptyCustomerStruct
}
// struct
type CustomerStruct struct {
	v int
}

// 定义pool
type CustomerStructPool struct {
	bufferChan chan *CustomerStruct
}

// 初始化对象池
func init() {
	bufferChan := make(chan EmptyCustomerStruct, 10)
	for i := 0; i < 10; i++ {
		bufferChan <- EmptyCustomerStruct{}
	}
	bufferStructChan := make(chan CustomerStruct, 10)
	for i := 0; i < 10; i++ {
		bufferStructChan <- CustomerStruct{1}
	}
}

// 获取池中的对象并设置等待超时时间
func (ePool *EmptyCustomerStructPool) Get(t time.Duration) (*EmptyCustomerStruct, error) {
	select {
	case s := <-ePool.bufferChan:
		return s, nil
	case <-time.After(t):
		return nil, errors.New("获取池中的对象超时")
	}
}

// 释放对象
func (ePool *EmptyCustomerStructPool) Release(e *EmptyCustomerStruct) error {
	select {
	case ePool.bufferChan <- e:
		return nil
	default:
		return errors.New("对象池已满")
	}
}


// 获取池中的对象并设置等待超时时间
func (ePool *CustomerStructPool) Get(t time.Duration) (*CustomerStruct, error) {
	select {
	case s := <-ePool.bufferChan:
		return s, nil
	case <-time.After(t):
		return nil, errors.New("获取池中的对象超时")
	}
}

// 释放对象
func (ePool *CustomerStructPool) Release(e *CustomerStruct) error {
	select {
	case ePool.bufferChan <- e:
		return nil
	default:
		return errors.New("对象池已满")
	}
}
