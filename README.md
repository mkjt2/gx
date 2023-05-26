# gx
## Compare Golang compression libraries

How to decide? Try out a bunch of them at once. `gx` will perform a compress/decompress cycle for
a given file with a broad selection of algorithms (standard and third party).

### Usage

#### Install `gx`
```
$ git clone https://github.com/mkjt2/gx.git
$ go build
```

#### Run `gx`

Compress/decompress gx executable itself, sort results by compression ratio.
```
$ ./gx -ratio
Test file:  ./gx
Size (bytes):  2869058
Algo                      Package                             CX Ratio  CX Time (s)   DX Time (s)   
zstd-best-compression     github.com/klauspost/compress/zstd  1.79      0.10 (0.000)  0.01 (0.000)  
flate-best-compression    compress/flate                      1.73      0.21 (0.000)  0.02 (0.000)  
zlib-best-compression     compress/zlib                       1.73      0.19 (0.000)  0.02 (0.000)  
gzip-best-compression     compress/gzip                       1.73      0.19 (0.000)  0.02 (0.000)  
flate-default             compress/flate                      1.73      0.07 (0.000)  0.02 (0.000)  
zlib-default-compression  compress/zlib                       1.73      0.06 (0.000)  0.02 (0.000)  
gzip-default-compression  compress/gzip                       1.73      0.06 (0.000)  0.02 (0.000)  
zstd-better-compression   github.com/klauspost/compress/zstd  1.72      0.02 (0.000)  0.00 (0.000)  
zstd-default              github.com/klauspost/compress/zstd  1.71      0.01 (0.000)  0.00 (0.000)  
zstd-fastest              github.com/klauspost/compress/zstd  1.67      0.01 (0.000)  0.00 (0.000)  
flate-best-speed          compress/flate                      1.64      0.02 (0.000)  0.00 (0.000)  
zlib-best-speed           compress/zlib                       1.64      0.02 (0.000)  0.00 (0.000)  
gzip-best-speed           compress/gzip                       1.64      0.02 (0.000)  0.00 (0.000)  
flate-huffman-only        compress/flate                      1.31      0.01 (0.000)  0.00 (0.000)  
zlib-huffman-only         compress/zlib                       1.31      0.01 (0.000)  0.00 (0.000)  
gzip-huffman-only         compress/gzip                       1.31      0.01 (0.000)  0.00 (0.000)  
lzw                       compress/lzw                        1.27      0.03 (0.000)  0.03 (0.000)  
```

Compress/decompress file $F, sort results by decompression time.
```
$ ./gx -dx-time "$F"
Test file: "$F"
Size (bytes):  71044112
Algo                      Package                             CX Ratio  CX Time (s)   DX Time (s)   
zstd-fastest              github.com/klauspost/compress/zstd  3.29      0.16 (0.000)  0.01 (0.000)  
zstd-default              github.com/klauspost/compress/zstd  3.61      0.19 (0.000)  0.02 (0.000)  
zstd-better-compression   github.com/klauspost/compress/zstd  3.78      0.35 (0.000)  0.03 (0.000)  
gzip-best-speed           compress/gzip                       3.07      0.46 (0.000)  0.04 (0.000)  
flate-best-speed          compress/flate                      3.07      0.46 (0.000)  0.05 (0.000)  
zlib-best-speed           compress/zlib                       3.07      0.47 (0.000)  0.05 (0.000)  
flate-huffman-only        compress/flate                      1.54      0.24 (0.000)  0.06 (0.000)  
gzip-huffman-only         compress/gzip                       1.54      0.25 (0.000)  0.06 (0.000)  
zlib-huffman-only         compress/zlib                       1.54      0.26 (0.000)  0.07 (0.000)  
zstd-best-compression     github.com/klauspost/compress/zstd  4.07      2.72 (0.000)  0.17 (0.000)  
flate-default             compress/flate                      3.44      1.34 (0.000)  0.31 (0.000)  
gzip-best-compression     compress/gzip                       3.45      4.18 (0.000)  0.31 (0.000)  
flate-best-compression    compress/flate                      3.45      4.12 (0.000)  0.31 (0.000)  
zlib-best-compression     compress/zlib                       3.45      4.17 (0.000)  0.32 (0.000)  
zlib-default-compression  compress/zlib                       3.44      1.35 (0.000)  0.32 (0.000)  
gzip-default-compression  compress/gzip                       3.44      1.35 (0.000)  0.32 (0.000)  
lzw                       compress/lzw                        2.08      0.70 (0.000)  0.41 (0.000)  
```

Compress/decompress gx executable itself, sort results by compression time. Try 7 times, show averages,
```
$ ./gx -cx-time -n 7
Test file:  /Users/jackie/dev/gx/gx
Size (bytes):  2869058
Algo                      Package                             CX Ratio  CX Time (s)   DX Time (s)   
zstd-fastest              github.com/klauspost/compress/zstd  1.67      0.01 (0.000)  0.00 (0.000)  
flate-huffman-only        compress/flate                      1.31      0.01 (0.000)  0.00 (0.000)  
gzip-huffman-only         compress/gzip                       1.31      0.01 (0.001)  0.00 (0.000)  
zlib-huffman-only         compress/zlib                       1.31      0.01 (0.000)  0.00 (0.000)  
zstd-default              github.com/klauspost/compress/zstd  1.71      0.01 (0.000)  0.00 (0.000)  
zstd-better-compression   github.com/klauspost/compress/zstd  1.72      0.02 (0.001)  0.00 (0.001)  
flate-best-speed          compress/flate                      1.64      0.02 (0.000)  0.00 (0.000)  
zlib-best-speed           compress/zlib                       1.64      0.02 (0.000)  0.00 (0.000)  
gzip-best-speed           compress/gzip                       1.64      0.02 (0.005)  0.00 (0.000)  
lzw                       compress/lzw                        1.27      0.03 (0.001)  0.04 (0.000)  
flate-default             compress/flate                      1.73      0.06 (0.001)  0.02 (0.000)  
gzip-default-compression  compress/gzip                       1.73      0.07 (0.001)  0.02 (0.000)  
zlib-default-compression  compress/zlib                       1.73      0.07 (0.001)  0.02 (0.000)  
zstd-best-compression     github.com/klauspost/compress/zstd  1.79      0.11 (0.004)  0.01 (0.002)  
flate-best-compression    compress/flate                      1.73      0.19 (0.001)  0.02 (0.000)  
gzip-best-compression     compress/gzip                       1.73      0.19 (0.004)  0.02 (0.000)  
zlib-best-compression     compress/zlib                       1.73      0.20 (0.002)  0.02 (0.000)  

```
