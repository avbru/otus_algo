hw02:
	go test -v ./HW02_happy_tickets/tickets*
	go test -v ./HW02_happy_tickets/strlen*
hw03:
	go test -v ./HW03_chess_bits/fen*
	go test -v ./HW03_chess_bits/bits*
hw05:
	go test -v ./HW05_data_structures/...
	go test -v ./HW05_data_structures/... -bench=.
hw06:
	go test -v ./HW06_heap_sort/.
	go test -v ./HW06_heap_sort/... -bench=.
hw17:
	go test -v ./HW_17_rle/...
	go build -o rle.exe ./HW_17_rle/...
