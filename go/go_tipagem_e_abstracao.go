package main

import "fmt"

/*
========================================
TIPAGEM E ABSTRAÇÃO EM GO
========================================

Este arquivo demonstra:

1. Funções
2. Structs (modelos de dados)
3. Interfaces (abstração e polimorfismo)
4. Generics (reutilização de código)

----------------------------------------
*/

// ========================================
// 1. FUNÇÕES
// ========================================

// Função simples (recebe 2 valores e retorna 1)
func somar(a int, b int) int {
	return a + b
}

// Função com múltiplos retornos (muito usado em Go)
func dividir(a int, b int) (int, int) {
	return a / b, a % b // quociente e resto
}

// Função com retorno nomeado
func saudacao(nome string) (mensagem string) {
	mensagem = "Olá, " + nome
	return
}

// ========================================
// 2. STRUCTS (modelos de dados)
// ========================================

// Struct representa uma entidade (ex: Produto)
type Produto struct {
	ID     int
	Nome   string
	Preco  float64
	Ativo  bool
}

// Método da struct (comportamento do Produto)
func (p Produto) Exibir() string {
	return fmt.Sprintf("Produto: %s | Preço: %.2f", p.Nome, p.Preco)
}

// Método com ponteiro (modifica o valor original)
func (p *Produto) AplicarDesconto(percentual float64) {
	p.Preco = p.Preco * (1 - percentual)
}

// ========================================
// 3. INTERFACES (abstração)
// ========================================

// Interface define um comportamento (contrato)
type Exibivel interface {
	Exibir() string
}

// Outra struct
type Usuario struct {
	Nome string
}

// Usuario implementa a interface automaticamente
func (u Usuario) Exibir() string {
	return "Usuário: " + u.Nome
}

// Função que aceita qualquer tipo que implemente Exibivel
func Mostrar(e Exibivel) {
	fmt.Println(e.Exibir())
}

// ========================================
// 4. GENERICS (Go 1.18+)
// ========================================

// Função genérica (aceita qualquer tipo)
func imprimir[T any](valor T) {
	fmt.Println("Valor:", valor)
}

// Função genérica com slice
func primeiro[T any](lista []T) T {
	return lista[0]
}

// ========================================
// MAIN
// ========================================

func main6() {

	// ========================================
	// FUNÇÕES
	// ========================================

	fmt.Println("=== Funções ===")

	fmt.Println("Soma:", somar(10, 5))

	div, resto := dividir(10, 3)
	fmt.Println("Divisão:", div, "| Resto:", resto)

	fmt.Println(saudacao("Régis"))

	// ========================================
	// STRUCTS
	// ========================================

	fmt.Println("\n=== Structs ===")

	produto := Produto{
		ID:    1,
		Nome:  "Coca-Cola",
		Preco: 5.50,
		Ativo: true,
	}

	fmt.Println(produto)
	fmt.Println(produto.Exibir())

	// Aplicando desconto (usa ponteiro)
	produto.AplicarDesconto(0.1)
	fmt.Println("Após desconto:", produto.Preco)

	// ========================================
	// INTERFACES
	// ========================================

	fmt.Println("\n=== Interfaces ===")

	usuario := Usuario{Nome: "Régis"}

	// Polimorfismo: mesma função, tipos diferentes
	Mostrar(produto)
	Mostrar(usuario)

	// ========================================
	// GENERICS
	// ========================================

	fmt.Println("\n=== Generics ===")

	imprimir("Texto")
	imprimir(123)
	imprimir(true)

	numeros := []int{10, 20, 30}
	fmt.Println("Primeiro número:", primeiro(numeros))

	nomes := []string{"Ana", "João"}
	fmt.Println("Primeiro nome:", primeiro(nomes))
}

/*
========================================
RESUMO FINAL
========================================

FUNÇÕES:
- Executam lógica
- Podem retornar múltiplos valores

STRUCTS:
- Representam dados (Produto, Usuario, etc)
- Base de APIs e banco de dados
- Podem ter métodos

INTERFACES:
- Definem comportamento
- Permitem desacoplamento e polimorfismo
- Implementação automática

GENERICS:
- Permitem reutilizar código
- Funcionam com múltiplos tipos

========================================
*/