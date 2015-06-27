package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/endeveit/guesslanguage"
	"github.com/kapsteur/franco"
	"github.com/taruti/langdetect"
)

var tests = []struct {
	// input filename
	fn string
	// list of acceptable languages, order does not matter
	lang []string
}{
	{fn: "fixtures/30232-0.txt", lang: []string{"hu"}},
	{fn: "fixtures/ai-50-aHR0cDovL2R4LmRvaS5vcmcvMTAuMTUxNS9qZWVoLTE5OTYtMi0zMTg", lang: []string{"fr"}},
	{fn: "fixtures/ai-50-aHR0cDovL2R4LmRvaS5vcmcvMTAuMTQzMTUvZXZ0aC0xOTY0LTA3MDE", lang: []string{"de", "el"}},
	{fn: "fixtures/pg12987.txt", lang: []string{"da", ""}},
	{fn: "fixtures/pg17489.txt", lang: []string{"fr", ""}},
}

func NewReader(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewReader(file)
}

func BenchmarkLD8k(b *testing.B) {
	for _, t := range tests {
		reader := NewReader(t.fn)
		buf := new(bytes.Buffer)
		io.CopyN(buf, reader, 8*1024)
		for n := 0; n < b.N; n++ {
			langdetect.DetectLanguage(buf.Bytes(), "")
		}
	}
}

func BenchmarkLD32k(b *testing.B) {
	for _, t := range tests {
		reader := NewReader(t.fn)
		buf := new(bytes.Buffer)
		io.CopyN(buf, reader, 32*1024)
		for n := 0; n < b.N; n++ {
			langdetect.DetectLanguage(buf.Bytes(), "")
		}
	}
}

func BenchmarkGL8k(b *testing.B) {
	for _, t := range tests {
		// in fixtures/30232-0.txt: Input string contains invalid UTF-8-encoded runes
		if t.fn == "fixtures/30232-0.txt" {
			continue
		}
		reader := NewReader(t.fn)
		buf := new(bytes.Buffer)
		io.CopyN(buf, reader, 8*1024)
		for n := 0; n < b.N; n++ {
			_, err := guesslanguage.Guess(buf.String())
			if err != nil {
				b.Errorf("in %s: %s", t.fn, err)
			}
		}
	}
}

func BenchmarkGL32k(b *testing.B) {
	for _, t := range tests {
		reader := NewReader(t.fn)
		buf := new(bytes.Buffer)
		io.CopyN(buf, reader, 32*1024)
		for n := 0; n < b.N; n++ {
			_, err := guesslanguage.Guess(buf.String())
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkFR8k(b *testing.B) {
	for _, t := range tests {
		reader := NewReader(t.fn)
		buf := new(bytes.Buffer)
		io.CopyN(buf, reader, 8*1024)
		for n := 0; n < b.N; n++ {
			franco.DetectOne(buf.String())
		}
	}
}

func BenchmarkFR32k(b *testing.B) {
	for _, t := range tests {
		reader := NewReader(t.fn)
		buf := new(bytes.Buffer)
		io.CopyN(buf, reader, 32*1024)
		for n := 0; n < b.N; n++ {
			franco.DetectOne(buf.String())
		}
	}
}
