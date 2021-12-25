package hyperloglog

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/spaolacci/murmur3"
	"github.com/stretchr/testify/require"
	"math/bits"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	hLog := New(2)
	for _, v := range users {
		hc := uint32(v) << 24
		hLog.Write(hc)
	}
	require.Equal(t, 18, hLog.Count())
}

func TestDataSet(t *testing.T) {
	domains := make(map[string]struct{})
	i := 0

	reader, err := os.Open("../.data/cdx-00000.gz")
	require.NoError(t, err)
	defer reader.Close()

	gr, err := gzip.NewReader(reader)
	require.NoError(t, err)
	scanner := bufio.NewScanner(gr)

	var hLogs []*HLog32
	for i := 2; i <= 10; i++ {
		hLogs = append(hLogs, New(i))
	}

	mur := murmur3.New32()
	for scanner.Scan() {
		i++
		domain := getDomain(scanner.Bytes())
		domains[domain] = struct{}{}

		_, _ = mur.Write([]byte(domain))
		for _, hLog := range hLogs {
			hLog.Write(mur.Sum32())
		}
		mur.Reset()
	}

	exact := float64(len(domains))
	fmt.Printf("Total: %d, exact: %.0f \n", i, exact)
	fmt.Println("N registers | Estimate")
	for _, hlog := range hLogs {
		est := float64(hlog.Count())
		fmt.Printf("%5d %.0f %.2f %% \n", hlog.m, est, (est-exact)/exact*100.0)
	}
}

func getDomain(line []byte) string {
	start := bytes.Index(line, []byte("://")) + 3
	end := 0

bytesloop:
	for next := start; next < len(line); next++ {
		switch line[next] {
		case '/', '"', ' ', ':', '?', '&', '=':
			end = next
			break bytesloop
		}
	}
	return string(line[start:end])
}

func Test_bits(t *testing.T) {
	k := 3
	hc := uint32(0b11100001111100010101110111010011)
	fmt.Printf("%b\n", int(hc>>(32-k)))
	fmt.Printf("%b\n", int(hc<<k))
	fmt.Printf("zeroes %d\n", bits.LeadingZeros32(hc<<k))

	h := New(k)
	h.Write(hc)

	fmt.Println(h.alpha, "k:", k, "m:", h.m, h.buckets)
}
