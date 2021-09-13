package HW02_happy_tickets

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"testing"
)

func TestTickets(t *testing.T) {
	reader := utils.NewReader(t, "1.Tickets")

	for reader.Next(t) {
		in := reader.In.Int()
		out := reader.Out.Int()

		res := Tickets(in)

		fmt.Printf("case %2d: input: %2d want: %19d, have: %19d\n", reader.Idx-1, in, out, res)
		if res != out {
			t.Fail()
		}
	}
}
