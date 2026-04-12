package main

// Go é uma linguagem criada pelo Google com foco em:

/*

Performance (compilada → rápida como C)
Simplicidade (menos complexidade que Java/C++)
Concorrência nativa (goroutines)
Facilidade para backend e APIs

*/


// Muito usada para:

/*

APIs (como você quer fazer)
Microservices
CLI tools
Sistemas distribuídos

*/


import "fmt"

func main1() {
    fmt.Println("Hello, Go!")
}


// Estrutura:

// package main → ponto de entrada
// func main() → função principal
// fmt.Println → imprime no terminal


/*

Criar projeto
mkdir minha-api
cd minha-api
go mod init minha-api

# Isso cria o go.mod (gerenciador de dependências)

🔹 4. Rodar código
go run main.go

🔹 Estrutura simples
minha-api/
├── go.mod
└── main.go

*/



// Fundamentos de Binários

/*
Aqui está o ponto forte do Go 👇

Go é compilado → gera um executável (.bin)

🔹 Compilar projeto
go build

Isso gera um arquivo executável:

minha-api
🔹 Rodar o binário
./minha-api
🔹 Diferença importante
Linguagem	Execução
Python	Interpretada
Go	Compilada (binário)

Go não precisa de runtime depois de compilado

🔹 Exemplo prático

Código:

package main

import "fmt"

func main() {
    fmt.Println("API rodando...")
}

Compilar:
go build main.go

Executar:
./main
🔹 Build para outros sistemas

Go permite gerar binário para outros sistemas:

GOOS=windows GOARCH=amd64 go build

Isso gera .exe mesmo no Linux
*/