package tax

import "testing"

// go test -bench=.
// go test -bench=. -run=^# => Roda somente os testes de benchmark.
//O ^# é uma expressão regular, como nao tem nenhum teste dessa forma, ele roda só o bench
// go test -bench=. -count=10 => Executa 10 vezes cada teste
// go test -bench=. -run=^# -count=10 => Executa 10 vezes cada teste rodando apenas os testes de bench
// go test -bench=. -run=^# -count=10 -benchtime=3s => roda 3s apenas em cada operacao
// go test -bench=. -run=^# -benchmem => Verifica a alocacao de memoria

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f but go %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but go %f", item.expected, result)
		}
	}
}

// verifica a performance
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
