package handlers

import (
	"net/http"
	"restaurante/internal/repository" 
	"github.com/gin-gonic/gin"
)

// Inputs para validação do JSON que vem do Frontend
type OpenDrawerInput struct {
	InitialValue int `json:"initial_value" binding:"required,min=0"`
}

type CloseDrawerInput struct {
	FinalCounted int `json:"final_counted" binding:"required,min=0"`
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
 