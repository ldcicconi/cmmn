package backoff

import (
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
	expFactor := 1 << uint(attempt)
	waitTime := time.Duration(expFactor) * b.baseWaitTime

	// Add jitter to avoid synchronization effects
	jitter := b.random.Int63n(int64(waitTime) / 2)
	waitTime += time.Duration(jitter)

	if waitTime > b.maxBackoffTime {
		return b.maxBackoffTime
	}
	return waitTime
}
