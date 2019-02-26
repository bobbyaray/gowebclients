# gowebclients
This repository is for http client testing of go clients. It uses a simple web application that will calculate the fibonacci sequence of a given size. Benchmark testing is done with the built in go benchmark suite.

## mathserver
As mentioned above the math server is a very simple web server designed to serve requests efficiently. It supports the following calls:

```
fibonacci/{size} - Returns a fibonacci sequeance of the size given
squared/{num} - Returns the square root of a number
factorial{basenum} - Returns the factorial value of a given base num.
```
