package containers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue(t *testing.T) {
	var q Queue

	q.Enqueue(1)
	q.Enqueue(2)
	require.Equal(t, []interface{}{1, 2}, q.List())
	require.Equal(t, 2, q.Len())

	el := q.Dequeue()
	require.Equal(t, 1, el)
	require.Equal(t, []interface{}{2}, q.List())
	require.Equal(t, 1, q.Len())

	el = q.Dequeue()
	require.Equal(t, 2, el)
	require.Equal(t, []interface{}{}, q.List())
	require.Equal(t, 0, q.Len())
}
