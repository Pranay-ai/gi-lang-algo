package tokenBucket

import (
	"sync"
	"time"
)



type TokenBucket struct {
	capacity int
	tokens float64
	refillRate float64
	lastRefillTime time.Time
	mu  sync.Mutex
}

func NewTokenBucket( capacity int, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		tokens: float64(capacity),
		refillRate: refillRate,
		lastRefillTime: time.Now(),
	}
}


func (tb *TokenBucket) refill(){

	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub((tb.lastRefillTime)).Seconds()
	tb.tokens += elapsedTime * tb.refillRate
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}
	tb.lastRefillTime = now
}


func (tb *TokenBucket) AllowRequest() bool {
	tb.refill()
	tb.mu.Lock()
	defer tb.mu.Unlock()

	if tb.tokens < 1 {
		return false
	}
	tb.tokens--
	return true
}