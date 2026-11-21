package ratelimit

import (
	"fmt"
	"sync"
	"time"
	"math"
)

//Creating the token bucket type
type TokenBucket struct {
	MaxTokens int
	Rate int 
	Tokens int
	LastRefill time.Time
}


//Creating the initialization function
func NewTokenBucket(capacity int, rate int) *TokenBucket){
	return &TokenBucket{
		MaxTokens: capacity,
		Rate: rate,
		Tokens: capacity,
		LastRefill: time.Now(),
	}
}


//Method for refilling bucket
func (tb *TokenBucket) refill (){
	now := time.Now()
	time_elapsed := now.sub(tb.LastRefill)

	tokensToAdd := int(time_elapsed.seconds * float64(tb.Rate))
	tb.Tokens = math.Min(tb.Tokens+tokensToAdd, tb.MaxTokens)
	tb.LastRefill = now
}

func (tb *TokenBucket) Take (size int) bool {
	mutex = sync.Mutex
	mutex.lock
	defer mutex.unlock

	tb.refill()

	if (size <= tb.tokens){
		tb.Tokens -= size
		return true
	}
	return false
}



