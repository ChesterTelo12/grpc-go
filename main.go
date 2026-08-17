package main

import (
	"fmt"
	"sync/atomic"
)

// MockCredentials implements PerRPCCredentials
type MockCredentials struct {
	counter int32
}

func (m *MockCredentials) GetRequestMetadata(ctx interface{}, uri ...string) (map[string]string, error) {
	val := atomic.AddInt32(&m.counter, 1)
	return map[string]string{"authorization": fmt.Sprintf("token-%d", val)}, nil
}

func main() {
	creds := &MockCredentials{}
	// Simulate retry loop
	for i := 0; i < 2; i++ {
		meta, _ := creds.GetRequestMetadata(nil)
		fmt.Printf("Attempt %d: %v\n", i+1, meta)
	}
}