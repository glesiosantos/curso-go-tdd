package main

import (
  "testing"
)

func TestOla(t *testing.T) {

  t.Run("diz olá quando as pessoas informar os nome", func(t *testing.T){
    resultado := Ola("Glêsio")
    esperado  := "Olá, Glêsio"
    
    if resultado != esperado {
      t.Errorf("resultado %s, esperado %s", resultado, esperado)
    }
  })

  t.Run("deverá diz 'Olá Mundo' quando nome não for informado", func(t *testing.T){
    resultado := Ola("")
    esperado  := "Olá, Mundo"

    if resultado != esperado {
      t.Errorf("resultado %s, esperado %s", resultado, esperado)
    }
  })

}
