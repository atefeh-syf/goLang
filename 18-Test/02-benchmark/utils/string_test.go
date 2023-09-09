package utils

import (
	"testing"
)

func BenchmarkConcatStringWithoutBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringWithoutBuilder("test", "test1")
	}
}

func BenchmarkConcatStringWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringWithBuilder("test", "test1")
	}
}

func BenchmarkStringWithoutBuilder(b *testing.B) {
	StringWithoutBuilder("test", "test1")
}

func BenchmarkStringWithBuilder(b *testing.B) {
	StringWithBuilder("test", "test1")
}
