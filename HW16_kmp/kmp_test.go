package HW16_kmp

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCreatePiSlow(t *testing.T) {
	s := "AABAABAAA"
	pi := []int{0, 0, 1, 0, 1, 2, 3, 4, 5, 2}
	require.Equal(t, pi, CreatePiSlow(s))
	require.Equal(t, pi, CreatePi(s))
}

func TestKMP(t *testing.T) {
	require.Equal(t, []int(nil), KMP("AABBBBAA", "C"))
	require.Equal(t, []int{0, 1, 2, 3, 4, 5}, KMP("AAAAAA", "A"))
	require.Equal(t, []int{0, 6}, KMP("AABBBBAA", "AA"))
	require.Equal(t, []int{0, 7}, KMP("AABABABAAB", "AAB"))
}

func BenchmarkCreatePi(b *testing.B) {
	s := strings.Repeat("AABAABAABB", 100)
	b.Run(fmt.Sprintf("PI SLOW"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CreatePiSlow(s)
		}
	})
	b.Run(fmt.Sprintf("PI FAST"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CreatePi(s)
		}
	})
}
