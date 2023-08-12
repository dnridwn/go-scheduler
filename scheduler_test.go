package goscheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduler(t *testing.T) {
	testCases := []struct {
		name     string
		timeout  time.Duration
		cronExp  string
		expected uint8
	}{
		{
			name:     "Test Schedule Every 2 Seconds",
			timeout:  10 * time.Second, // 10 seconds
			cronExp:  "*/2 * * * * *",  // every 2 seconds
			expected: 5,
		},
		{
			name:     "Test Schedule Every Second",
			timeout:  10 * time.Second, // 10 seconds
			cronExp:  "* * * * * *",    // every second
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var counter uint8 = 0
			do := func() {
				counter++
			}

			scheduler := NewScheduler()
			scheduler.Add(tc.cronExp, do)
			scheduler.Run()

			time.Sleep(tc.timeout)
			time.Sleep(10 * time.Millisecond) // add time to wait for make sure

			assert.Equal(t, tc.expected, counter)
		})
	}
}
