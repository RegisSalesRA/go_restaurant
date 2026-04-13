package handler

import (
	"net/http"
	"restaurante/internal/repository" 
	"github.com/gin-gonic/gin"
)

// Inputs para validação do JSON que vem do Frontend
type OpenDrawerInput struct {
	InitialValue int `json:"initial_value" binding:"required,min=0"`
}

// Struct do Handler que "segura" o repositório
type CashDrawerHandler struct {
	Repo *repository.CashDrawerRepository
}

// AbrirCaixaHandler lida com o POST /caixa/abrir
func (h *CashDrawerHandler) AbrirCaixaHandler(c *gin.Context) {
	var input OpenDrawerInput

	// Valida o JSON de entrada
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor inicial inválido"})
		return
	}

	// Chama o Repository usando o contexto da requisição
	drawer, err := h.Repo.AbrirCaixa(c.Request.Context(), input.InitialValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, drawer)
}
 


/*

1. O que é essa Struct no Handler?
Go
type CashDrawerHandler struct {
    Repo *repository.CashDrawerRepository
}
Imagine que o CashDrawerHandler é uma caixa de ferramentas. Para que essa caixa funcione, ela precisa de uma ferramenta específica: o acesso ao banco de dados (Repo).

Em vez de deixar o banco de dados "solto" no código (como uma variável global), nós o guardamos dentro dessa Struct.

Isso se chama Injeção de Dependência. Significa: "Para eu existir, você precisa me entregar um Repository pronto".

2. O que é esse (h *CashDrawerHandler)?
Go
func (h *CashDrawerHandler) AbrirCaixaHandler(c *gin.Context) { ... }
Isso é um Method Receiver (Receptor de Método).

Diferença: Uma função comum é func Somar(a, b int). Você chama ela de qualquer lugar.

O seu caso: Ao colocar (h *CashDrawerHandler) antes do nome da função, você está dizendo: "Essa função pertence à struct CashDrawerHandler".

Por que fazemos isso?
Porque dentro da função, você ganha acesso ao h.Repo. É como se o h fosse o this do JavaScript ou Java, ou o self do Python. Ele permite que a função use o banco de dados que está guardado na struct.

3. O Fluxo da Requisição (O caminho do dado)
Imagine que um cliente faz um pedido para abrir o caixa. O caminho é este:

Gin (Router): Recebe a batida na porta (/caixa/abrir) e chama o AbrirCaixaHandler.

Handler (O Porteiro): * Ele olha para o JSON que chegou (ShouldBindJSON).

Se estiver errado, ele expulsa o usuário com um erro 400.

Se estiver certo, ele diz: "Ei, Repository, salva esse valor inicial de 50 reais lá no banco pra mim?" (h.Repo.AbrirCaixa).

Repository (O Operário): Ele não sabe nada de JSON ou HTTP. Ele só pega o número 50, monta o comando SQL e joga no PostgreSQL.

Banco de Dados: Grava e devolve o ID.

4. Por que usar c.Request.Context()?
Isso é uma das coisas mais geniais do Go.

Se o banco de dados travar e a query demorar 30 segundos, mas o usuário cansar e fechar o navegador antes, o Gin percebe que o "Contexto" foi cancelado.

O Go avisa o banco: "Pode parar de trabalhar, o usuário já foi embora!". Isso economiza muita memória e processamento.

Em resumo, o que você está fazendo é:
Structs: Para organizar onde as coisas ficam guardadas (como gavetas).

Ponteiros (*): Para não ficar copiando os dados na memória, mas sim apontar para o lugar onde o banco já está conectado.

Métodos: Para dar "poderes" às suas structs.

O segredo para não se perder:
Sempre que você vir h.Repo.AlgumaCoisa, lembre-se: "Estou usando a conexão com o banco que está guardada na minha gaveta h".

*/