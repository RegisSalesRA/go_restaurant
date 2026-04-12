package main

import "fmt"

/*
========================================
COLEÇÕES DE DADOS EM GO
========================================

Este arquivo demonstra:

1. Arrays (tamanho fixo)
2. Slices (listas dinâmicas)
3. Maps (chave e valor)

----------------------------------------
*/

func main4() {

	// ========================================
	// 1. ARRAYS (tamanho fixo)
	// ========================================

	fmt.Println("=== Arrays ===")

	// Array com tamanho fixo
	var numeros [3]int = [3]int{10, 20, 30}

	fmt.Println("Array:", numeros)
	fmt.Println("Primeiro elemento:", numeros[0])

	// Alterando valor
	numeros[1] = 50
	fmt.Println("Array atualizado:", numeros)

	// Percorrendo array
	for i := 0; i < len(numeros); i++ {
		fmt.Println("Index:", i, "Valor:", numeros[i])
	}

	// ========================================
	// 2. SLICES (listas dinâmicas)
	// ========================================

	fmt.Println("\n=== Slices ===")

	// Criando slice
	lista := []int{1, 2, 3}

	fmt.Println("Slice:", lista)

	// Adicionando elementos
	lista = append(lista, 4, 5)
	fmt.Println("Após append:", lista)

	// Tamanho e capacidade
	fmt.Println("Tamanho:", len(lista))
	fmt.Println("Capacidade:", cap(lista))

	// Criando slice com make
	sliceMake := make([]int, 3) // tamanho 3
	fmt.Println("Slice com make:", sliceMake)

	// Alterando valores
	sliceMake[0] = 100
	sliceMake[1] = 200
	fmt.Println("Slice atualizado:", sliceMake)

	// Percorrendo slice (forma mais usada)
	for index, valor := range lista {
		fmt.Println("Index:", index, "Valor:", valor)
	}

	// ========================================
	// 3. MAPS (chave e valor)
	// ========================================

	fmt.Println("\n=== Maps ===")

	// Criando map
	pessoas := map[string]int{
		"Régis": 30,
		"Ana":   25,
	}

	fmt.Println("Map:", pessoas)

	// Acessando valor
	fmt.Println("Idade do Régis:", pessoas["Régis"])

	// Adicionando/alterando
	pessoas["João"] = 40
	fmt.Println("Map atualizado:", pessoas)

	// Verificando se chave existe
	idade, existe := pessoas["Maria"]

	if existe {
		fmt.Println("Maria existe:", idade)
	} else {
		fmt.Println("Maria não encontrada")
	}

	// Removendo item
	delete(pessoas, "Ana")
	fmt.Println("Após remover Ana:", pessoas)

	// Percorrendo map
	for chave, valor := range pessoas {
		fmt.Println("Nome:", chave, "Idade:", valor)
	}

}