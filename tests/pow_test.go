// pow_test.go
package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPow(t *testing.T) {
	t.Parallel()

	got := Pow(2, 3)
	want := 8

	assert.Equal(t, want, got)
	require.Equal(t, want, got)
}

func TestPow_Table(t *testing.T) {
	t.Parallel()

	cases := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{2, 3, 8},
		{1, 10, 1},
		{-2, 2, 4},
		{-2, 3, -8},
		{100, 0, 1},
	}

	for _, tc := range cases {
		got := Pow(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("Expected '%d', but got '%d'", tc.want, got)
		}
	}
}

func TestPow_Table_v2(t *testing.T) {
	t.Parallel()

	cases := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{2, 3, 8},
		{1, 10, 1},
		{-2, 2, -4},
		{-2, 3, -8},
		{100, 0, 1},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d**%d=%d", tc.arg1, tc.arg2, tc.want), func(t *testing.T) {
			got := Pow(tc.arg1, tc.arg2)
			if tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

// pow_test.go
func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Pow(2, 15)
	}
}

// FastPow returns x**y, the base-x exponential of y.
func FastPow(x, y int) int {
	if y == 0 {
		return 1
	} else if y < 0 {
		return 0
	}

	tmp := FastPow(x, y/2)
	if y%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

func BenchmarkPowCompare(b *testing.B) {
	// prepare data
	const (
		x = 2
		y = 64
	)
	b.ResetTimer()

	b.Run("Pow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Pow(x, y)
		}
	})
	b.Run("FastPow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FastPow(x, y)
		}
	})
	b.Run("math.Pow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = math.Pow(x, y)
		}
	})
}
