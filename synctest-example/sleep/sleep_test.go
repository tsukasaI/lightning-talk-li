package sleep_test

import (
	"fmt"
	"synctest-example/sleep"
	"testing"
	"testing/synctest"
	"time"
)

// GOEXPERIMENT=synctest go test ./sleep -v -run TestSleepOneSecFlaky
func TestSleepOneSecFlaky(t *testing.T) {
	before := time.Now()
	sleep.Duration(1 * time.Second)
	after := time.Now()

	if d := after.Sub(before); d != 1*time.Second {
		t.Errorf("expected: 1s, took: %v", d)
	}
}

// GOEXPERIMENT=synctest go test ./sleep -v -run TestSleepOneSecProper
func TestSleepOneSecProper(t *testing.T) {
	before := time.Now()
	sleep.Duration(1 * time.Second)
	after := time.Now()

	if d := after.Sub(before); d < 1*time.Second {
		t.Errorf("expected: 1s, took: %v", d)
	}
}

// GOEXPERIMENT=synctest go test ./sleep -v -run TestSleepOneSecMultipleCases
func TestSleepOneSecMultipleCases(t *testing.T) {
	tests := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second, 4 * time.Second, 5 * time.Second}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			before := time.Now()
			sleep.Duration(tt)
			after := time.Now()

			if d := after.Sub(before); d < 1*time.Second {
				t.Errorf("expected: 1s, took: %v", d)
			}
		})
	}
}

// GOEXPERIMENT=synctest go test ./sleep -v -run TestSleepWithSynctest
func TestSleepWithSynctest(t *testing.T) {
	synctest.Run(func() {
		before := time.Now()
		sleep.Duration(1 * time.Second)
		after := time.Now()

		if d := after.Sub(before); d != 1*time.Second {
			t.Errorf("expected: 1s, took: %v", d)
		}
	})
}

// GOEXPERIMENT=synctest go test ./sleep -v -run TestSleepOneSecMultipleCasesWithSynctest
func TestSleepOneSecMultipleCasesWithSynctest(t *testing.T) {
	tests := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second, 4 * time.Second, 5 * time.Second}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			synctest.Run(func() {
				before := time.Now()
				sleep.Duration(tt)
				after := time.Now()

				if d := after.Sub(before); d < 1*time.Second {
					t.Errorf("expected: 1s, took: %v", d)
				}
			})
		})
	}
}
