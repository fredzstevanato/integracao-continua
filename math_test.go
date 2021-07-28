package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Error("Resultado da soma é inválido: Resudado %d. Esperado,", total, 30)
	}
}
