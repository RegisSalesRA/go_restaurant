package main
 

import "fmt"

/*
========================================
ESTRUTURAS DE CONTROLE EM GO
========================================

Este arquivo demonstra:

1. Condicionais (if / else)
2. Switch (múltiplas decisões)
3. Looping com for (substitui while)
----------------------------------------
*/

func main3() {

	// ========================================
	// 1. CONDICIONAIS (if / else)
	// ========================================

	fmt.Println("=== Condicionais ===")

	idade := 18

	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}

	// if com condição extra
	if idade >= 18 && idade < 60 {
		fmt.Println("Adulto")
	}

	// if com inicialização (muito usado em Go)
	if valor := 10; valor > 5 {
		fmt.Println("Valor é maior que 5:", valor)
	}

	// ========================================
	// 2. SWITCH (múltiplas decisões)
	// ========================================

	fmt.Println("\n=== Switch ===")

	dia := 2

	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	default:
		fmt.Println("Outro dia")
	}

	// switch sem variável (tipo if encadeado)
	numero := 15

	switch {
	case numero < 10:
		fmt.Println("Menor que 10")
	case numero >= 10 && numero < 20:
		fmt.Println("Entre 10 e 20")
	default:
		fmt.Println("Maior ou igual a 20")
	}

	// múltiplos valores no mesmo case
	letra := "a"

	switch letra {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vogal")
	default:
		fmt.Println("Consoante")
	}

	// ========================================
	// 3. LOOPING COM FOR
	// ========================================

	fmt.Println("\n=== Loop For ===")

	// 🔹 For clássico (igual outras linguagens)
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// 🔹 For como while
	j := 0
	for j < 3 {
		fmt.Println("j =", j)
		j++
	}

	// 🔹 Loop infinito (cuidado!)
	k := 0
	for {
		fmt.Println("Loop infinito controlado:", k)
		k++
		if k == 3 {
			break // interrompe o loop
		}
	}

	// 🔹 Percorrendo slice (array dinâmico)
	lista := []string{"Go", "Python", "Java"}

	for index, valor := range lista {
		fmt.Println("Index:", index, "| Valor:", valor)
	}

	// 🔹 Ignorando index
	for _, valor := range lista {
		fmt.Println("Valor:", valor)
	}

}