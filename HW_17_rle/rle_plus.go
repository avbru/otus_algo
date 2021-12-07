package main

import (
	"bytes"
)

func encodeP(in []byte) []byte {
	var res []byte
	reader := bytes.NewReader(in)

	count, buf := int8(0), []byte{}

	for {
		b, err := reader.ReadByte()
		if err != nil {
			res = dropBuf(res, buf, count)
			return res
		}

		switch {
		case count == 0:
			count, buf = -1, []byte{b}
		case count == -127 || count == 127:
			res = dropBuf(res, buf, count)
			count, buf = -1, []byte{b}
		case count > 0 && b == buf[len(buf)-1]:
			count++
		case count > 0 && b != buf[len(buf)-1]:
			res = dropBuf(res, buf, count)
			count = -1
			buf = []byte{b}
		case count < 0 && b == buf[len(buf)-1]:
			res = dropBuf(res, buf[:-count-1], count+1)
			buf = []byte{b}
			count = 2
		case count < 0 && b != buf[len(buf)-1]:
			buf = append(buf, b)
			count--
		}
	}
}

func dropBuf(res []byte, buf []byte, count int8) []byte {
	if count == 0 {
		return res
	}

	res = append(res, byte(count))

	if count > 0 {
		res = append(res, buf[0])
	} else {
		res = append(res, buf...)
	}

	return res
}

func decodeP(in []byte) []byte {
	res := []byte{}
	reader := bytes.NewReader(in)
	for {
		count, err := reader.ReadByte()
		if err != nil {
			break
		}

		ct := int8(count)

		switch {
		case ct > 0:
			b, _ := reader.ReadByte()
			res = append(res, bytes.Repeat([]byte{b}, int(ct))...)
		case ct < 0:
			for i := int8(0); i < -ct; i++ {
				b, _ := reader.ReadByte()
				res = append(res, b)
			}
		}
	}

	return res
}
