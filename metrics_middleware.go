package main

import (
	"fmt"
	"time"

	"github.com/go-redis/store"
)

type MetricsMiddleware struct {
	inner          store.Storer
	getCalls       int
	getCallsFailed int
	setCalls       int
	delCalls       int
	lenCalls       int

	totalGetlatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricsMiddleware(store store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{
		inner: store,
	}
}
func (mm *MetricsMiddleware) Keys() []string {
	got := mm.inner.Keys()
	return got
}

func (mm *MetricsMiddleware) Get(key string) (string, error) {
	start := time.Now()
	val, err := mm.inner.Get(key)

	if err != nil {
		mm.getCallsFailed++
	}

	mm.getCalls++

	mm.totalGetlatency += time.Since(start)

	return val, err
}

func (mm *MetricsMiddleware) Set(key, value string) error {
	start := time.Now()

	err := mm.inner.Set(key, value)
	mm.setCalls++

	mm.totalSetLatency += time.Since(start)
	return err
}
func (mm *MetricsMiddleware) Del(key string) {
	mm.inner.Del(key)
	mm.delCalls++
}

func (mm *MetricsMiddleware) Len() int {
	mm.lenCalls++
	return mm.inner.Len()
}

func (mm *MetricsMiddleware) Report() {
	fmt.Printf("GET calls %d. (Failed %d)\n", mm.getCalls, mm.getCallsFailed)
	fmt.Printf("SET calls %d.\n", mm.setCalls)
	fmt.Printf("DEL calls %d.\n", mm.delCalls)
	fmt.Printf("LEN calls %d.\n", mm.lenCalls)

	if mm.getCalls > 0 {
		avg := mm.totalGetlatency / time.Duration(mm.getCalls)
		fmt.Printf("GET avg latency %v\n", avg)
	}

	if mm.setCalls > 0 {
		avg := mm.totalSetLatency / time.Duration(mm.setCalls)
		fmt.Printf("SET avg latency %v\n", avg)
	}
}
