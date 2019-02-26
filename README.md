# gowebclients
This repository is for http client testing of go clients. It uses a simple web application that will calculate the fibonacci sequence of a given size. Benchmark testing is done with the built in go benchmark suite.

## mathserver
As mentioned above the math server is a very simple web server designed to serve requests efficiently. It supports the following calls:

```
fibonacci/{size} - Returns a fibonacci sequeance of the size given
squared/{num} - Returns the square root of a number
factorial{basenum} - Returns the factorial value of a given base num.
```
## goclienttest
The go client test uses a variety of go web clients. Current clients used:

* Standard net/http
* Fasthttp - https://github.com/valyala/fasthttp
* Gorilla - https://github.com/gorilla/http (deprecated)
* GRequests - https://github.com/levigross/grequests
* GoRequests - https://github.com/parnurzeal/gorequest
* Heimdall - https://github.com/gojektech/heimdall

The clients are tested using a fibonacci/20 call on the math server. After getting the needed client libraries and building the project files you can execute the benchmark test with the following command in the goclienttest folder:

`go test -bench=.`

Here are some results after running the benchmark on an 2012 iMac (quad core 3.2 GHz i5) 24 GB RAM

```
BenchmarkFastHttpClient-4          10000            108941 ns/op
BenchmarkGoClient-4                10000            155577 ns/op
BenchmarkGoRequestCall-4            2000            798044 ns/op
BenchmarkGRequestsCall-4           10000            160296 ns/op
BenchmarkGorillaCall-4              2000            732623 ns/op
BenchmarkHeimdallCall-4             2000            841548 ns/op
```

According to these numbers FastHttp performs the best for a simple get command with the standard net/http library performing very well.

