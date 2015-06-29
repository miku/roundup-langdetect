README
======

Measure performance of various language detection libraries.

* https://github.com/taruti/langdetect
* https://github.com/endeveit/guesslanguage
* https://github.com/kapsteur/franco
* https://github.com/jbowles/cld2_nlpt

Run:

    $ make test
    go test -v .
    === RUN TestLD
    --- PASS: TestLD (0.02s)
    === RUN TestGL
    --- FAIL: TestGL (0.02s)
        language_detection_test.go:104: fixtures/pg12987.txt: got nb, want one of [da]
        language_detection_test.go:104: fixtures/pg12987.txt: got nb, want one of [da]
    === RUN TestFR
    --- PASS: TestFR (0.25s)
    === RUN TestCL
    --- PASS: TestCL (0.03s)
    FAIL
    exit status 1
    FAIL    github.com/miku/roundup-langdetect  0.378s
    make: *** [test] Error 1

    $ make bench
    go test -v -test.bench=. -test.run xxx
    PASS
    BenchmarkLD8k        300       5297970 ns/op
    BenchmarkLD32k       100      14129548 ns/op
    BenchmarkGL8k        200       8593321 ns/op
    BenchmarkGL32k       100      13808069 ns/op
    BenchmarkFR8k         10     126377321 ns/op
    BenchmarkFR32k        10     128064015 ns/op
    BenchmarkCL8k       1000       1526382 ns/op
    BenchmarkCL32k       300       5664950 ns/op
    ok      github.com/miku/roundup-langdetect  14.428s
