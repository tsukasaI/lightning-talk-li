package basic_test

import (
	"fmt"
	"synctest-example/basic"
	"testing"
)

func TestAdd(t *testing.T) {
	a, b, want := 1, 2, 3
	actual := basic.Add(a, b)
	if actual != want {
		t.Errorf("actual: %d, want: %d", actual, want)
	}
}

func TestAddFail(t *testing.T) {
	a, b, want := 1, 2, 4
	actual := basic.Add(a, b)
	if actual != want {
		t.Errorf("actual: %d, want: %d", actual, want)
	}
}

func TestAddWithSubTests(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -2, -3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			actual := basic.Add(tt.a, tt.b)
			if actual != tt.want {
				t.Errorf("actual: %d, want: %d", actual, tt.want)
			}
		})
	}
}
