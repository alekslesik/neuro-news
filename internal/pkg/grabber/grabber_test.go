package grabber

import (
	"testing"
)

func BenchmarkTranslit(b *testing.B) {
	testString := "example test string"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		translit(testString)
	}
}
