# XXH3 hash algorithm
A Go implementation of the 64/128 bit xxh3 algorithm, added the SIMD vector instruction set: AVX2 and SSE2 support to accelerate the hash processing.\
The original repository can be found here: https://github.com/Cyan4973/xxHash.


## Overview

For the input length larger than 240, the 64-bit version of xxh3 algorithm goes along with following steps to get the hash result.

### step1.  Initialize 8 accumulators used to store the middle result of each Iterator.
```
xacc[0] = prime32_3
xacc[1] = prime64_1
xacc[2] = prime64_2
xacc[3] = prime64_3
xacc[4] = prime64_4
xacc[5] = prime32_2
xacc[6] = prime64_5
xacc[7] = prime32_1
```

### step2.  Process 1024 bytes of input data as one block each time
```
while remaining_length > 1024{
    for i:=0, j:=0; i < 1024; i += 64, j+=8 {
        for n:=0; n<8; n++{
            inputN := input[i+8*n:i+8*n+8]
            secretN := inputN ^ secret[j+8*n:j+8*n+8]
            
            xacc[n^1] += inputN
            xacc[n]   +=  (secretN & 0xFFFFFFFF) * (secretN >> 32)
        }
    }
    
    xacc[n]   ^= xacc[n] >> 47
    xacc[n]   ^= secret[128+8*n:128+8*n:+8]
    xacc[n]   *= prime32_1
    
    remaining_length -= 1024
}
```

### step3.  Process remaining stripes (totally 1024 bytes at most)
```

for i:=0, j:=0; i < remaining_length; i += 64, j+=8 {
    for n:=0; n<8; n++{
        inputN := input[i+8*n:i+8*n+8]
        secretN := inputN ^ secret[j+8*n:j+8*n+8]
    
        xacc[n^1] += inputN
        xacc[n]   += (secretN & 0xFFFFFFFF) * (secretN >> 32)
    }

    remaining_length -= 64
}
```

### step4.  Process last stripe  (align to last 64 bytes)
```
for n:=0; n<8; n++{
    inputN := input[(length-64): (length-64)+8]
    secretN := inputN ^ secret[121+8*n, 121+8*n+8]

    xacc[n^1] += inputN
    xacc[n]   += (secretN & 0xFFFFFFFF) * (secretN >> 32)
}
```

### step5.  Merge & Avalanche accumulators
```
acc = length * prime64_1
acc += mix(xacc[0]^secret11, xacc[1]^secret19) + mix(xacc[2]^secret27, xacc[3]^secret35) +
    mix(xacc[4]^secret43, xacc[5]^secret51) + mix(xacc[6]^secret59, xacc[7]^secret67)

acc ^= acc >> 37
acc *= 0x165667919e3779f9
acc ^= acc >> 32
```

If the input data size is not larger than 240 bytes, the calculating steps are similar to the above description. The major difference lies in the data alignment. In the case of smaller input, the alignment size is 16 bytes. 

## Quickstart
The SIMD assembly file can be generated by the following command:
```
cd internal/avo && ./build.sh
```

Use Hash functions in your code:
```
package main

import "github.com/songzhibin97/gkit/sys/xxhash3"

func main() {
	println(xxhash3.HashString("hello world!"))
	println(xxhash3.Hash128String("hello world!"))
}
```
## Benchmark
go version: go1.15.10 linux/amd64\
CPU: Intel(R) Core(TM) i7-10700K CPU @ 3.80GHz\
OS: Linux bluehorse 5.8.0-48-generic #54~20.04.1-Ubuntu SMP\
MEMORY: 32G

```
go test -run=None -bench=. -benchtime=1000x -count=10 > 1000_10.txt && benchstat 1000_10.txt
```
```
name                               time/op
Default/Len0_16/Target64-4         88.6ns ± 0%
Default/Len0_16/Target128-4        176ns ± 0%
Default/Len17_128/Target64-4       1.07µs ± 2%
Default/Len17_128/Target128-4      1.76µs ± 1%
Default/Len129_240/Target64-4      1.89µs ± 2%
Default/Len129_240/Target128-4     2.82µs ± 3%
Default/Len241_1024/Target64-4     47.9µs ± 0%
Default/Len241_1024/Target128-4    52.8µs ± 1%
Default/Scalar/Target64-4          3.52ms ± 2%
Default/Scalar/Target128-4         3.52ms ± 1%
Default/AVX2/Target64-4            1.93ms ± 2%
Default/AVX2/Target128-4           1.91ms ± 1%
Default/SSE2/Target64-4            2.61ms ± 2%
Default/SSE2/Target128-4           2.63ms ± 4%
```