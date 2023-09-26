package backoff_test

import (
	"github.com/ldcicconi/cmmn/backoff"
	"testing"
	"time"
)

func TestExponentialBackoffWaitTime(t *testing.T) {
	testCases := []struct {
		attempt  int
		baseTime time.Duration
		expected time.Duration
	}{
		{attempt: 0, baseTime: time.Second, expected: time.Second},
		{attempt: 1, baseTime: time.Second, expected: 2 * time.Second},
		{attempt: 2, baseTime: time.Second, expected: 4 * time.Second},
		{attempt: 3, baseTime: time.Second, expected: 8 * time.Second},
		{attempt: 4, baseTime: time.Second, expected: 16 * time.Second},
		{attempt: 5, baseTime: time.Second, expected: 32 * time.Second},
		{attempt: 6, baseTime: time.Second, expected: 64 * time.Second},
		{attempt: 0, baseTime: time.Second * 2, expected: time.Second * 2},
		{attempt: 1, baseTime: time.Second * 2, expected: 4 * time.Second},
		{attempt: 2, baseTime: time.Second * 2, expected: 8 * time.Second},
		{attempt: 3, baseTime: time.Second * 2, expected: 16 * time.Second},
		{attempt: 4, baseTime: time.Second * 2, expected: 32 * time.Second},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := backoff.ExponentialBackoffWaitTime(tc.attempt, tc.baseTime)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
			if int64(actual)/2 <= 0 {
				t.Errorf("Expected jitter randomness arg to be greater than zero, got %v", actual)
			}
		})
	}
}
