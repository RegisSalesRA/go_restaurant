package main


import "fmt"

/*
========================================
SINTAXE BÁSICA E VARIÁVEIS EM GO
========================================

Este arquivo demonstra:

1. Tipos de dados (int, string, bool, etc.)
2. Atribuição implícita (:=)
3. Saída no console com fmt
4. Operadores aritméticos

----------------------------------------
*/

// ========================================
// FUNÇÃO PRINCIPAL
// ========================================
func main2() {

	// ========================================
	// 1. TIPOS DE DADOS E VARIÁVEIS
	// ========================================

	// 🔹 Inteiros (com sinal)
	var idade1 int = 30        // padrão (32 ou 64 bits)
	var pequeno int8 = 127    // -128 até 127
	var medio int16 = 32000
	var grande int64 = 9000000000

	// 🔹 Inteiros sem sinal (não aceitam negativos)
	var positivo uint = 10
	var pequenoUint uint8 = 255   // também conhecido como byte
	var grandeUint uint64 = 100000

	// 🔹 Float (decimais)
	var altura float32 = 1.75
	var peso float64 = 80.50

	// 🔹 Booleano
	var ativo1 bool = true

	// 🔹 String
	var nome1 string = "Régis"

	// 🔹 Rune (caractere Unicode)
	var letraRune rune = 'A'   // representa um caractere (int32)

	// 🔹 Byte (alias de uint8)
	var b1 byte = 255


	fmt.Println("=== Tipos de Dados ===")
	fmt.Println("Nome:", nome1)
	fmt.Println("Idade:", idade1)
	fmt.Println("Ativo:", ativo1)
	fmt.Println("Altura:", altura)
	fmt.Println("Letra Rune:", letraRune)
	fmt.Println("Peso:", peso)
	fmt.Println("GrandeUint:", grandeUint)
	fmt.Println("PequenoUint:", pequenoUint)
	fmt.Println("Pequeno:", pequeno)
	fmt.Println("Medio:", medio)
	fmt.Println("Grande:", grande)
	fmt.Println("Positivo:", positivo)
	fmt.Println("Byte:", b1)


	// ========================================
	// 2. ATRIBUIÇÃO IMPLÍCITA (:=)
	// ========================================

	// Go infere automaticamente o tipo
	
	//var cidade = "Fortaleza"
	//var ano = 2026
	//var nome string = "Régis"
	// var idade int = 30

	cidade := "Fortaleza"
	ano := 2026

	// Go define valor padrão
	var nomeSemValor string   // ""
	var idadeSemValor int     // 0
	var ativoSemValor bool    // false

	// Várias variáveis na mesma linha
	//var nome, idade = "Régis", 30
	
	// Declaração em bloco

	/*
	var (
    nome  = "Régis"
    idade = 30
    ativo = true
	)
	*/	

	// Múltiplos com :=
	// nome, idade := "Régis", 30

	// constantes
	//  1. Constante simples
	const pi = 3.14
	const nome = "Golang"
	//  2. Constante com tipo
	const idade int = 30
	//  3. Constantes em bloco
	const (
	    statusAtivo  = true
	    statusInativo = false
	)
	// 4. iota (auto incremento)
	const (
	    Domingo = iota // 0
	    Segunda        // 1
	    Terca          // 2
	    Quarta         // 3
	)
	
	fmt.Println("\n=== Atribuição Implícita ===")
	fmt.Println("Cidade:", cidade)
	fmt.Println("Ano:", ano)
	fmt.Println("Ano:", nomeSemValor)
	fmt.Println("Ano:", idadeSemValor)
	fmt.Println("Ano:", ativoSemValor)
	fmt.Println("Ano:", ano)
	fmt.Println("Ano:", ano)


	// ========================================
	// 3. SAÍDA NO CONSOLE (fmt)
	// ========================================

	fmt.Println("\n=== Saída no Console ===")

	// Println → quebra linha automaticamente
	fmt.Println("Olá, Go!")

	// Print → não quebra linha
	fmt.Print("Texto sem quebra ")
	fmt.Print("continua aqui\n")

	// Printf → formatação
	fmt.Printf("Nome: %s | Idade: %d\n", nome, idade)

	// %s → string
	// %d → inteiro
	// %f → float

	// ========================================
	// 4. OPERADORES ARITMÉTICOS
	// ========================================

	fmt.Println("\n=== Operadores Aritméticos ===")

	a := 10
	c := 3 

	soma := a + c
	subtracao := a - c
	multiplicacao := a * c
	divisao := a / c        // divisão inteira
	resto := a % c          // resto da divisão

	fmt.Println("a =", a, "c =", c)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Resto:", resto)

	// Importante: divisão com float
	divisaoFloat := float64(a) / float64(c)
	fmt.Println("Divisão com float:", divisaoFloat)

}
 