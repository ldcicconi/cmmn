package backoff_test

import (
	"github.com/ldcicconi/cmmn/backoff"
	"testing"
	"time"
)

// TODO: fix these to handle the jitter
func TestExponentialBackoffWaitTime(t *testing.T) {
	testCases := []struct {
		attempt  int
		baseTime time.Duration
		maxTime  time.Duration
		expected time.Duration
	}{
		{attempt: 0, baseTime: time.Second, maxTime: time.Second * 100, expected: time.Second},
		{attempt: 1, baseTime: time.Second, maxTime: time.Second * 100, expected: 2 * time.Second},
		{attempt: 2, baseTime: time.Second, maxTime: time.Second * 100, expected: 4 * time.Second},
		{attempt: 3, baseTime: time.Second, maxTime: time.Second * 100, expected: 8 * time.Second},
		{attempt: 4, baseTime: time.Second, maxTime: time.Second * 100, expected: 16 * time.Second},
		{attempt: 5, baseTime: time.Second, maxTime: time.Second * 100, expected: 32 * time.Second},
		{attempt: 6, baseTime: time.Second, maxTime: time.Second * 100, expected: 64 * time.Second},
		{attempt: 0, baseTime: time.Second * 2, maxTime: time.Second * 100, expected: time.Second * 2},
		{attempt: 1, baseTime: time.Second * 2, maxTime: time.Second * 100, expected: 4 * time.Second},
		{attempt: 2, baseTime: time.Second * 2, maxTime: time.Second * 100, expected: 8 * time.Second},
		{attempt: 3, baseTime: time.Second * 2, maxTime: time.Second * 100, expected: 16 * time.Second},
		{attempt: 4, baseTime: time.Second * 2, maxTime: time.Second * 100, expected: 32 * time.Second},
		{attempt: 31, baseTime: time.Second, maxTime: time.Second * 100, expected: time.Second * 100},
		//{attempt: 63, baseTime: time.Second, expected: 1 << 63 * time.Second},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			backoffer := backoff.NewExponentialBackoff(tc.maxTime, tc.baseTime)

			actual := backoffer.WaitTime(tc.attempt)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}

			if int64(actual)/2 <= 0 {
				t.Errorf("Expected jitter randomness arg to be greater than zero, got %v", actual)
			}
		})
	}

}
