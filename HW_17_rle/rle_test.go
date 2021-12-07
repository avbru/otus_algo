package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const datadir = "data"

func TestRLE(t *testing.T) {
	cases := []string{"", "a", "Я", "aaa fff ggg hhh", strings.Repeat("a", 300), "abbbbbaaaaafgfgfgfgfgfgfgfgfgfgfgfgfgfgfgfgfgfgfgfgf"}
	for _, c := range cases {
		in := []byte(c)
		require.Equal(t, in, decode(encode([]byte(c))))
		require.Equal(t, in, decodeP(encodeP([]byte(c))))
	}
}

func TestFiles(t *testing.T) {
	files, err := os.ReadDir(datadir)
	require.NoError(t, err)
	for _, f := range files {
		file, err := os.Open(path.Join(datadir, f.Name()))
		require.NoError(t, err)
		data, err := io.ReadAll(file)
		require.NoError(t, err)
		file.Close()

		res := encode(data)
		resP := encodeP(data)

		fmt.Printf("file: %10s rle/rlePlus ratio: %d%% / %d%%  \n", f.Name(), len(res)*100/len(data), len(resP)*100/len(data))

		require.Equal(t, data, decode(res))
		require.Equal(t, data, decodeP(resP))
	}
}
