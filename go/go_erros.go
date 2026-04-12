package main

import "fmt"

/*
========================================
TRATAMENTO DE ERROS EM GO
========================================

Este arquivo demonstra:

1. Como funções retornam erros
2. O padrão obrigatório: if err != nil
3. Boas práticas no tratamento de erros

----------------------------------------
*/

// ========================================
// FUNÇÃO QUE PODE GERAR ERRO
// ========================================

func dividirSeguroErro(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("não é possível dividir por zero")
	}
	return a / b, nil
}

// ========================================
// OUTRO EXEMPLO (simulando banco/API)
// ========================================

func buscarUsuario(id int) (string, error) {

	if id == 0 {
		return "", fmt.Errorf("id inválido")
	}

	// Simulando sucesso
	return "Régis", nil
}

// ========================================
// MAIN
// ========================================

func main7() {

	// ========================================
	// EXEMPLO 1 - DIVISÃO
	// ========================================

	fmt.Println("=== Exemplo 1: Divisão ===")

	resultado, err := dividirSeguroErro(10, 2)

	// Sempre verificar erro
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", resultado)

	// Testando erro
	fmt.Println("\nTestando erro (divisão por zero):")

	_, err = dividirSeguroErro(10, 0)

	if err != nil {
		fmt.Println("Erro esperado:", err)
	}

	// ========================================
	// EXEMPLO 2 - BUSCAR USUÁRIO
	// ========================================

	fmt.Println("\n=== Exemplo 2: Buscar Usuário ===")

	usuario, err := buscarUsuario(1)

	if err != nil {
		fmt.Println("Erro ao buscar usuário:", err)
		return
	}

	fmt.Println("Usuário encontrado:", usuario)

	// Testando erro
	fmt.Println("\nTestando erro (ID inválido):")

	_, err = buscarUsuario(0)

	if err != nil {
		fmt.Println("Erro esperado:", err)
	}

}

/*
========================================
RESUMO
========================================

- Go NÃO usa try/catch
- Funções retornam (valor, error)
- Sempre verificar erro com:

    if err != nil {
        // tratar erro
    }

- nil significa "sem erro"

- Nunca ignore erros (evite usar "_")

========================================
*/