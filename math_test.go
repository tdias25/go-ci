package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("resultado da soma eh invalido: %d,  esperado: %d", total, 30)
	}
}
