package main

import "bytes"

func encode(in []byte) []byte {
	res := []byte{}
	reader := bytes.NewReader(in)

	count, prev := uint8(0), byte(0)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if prev == b && count < 255 {
			count++
			continue
		}

		res = append(res, count, prev)
		count, prev = 1, b
	}

	res = append(res, count, prev)
	return res
}

func decode(in []byte) (res []byte) {
	reader := bytes.NewReader(in)
	for {
		count, err := reader.ReadByte()
		b, err := reader.ReadByte()
		if err != nil {
			break
		}

		for i := uint8(0); i < count; i++ {
			res = append(res, b)
		}
	}

	if res == nil {
		return []byte{}
	}
	return
}
