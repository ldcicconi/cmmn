package backoff

import (
	"math"
	"math/rand"
	"time"
)

type Exponential struct {
	maxBackoffTime time.Duration
	baseWaitTime   time.Duration
	random         *rand.Rand
}

func NewExponentialBackoff(maxBackoffTime, baseWaitTime time.Duration) *Exponential {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	return &Exponential{
		maxBackoffTime: maxBackoffTime,
		baseWaitTime:   baseWaitTime,
		random:         random,
	}
}

func (b Exponential) WaitTime(attempt int) time.Duration {
	if attempt <= 0 {
		return b.baseWaitTime
	}

	if attempt >= 64 { // 2^64 will overflow int64
		return b.maxBackoffTime
	}

	expFactor := math.Pow(2, float64(attempt))
	// Check for overflow on multiplication
	if expFactor > math.MaxInt64/float64(b.maxBackoffTime) {
		return b.maxBackoffTime
	}

	waitTime := time.Duration(expFactor) * b.baseWaitTime
	// Add jitter to avoid synchronization effects
	interval := int64(waitTime) / 2
	if interval > 0 {
		jitter := b.random.Int63n(interval)
		waitTime += time.Duration(jitter)
	}

	if waitTime > b.maxBackoffTime {
		return b.maxBackoffTime
	}
	return waitTime
}
