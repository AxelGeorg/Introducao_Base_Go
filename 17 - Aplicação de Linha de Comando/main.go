package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

// Para executar
//
// go run main.go ip --host amazon.com.br
// go run main.go servidores --host amazon.com.br
//
// ou
// usar o arquivo gerado com o go build
//
// ./linha-de-comando ip --host amazon.com.br
// ./linha-de-comando servidores --host amazon.com.br
func main() {
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
