package hyperloglog

import (
	"math"
	"math/bits"
)

type HLog32 struct {
	buckets []int
	alpha   float64
	k       int //number of bits used for bucketID
	m       int //number of buckets
}

func New(k int) *HLog32 {
	m := 2 << (k - 1)
	return &HLog32{
		buckets: make([]int, m),
		alpha:   alpha(m),
		k:       k,
		m:       m,
	}
}

// Write - appends HASH uint32 value to buckets
func (h *HLog32) Write(hc uint32) {
	idx := int(hc >> (32 - h.k))
	zeroes := bits.LeadingZeros32(hc<<h.k) + 1
	if zeroes > h.buckets[idx] {
		h.buckets[idx] = zeroes
	}
}

// Count - return estimate of unique values count
func (h *HLog32) Count() int {
	sum := 0.0
	m := float64(h.m)

	for _, v := range h.buckets {
		sum += math.Pow(math.Pow(2, float64(v)), -1)
	}

	sum = h.alpha * m * m / sum
	return int(sum)
}

// Empty - return estimate of unique values count
func (h *HLog32) Empty() int {
	count := 0
	for _, v := range h.buckets {
		if v == 0 {
			count++
		}
	}
	return count
}

func alpha(m int) float64 {
	switch {
	case m == 4:
		return 0.5324
	case m == 8:
		return 0.6355
	case m == 16:
		return 0.673
	case m == 32:
		return 0.697
	case m == 64:
		return 0.709
	default:
		alpha := 0.7213 / (1 + 1.079/float64(m))
		return alpha
	}
}
