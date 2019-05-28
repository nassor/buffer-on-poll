package main

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
	"testing"
)

var data = []byte(`
{
	"id": 100,
	"first_name": "Yvor",
	"last_name": "Hasnney",
	"email": "yhasnney2r@reuters.com",
	"gender": "Male",
	"ip_address": "26.183.247.4"
	}
`)

type Data struct {
	ID        json.Number `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Gender    string      `json:"gender"`
	IPAddress string      `json:"ip_address"`
}

func BenchmarkByteSliceNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := make([]byte, len(data), len(data))
		copy(payload, data)
		d := Data{}
		if err := json.Unmarshal(payload, &d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkByteSliceWithPool(b *testing.B) {
	pool := sync.Pool{New: func() interface{} { return []byte{} }}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payload := pool.Get().([]byte)
		defer func() { pool.Put(payload) }()
		payload = payload[:0] // reset
		payload = data
		d := Data{}
		if err := json.Unmarshal(payload, &d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkBufferNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(data)
		d := Data{}
		if err := json.NewDecoder(buf).Decode(&d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkBufferWithPool(b *testing.B) {
	pool := sync.Pool{New: func() interface{} { return bytes.NewBuffer([]byte{}) }}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := pool.Get().(*bytes.Buffer)
		defer func() { pool.Put(buf) }()
		buf.Reset()
		buf.Write(data)
		d := Data{}
		if err := json.NewDecoder(buf).Decode(&d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}

func BenchmarkBufferAndDataWithPool(b *testing.B) {
	pool := sync.Pool{New: func() interface{} { return bytes.NewBuffer([]byte{}) }}
	dpool := sync.Pool{New: func() interface{} { return &Data{} }}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := pool.Get().(*bytes.Buffer)
		defer func() { pool.Put(buf) }()
		buf.Reset()
		buf.Write(data)

		d := dpool.Get().(*Data)
		defer func() { dpool.Put(d) }()
		*d = Data{}
		if err := json.NewDecoder(buf).Decode(d); err != nil && err != io.EOF {
			b.Fatal(err)
		}
	}
}
