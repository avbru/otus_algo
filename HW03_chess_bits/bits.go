package HW03_chess_bits

import (
	"github.com/avbru/algo/utils"
	"math/bits"
)

const (
	noA       = 18374403900871474942
	noAB      = 18229723555195321596
	noABC     = 17940362863843014904
	noABCD    = 17361641481138401520
	noABCDE   = 16204198715729174752
	noABCDEF  = 13889313184910721216
	noABCDEFG = 9259542123273814144

	noH       = 9187201950435737471
	noGH      = 4557430888798830399
	noFGH     = 2242545357980376863
	noEFGH    = 1085102592571150095
	noDEFGH   = 506381209866536711
	noCDEFGH  = 217020518514230019
	noBCDEFGH = 72340172838076673
)

func setBit(n uint64, pos int) uint64 {
	n |= 1 << pos
	return n
}

func king(n int) utils.ChessMoves {
	var K uint64
	K = setBit(K, n)

	Ka := K & noA
	Kh := K & noH
	M := (Ka << 7) | (K << 8) | (Kh << 9) |
		(Ka >> 1) | (Kh << 1) |
		(Ka >> 9) | (K >> 8) | (Kh >> 7)

	return utils.ChessMoves{
		N:    bits.OnesCount64(M),
		Mask: M,
	}
}

func knight(n int) utils.ChessMoves {
	var K uint64
	K = setBit(K, n)

	Ka := K & noA
	Kh := K & noH
	Kab := K & noAB
	Kgh := K & noGH
	M :=
		(Kgh << 17) |
			(Ka << 15) |
			(Kgh << 10) |
			(Kab << 6) |
			(Kgh >> 6) |
			(Kab >> 17) |
			(Kab >> 10) |
			(Kh >> 15)

	return utils.ChessMoves{
		N:    bits.OnesCount64(M),
		Mask: M,
	}
}

func rock(n int) utils.ChessMoves {
	var K uint64
	K = setBit(K, n)

	M :=
		(K << 8) |
			(K << 16) |
			(K << 24) |
			(K << 32) |
			(K << 40) |
			(K << 48) |
			(K << 56) |
			//
			(K >> 8) |
			(K >> 16) |
			(K >> 24) |
			(K >> 32) |
			(K >> 40) |
			(K >> 48) |
			(K >> 56) |
			//
			((K & noA) >> 1) |
			((K & noAB) >> 2) |
			((K & noABC) >> 3) |
			((K & noABCD) >> 4) |
			((K & noABCDE) >> 5) |
			((K & noABCDEF) >> 6) |
			((K & noABCDEFG) >> 7) |
			//
			((K & noH) << 1) |
			((K & noGH) << 2) |
			((K & noFGH) << 3) |
			((K & noEFGH) << 4) |
			((K & noDEFGH) << 5) |
			((K & noCDEFGH) << 6) |
			((K & noBCDEFGH) << 7)

	return utils.ChessMoves{
		N:    bits.OnesCount64(M),
		Mask: M,
	}
}

func bishop(n int) utils.ChessMoves {
	var K uint64
	K = setBit(K, n)

	M :=
	//вправо вниз
		((K & noH) >> 7) |
			((K & noGH) >> 14) |
			((K & noFGH) >> 21) |
			((K & noEFGH) >> 28) |
			((K & noDEFGH) >> 35) |
			((K & noCDEFGH) >> 42) |
			((K & noBCDEFGH) >> 49) |

			//влево вверх
			((K & noA) << 7) |
			((K & noAB) << 14) |
			((K & noABC) << 21) |
			((K & noABCD) << 28) |
			((K & noABCDE) << 35) |
			((K & noABCDEF) << 42) |
			((K & noABCDEFG) << 49) |
			//
			((K & noA) >> 9) |
			((K & noAB) >> 18) |
			((K & noABC) >> 27) |
			((K & noABCD) >> 36) |
			((K & noABCDE) >> 45) |
			((K & noABCDEF) >> 54) |
			((K & noABCDEFG) >> 63) |
			//
			((K & noH) << 9) |
			((K & noGH) << 18) |
			((K & noFGH) << 27) |
			((K & noEFGH) << 36) |
			((K & noDEFGH) << 45) |
			((K & noCDEFGH) << 54) |
			((K & noBCDEFGH) << 63)

	return utils.ChessMoves{
		N:    bits.OnesCount64(M),
		Mask: M,
	}
}
func queen(n int) utils.ChessMoves {
	var K uint64
	K = setBit(K, n)

	mRock := rock(n)
	mBishop := bishop(n)

	M := mRock.Mask | mBishop.Mask

	return utils.ChessMoves{
		N:    bits.OnesCount64(M),
		Mask: M,
	}
}
