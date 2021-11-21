hw02:
	go test -v ./HW02_happy_tickets/tickets*
	go test -v ./HW02_happy_tickets/strlen*
hw03:
	go test -v ./HW03_chess_bits/fen*
	go test -v ./HW03_chess_bits/bits*
hw04:
	go test -v ./HW05_data_structures/...
	go test -v ./HW05_data_structures/... -bench=.
