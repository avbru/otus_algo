package HW09_b_trees

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandom(t *testing.T) {
	tr := NewRandom()
	//Test Length
	tr.Insert(1)
	require.Equal(t, 1, tr.Len())
	tr.Insert(2)
	require.Equal(t, 2, tr.Len())
	tr.Insert(2)
	require.Equal(t, 2, tr.Len())
}
