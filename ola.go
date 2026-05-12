package main

import (
  "fmt"
)

const espanhol = "espanhol"
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaPortugues = "Olá, "

func Ola(nome string, idioma string) string {

  if nome == "" {
    nome = "Mundo"
  }

	if idioma == espanhol {
		return prefixoOlaEspanhol + nome
	}

	if idioma == "frances" {
		return "Bonjour, "+ nome
	}

  return prefixoOlaPortugues + nome
}

func main() {
  fmt.Println(Ola("", ""))
}
