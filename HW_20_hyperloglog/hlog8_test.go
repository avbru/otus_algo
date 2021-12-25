package hyperloglog

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHLog(t *testing.T) {
	require.Equal(t, 18.799227586206896, HLog8())
}
