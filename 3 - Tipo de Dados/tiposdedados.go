package main

import (
	"errors"
	"fmt"
)

func main() {

	// NUMEROS INTEIROS

	var numero int64 = -100000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM NUMEROS INTEIROS

	// NUMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234555555445.45
	fmt.Println(numeroReal2)

	// FIM NUMEROS REAIS

	// STRINGS

	var str string = "ALALALALAA"
	fmt.Println(str)

	char := 'A'
	fmt.Println(char)

	// FIM STRINGS

	var valorVazio string
	fmt.Println(valorVazio)

	var valorVazio2 int
	fmt.Println(valorVazio2)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
