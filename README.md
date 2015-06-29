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
    === RUN TestCL
    --- PASS: TestCL (0.03s)
    PASS
    BenchmarkLD8k        300       5239412 ns/op
    BenchmarkLD32k       100      13962661 ns/op
    BenchmarkGL8k        200       8263652 ns/op
    BenchmarkGL32k       100      13360719 ns/op
    BenchmarkFR8k         10     119532226 ns/op
    BenchmarkFR32k        10     119533394 ns/op
    BenchmarkCL8k       1000       1528074 ns/op
    BenchmarkCL32k       300       5430894 ns/op
    ok      github.com/miku/roundup-langdetect  14.296s
