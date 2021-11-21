package HW05_data_structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue(t *testing.T) {
	values := [][2]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 1},
	}
	result := []int{3, 4, 1, 2}
	q := NewPQueue(2)

	for _, v := range values {
		q.Enqueue(v[0], v[1])
	}

	for i := 0; i < len(values); i++ {
		require.Equal(t, result[i], q.Dequeue(), i)
	}
}
