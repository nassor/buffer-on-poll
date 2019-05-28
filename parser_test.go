package main

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkStdLibUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := make([]byte, len(data), len(data))
		copy(payload, data)
		d := Data{}
		if err := json.Unmarshal(payload, &d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkStdLibDecoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := bytes.NewBuffer(data)
		d := Data{}
		if err := json.NewDecoder(payload).Decode(&d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonInterConfigStd(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := Data{}
		if err := json.Unmarshal(data, &d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonInterConfigStdDecoder(b *testing.B) {
	var responsePool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payload := responsePool.Get().(*bytes.Buffer)
		defer responsePool.Put(payload)
		payload.Reset()
		payload.ReadFrom(bytes.NewReader(data))

		d := Data{}
		if err := json.NewDecoder(payload).Decode(&d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonInterConfigFastest(b *testing.B) {
	var json = jsoniter.ConfigFastest
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := Data{}
		if err := json.Unmarshal(data, &d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonInterConfigFastestDecoder(b *testing.B) {
	var json = jsoniter.ConfigFastest
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payload := bytes.NewBuffer(data)
		d := Data{}
		if err := json.NewDecoder(payload).Decode(&d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonInterConfigFastestBufferedDecoderWithWarmup(b *testing.B) {
	var responsePool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	// warmup
	for i := 0; i < 1000; i++ {
		payload := responsePool.Get().(*bytes.Buffer)
		payload.ReadFrom(bytes.NewReader(data))
		responsePool.Put(payload)
	}
	var json = jsoniter.ConfigFastest
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payload := responsePool.Get().(*bytes.Buffer)
		defer responsePool.Put(payload)
		payload.Reset()
		payload.ReadFrom(bytes.NewReader(data))

		d := Data{}
		if err := json.NewDecoder(payload).Decode(&d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}
