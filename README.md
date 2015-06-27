README
======

Measure performance of various language detection libraries.

* https://github.com/taruti/langdetect
* https://github.com/endeveit/guesslanguage
* https://github.com/kapsteur/franco

	$ make bench
	go test -bench=.
	PASS
	BenchmarkLD8k	     300	   5385115 ns/op
	BenchmarkLD32k	     100	  14242785 ns/op
	BenchmarkGL8k	     200	   8677636 ns/op
	BenchmarkGL32k	     100	  14081606 ns/op
	BenchmarkFR8k	      10	 124369006 ns/op
	BenchmarkFR32k	      10	 124560068 ns/op
	ok  	github.com/miku/roundup-langdetect	10.488s
