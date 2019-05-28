## Results

```
goos: darwin
goarch: amd64
pkg: github.com/nassor/buffer-on-pool
BenchmarkStdLibUnmarshal-8                                   	  500000	      2567 ns/op	     496 B/op	      10 allocs/op
BenchmarkStdLibDecoder-8                                     	  500000	      2823 ns/op	    1088 B/op	      12 allocs/op
BenchmarkJsonInterConfigStd-8                                	 2000000	       642 ns/op	     176 B/op	       8 allocs/op
BenchmarkJsonInterConfigStdDecoder-8                         	  500000	     17471 ns/op	    3138 B/op	      16 allocs/op
BenchmarkJsonInterConfigFastest-8                            	 2000000	       623 ns/op	     176 B/op	       8 allocs/op
BenchmarkJsonInterConfigFastestDecoder-8                     	 2000000	       755 ns/op	     872 B/op	      12 allocs/op
BenchmarkJsonInterConfigFastestBufferedDecoderWithWarmup-8   	 1000000	     31342 ns/op	    3136 B/op	      16 allocs/op
BenchmarkByteSliceNoPool-8                                   	  500000	      2757 ns/op	     496 B/op	      10 allocs/op
BenchmarkByteSliceWithPool-8                                 	  500000	      5635 ns/op	     594 B/op	      13 allocs/op
BenchmarkBufferNoPool-8                                      	  500000	      3037 ns/op	    1088 B/op	      12 allocs/op
BenchmarkBufferWithPool-8                                    	  300000	      6963 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBufferAndDataWithPool-8                             	  300000	     11068 ns/op	    1568 B/op	      15 allocs/op
```