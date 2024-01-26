package utils

import (
	"math"
	"sync"
	"time"
)

// TokenBucket represents a token bucket implementation.
type TokenBucket struct {
	rate           int64
	maxTokens      int64
	currentTokens  int64
	lastRefillTime time.Time
	mutex          sync.Mutex
}

// NewTokenBucket creates a new TokenBucket instance with the given rate and maxTokens.
func NewTokenBucket(rate, maxTokens int64) *TokenBucket {
	return &TokenBucket{
		rate:           rate,
		maxTokens:      maxTokens,
		lastRefillTime: time.Now(),
		currentTokens:  maxTokens,
	}
}

// refill adds tokens to the bucket based on the elapsed time since the last refill.
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTime)
	tokensToBeAdded := (elapsed.Nanoseconds() * tb.rate) / 1000000000
	tb.currentTokens = int64(math.Min(float64(tb.currentTokens+tokensToBeAdded), float64(tb.maxTokens)))
	tb.lastRefillTime = now
}

// IsRequestAllowed checks if the given number of tokens can be consumed.
func (tb *TokenBucket) IsRequestAllowed(tokens int64) bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.refill()
	if tb.currentTokens >= tokens {
		tb.currentTokens -= tokens
		return true
	}
	return false
}
