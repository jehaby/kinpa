package kinpa

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"testing"
)

var benchFile = "examples/ex3.txt"

var benchData io.ReadSeeker

func loadBenchData() {
	data, err := ioutil.ReadFile(benchFile)
	if err != nil {
		log.Panicf("couldn't open bench file: %v", err)
	}
	benchData = bytes.NewReader(data)
}

func BenchmarkClippings(b *testing.B) {
	loadBenchData()
	p := new(Parser)
	for i := 0; i < b.N; i++ {
		p.ParseClippings(benchData)
		benchData.Seek(0, io.SeekStart)
	}
}
