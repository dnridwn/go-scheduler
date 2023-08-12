package goscheduler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// define global context
var gCtx context.Context = context.Background()

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

			ctx, cancel := context.WithTimeout(gCtx, tc.timeout)
			scheduler := NewScheduler()
			scheduler.Add(ctx, tc.cronExp, do)
			scheduler.Run()

			<-ctx.Done()
			time.Sleep(10 * time.Millisecond) // add time to wait to make sure
			cancel()

			assert.Equal(t, tc.expected, counter)
		})
	}
}
