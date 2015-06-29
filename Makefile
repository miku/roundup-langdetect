test:
	go test -v .

bench:
	go test -v -test.bench=. -test.run xxx
