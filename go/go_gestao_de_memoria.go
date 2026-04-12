package main

import "fmt"

/*
========================================
PONTEIROS EM GO (MEMÓRIA)
========================================

Este arquivo demonstra:

1. Diferença entre passagem por valor e referência
2. Uso de ponteiros (*)
3. Uso de endereço (&)

----------------------------------------
*/

// ========================================
// PASSAGEM POR VALOR (cópia)
// ========================================

// Aqui NÃO altera o valor original
func alterarValor(x int) {
	x = x + 10
}

// ========================================
// PASSAGEM POR REFERÊNCIA (ponteiro)
// ========================================

// Aqui altera o valor original
func alterarValorComPonteiro(x *int) {
	*x = *x + 10
}

// ========================================
// MAIN
// ========================================

func main() {

	fmt.Println("=== Ponteiros ===")

	valor := 10

	fmt.Println("Valor original:", valor)

	// 🔹 Passagem por valor (cópia)
	// cria uma cópia de 'valor'
	alterarValor(valor)
	fmt.Println("Após alterarValor (cópia):", valor)

	// 🔹 Passagem por referência (ponteiro)
	// passa o endereço de memória
	alterarValorComPonteiro(&valor)
	fmt.Println("Após alterarValorComPonteiro:", valor)

	// ========================================
	// EXEMPLO DIRETO DE PONTEIRO
	// ========================================

	fmt.Println("\n=== Referência de Memória ===")

	x := 20

	// p guarda o ENDEREÇO de x
	var p *int = &x

	fmt.Println("Valor de x:", x)

	// &x → endereço de memória
	fmt.Println("Endereço de x:", &x)

	// p → também é o endereço
	fmt.Println("Valor de p (endereço):", p)

	// *p → valor armazenado naquele endereço
	fmt.Println("Valor apontado por p:", *p)

	// Alterando valor via ponteiro
	*p = 50

	fmt.Println("Novo valor de x:", x)

}

/*
========================================
RESUMO
========================================

VALOR:
- Passa uma cópia
- Não altera o original

PONTEIRO:
- Passa o endereço de memória
- Altera o valor original

SÍMBOLOS:

& → pega o endereço
* → acessa o valor

========================================
*/