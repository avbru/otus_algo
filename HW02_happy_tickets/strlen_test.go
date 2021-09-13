package HW02_happy_tickets

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"testing"
)

func TestString(t *testing.T) {
	reader := utils.NewReader(t, "0.String")

	for reader.Next(t) {
		in := reader.In.String()
		out := reader.Out.Int()

		res := Len(in)
		fmt.Printf("case %d: want: %2d, have: %2d\n", reader.Idx-1, out, res)

		if res != out {
			t.Fail()
		}
	}
}
