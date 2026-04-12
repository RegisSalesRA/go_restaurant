package main

import (
	"fmt"
	"time"
)

/*
========================================
CONCORRÊNCIA EM GO
========================================

Este arquivo demonstra:

1. O que é concorrência
2. Uso de Goroutines (execução assíncrona)
3. Sincronização básica com sleep

----------------------------------------
*/

// ========================================
// FUNÇÃO SIMPLES
// ========================================

func tarefa(nome string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(nome, "executando:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

// ========================================
// MAIN
// ========================================

func main8() {

	fmt.Println("=== Execução normal (sequencial) ===")

	tarefa("Tarefa 1")
	tarefa("Tarefa 2")

	/*
		👉 Aqui executa:
		- Primeiro Tarefa 1 completa
		- Depois Tarefa 2 começa
	*/

	fmt.Println("\n=== Execução concorrente (Goroutines) ===")

	// Goroutines (execução paralela/assíncrona)
	go tarefa("Goroutine 1")
	go tarefa("Goroutine 2")

	/*
		👉 Aqui executa:
		- Ambas ao mesmo tempo
		- Intercalando execução
	*/

	// IMPORTANTE: esperar goroutines terminarem
	time.Sleep(3 * time.Second)

	fmt.Println("\nFinal do programa")
}

/*
========================================
RESUMO
========================================

CONCORRÊNCIA:
- Executar várias tarefas ao mesmo tempo

GOROUTINES:
- São funções executadas de forma assíncrona
- Extremamente leves (mais leves que threads)

SINTAXE:

    go minhaFuncao()

IMPORTANTE:
- O main NÃO espera goroutines automaticamente
- Use sleep ou sync.WaitGroup (mais avançado)

========================================
*/