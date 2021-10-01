package HW03_chess_bits

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"testing"
)

type Case struct {
	Filename string
	F        func(int) utils.ChessMoves
}

func Test_D(t *testing.T) {
	cases := []Case{
		{"0.BITS/1.Bitboard - Король", king},
		{"0.BITS/2.Bitboard - Конь", knight},
		{"0.BITS/3.Bitboard - Ладья", rock},
		{"0.BITS/4.Bitboard - Слон", bishop},
		{"0.BITS/5.Bitboard - Ферзь", queen},
	}

	for _, tCase := range cases {
		println(tCase.Filename)
		reader := utils.NewReader(t, tCase.Filename)
		for reader.Next(t) {

			in := reader.In.Int()
			out := reader.Out.ChessMoves()
			res := tCase.F(in)

			fmt.Printf("%v,%v\n", out, res)

			if out != res {
				t.Fail()
			}
		}

	}
}
