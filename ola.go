package main

import (
  "fmt"
)

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaPortugues = "Olá, "

func Ola(nome string, idioma string) string {

  if nome == "" {
    nome = "Mundo"
  }
	
	return saudacao(idioma) + nome
}

func saudacao(idioma string) (prefixo string) {
	
	switch idioma {
		case frances:
			prefixo = prefixoOlaFrances
		case espanhol:
			prefixo = prefixoOlaEspanhol
		default:
				prefixo = prefixoOlaPortugues
	}

  return

}


func main() {
  fmt.Println(Ola("", ""))
}
