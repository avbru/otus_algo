package HW18_dt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_five8(t *testing.T) {
	require.Equal(t, 16, five8(5))
	require.Equal(t, 21892, five8(20))
}
