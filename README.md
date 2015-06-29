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
    --- FAIL: TestLD (0.04s)
        language_detection_test.go:64: fixtures/pg46456.txt: got it, want one of [la en]
    === RUN TestGL
    --- FAIL: TestGL (0.05s)
        language_detection_test.go:115: fixtures/pg12987.txt: got nb, want one of [da]
        language_detection_test.go:115: fixtures/pg12987.txt: got nb, want one of [da]
        language_detection_test.go:115: fixtures/pg46456.txt: got ca, want one of [la en]
        language_detection_test.go:115: fixtures/pg46456.txt: got ca, want one of [la en]
    === RUN TestFR
    --- FAIL: TestFR (0.42s)
        language_detection_test.go:168: fixtures/pg29640.txt: got , want one of [es]
        language_detection_test.go:168: fixtures/pg29640.txt: got , want one of [es]
        language_detection_test.go:168: fixtures/pg29409.txt: got , want one of [it en]
        language_detection_test.go:168: fixtures/pg29409.txt: got , want one of [it en]
        language_detection_test.go:168: fixtures/pg46456.txt: got , want one of [la en]
        language_detection_test.go:168: fixtures/pg46456.txt: got , want one of [la en]
    === RUN TestCL
    --- PASS: TestCL (0.04s)
    FAIL
    exit status 1
    FAIL    github.com/miku/roundup-langdetect  0.603s
    make: *** [test] Error 1

    $ make bench
    go test -v -test.bench=. -test.run xxx
    PASS
    BenchmarkLD8k        200       9842191 ns/op
    BenchmarkLD32k        50      27101816 ns/op
    BenchmarkGL8k        100      18258429 ns/op
    BenchmarkGL32k        50      24789546 ns/op
    BenchmarkFR8k         10     201354804 ns/op
    BenchmarkFR32k        10     199869835 ns/op
    BenchmarkCL8k        500       2880417 ns/op
    BenchmarkCL32k       100      10694101 ns/op
    ok      github.com/miku/roundup-langdetect  15.679s
