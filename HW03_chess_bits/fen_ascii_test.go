package HW03_chess_bits

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"testing"
)

func Test_fen2ascii(t *testing.T) {
	reader := utils.NewReader(t, "1743.1.FEN - ASCII")

	for reader.Next(t) {
		in := reader.In.RawString()
		out := reader.Out.RawString()

		res := fen2ascii(in)
		pass := res == out
		fmt.Printf("case %d: pass: %v \n want:\n%shave:\n%s", reader.Idx-1, pass, out, res)

		if !pass {
			t.Fail()
		}
	}
}
