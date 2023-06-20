package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(11, 15)

	esperado := 30

	if total != esperado {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, esperado)
	}
}