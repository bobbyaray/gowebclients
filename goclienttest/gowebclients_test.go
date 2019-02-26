package goclienttest

import "testing"

func BenchmarkFastHttpClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendFastHttpCall("http://localhost:9080/fibonacci/20")
	}
}

func BenchmarkGoClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendCall("http://localhost:9080/fibonacci/20")
	}
}

func BenchmarkGoRequestCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendGoRequestCall("http://localhost:9080/fibonacci/20")
	}
}

func BenchmarkGRequestsCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendGRequestsCall("http://localhost:9080/fibonacci/20")
	}
}

func BenchmarkGorillaCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendGorillaCall("http://localhost:9080/fibonacci/20")
	}
}

func BenchmarkHeimdallCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendHeimdallRequestsCall("http://localhost:9080/fibonacci/20")
	}
}
