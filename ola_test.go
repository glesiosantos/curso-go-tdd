package main

import (
  "testing"
)

func TestOla(t *testing.T) {

	verificarMessagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

  t.Run("diz olá quando as pessoas informar os nome", func(t *testing.T){
    resultado := Ola("Glêsio","")
    esperado  := "Olá, Glêsio"
    
   verificarMessagemCorreta(t, resultado, esperado) 
  })

  t.Run("diz 'Olá Mundo' quando nome não for informado", func(t *testing.T){
    resultado := Ola("","")
    esperado  := "Olá, Mundo"
		verificarMessagemCorreta(t, resultado, esperado)
  })

	t.Run("diz em espanhol", func(t *testing.T){
		resultado := Ola("Glêsio", "espanhol")
		esperado  := "Hola, Glêsio"
		verificarMessagemCorreta(t, resultado, esperado)
	})

}
