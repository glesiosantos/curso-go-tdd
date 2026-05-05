package main

import (
  "testing"
)

func OlaTest(t *testing.T) {
  resultado := Ola()
  esperado := "Olá Mundo"

  if resultado != esperado {
    t.Errorf("resultado %q, esperado %q", resultado, esperado)
  }
}
