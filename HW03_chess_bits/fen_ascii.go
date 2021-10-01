package HW03_chess_bits

import (
	"fmt"
	"strconv"
	"strings"
)

const line = "  +-----------------+\r\n"

func fen2ascii(fen string) string {
	buf := strings.Builder{}
	buf.WriteString(line)
	buf.WriteString("8 |")

	reader := strings.NewReader(fen)
	n := 8
loop:
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		switch char {
		case '/':
			n--
			buf.WriteString(fmt.Sprintf(" |\r\n%d |", n))
		case 'r', 'n', 'b', 'q', 'k', 'p',
			'R', 'N', 'B', 'Q', 'K', 'P':
			buf.WriteString(" " + string(char))

		case '1', '2', '3', '4', '5', '6', '7', '8':
			count, _ := strconv.Atoi(string(char))
			buf.WriteString(strings.Repeat(" .", count))
		case ' ':
			break loop
		}
	}
	buf.WriteString(" |\r\n")
	buf.WriteString(line)
	buf.WriteString("    a b c d e f g h  \r\n\r\n")
	return buf.String()
}
