package main

import (
  "fmt"
)

const prefixoOlaPortugues = "Olá, "

func Ola(nome string, idioma string) string {

  if nome == "" {
    nome = "Mundo"
  }

	if idioma == "espanhol" {
		return "Hola, "+ nome
	}

  return prefixoOlaPortugues + nome
}

func main() {
  fmt.Println(Ola("", ""))
}
