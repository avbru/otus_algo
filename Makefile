hw02:
	go test -v ./HW02_happy_tickets/tickets*
	go test -v ./HW02_happy_tickets/strlen*
hw03:
	go test -v ./HW03_chess_bits/fen*
	go test -v ./HW03_chess_bits/bits*
hw04:
	go test -v ./HW04_fib_primes/...
hw05:
	go test -v ./HW05_data_structures/...
	go test -v ./HW05_data_structures/... -bench=.
hw06:
	go test -v ./HW06_heap_sort/.
	go test -v ./HW06_heap_sort/... -bench=.
hw16:
	go test -v ./HW16_kmp/.
	go test -v ./HW16_kmp/... -bench=.
hw17:
	go test -v ./HW_17_rle/...
	go build -o rle.exe ./HW_17_rle/...
hw20:
	mkdir -p .data
	wget -nc -O .data/cdx-00000.gz https://commoncrawl.s3.amazonaws.com/cc-index/collections/CC-MAIN-2021-49/indexes/cdx-00000.gz || true
	go test -v ./HW_20_hyperloglog/... -run DataSet