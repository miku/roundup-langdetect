README
======

Measure performance of various language detection libraries.

* https://github.com/taruti/langdetect
* https://github.com/endeveit/guesslanguage
* https://github.com/kapsteur/franco
* https://github.com/jbowles/cld2_nlpt

Run:

    $ make bench
    go test -v -bench=.
    === RUN TestLD
    --- PASS: TestLD (0.02s)
    === RUN TestGL
    --- PASS: TestGL (0.02s)
    === RUN TestFR
    --- PASS: TestFR (0.24s)
    PASS
    BenchmarkLD8k        300       5227946 ns/op
    BenchmarkLD32k       100      13773107 ns/op
    BenchmarkGL8k        200       8364864 ns/op
    BenchmarkGL32k       100      13961805 ns/op
    BenchmarkFR8k         10     122466794 ns/op
    BenchmarkFR32k        10     122241139 ns/op
    BenchmarkCL8k       1000       1488564 ns/op
    BenchmarkCL32k       300       5549939 ns/op
    ok      github.com/miku/roundup-langdetect  14.402s
