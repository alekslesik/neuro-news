package grabber

import (
	"testing"
)

// Ваша функция translit здесь

func BenchmarkTranslit(b *testing.B) {
	// Пример тестовой строки для бенчмарка
	testString := "Пример строки для бенчмарка"

	// Сброс бенчмарка
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		translit(testString)
	}
}
