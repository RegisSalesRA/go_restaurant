package handler

import (
	"net/http"
	"restaurante/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateTableInput struct {
	Number int  `json:"number"`
	Status string `json:"status"`
}

type TablesHandler struct {
	Repo *repository.TablesRepository
}

func (h *TablesHandler) SaveTable(c *gin.Context) {
	var input CreateTableInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O nome da Table é obrigatório"})
		return
	}

	table, err := h.Repo.SaveTables(c.Request.Context(), input.Number, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, table)
}

func (h *TablesHandler) TableList(c *gin.Context) {
	tablesHandlers, err := h.Repo.TablesList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tablesHandlers)
}

func (h *TablesHandler) GetTableByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	table, err := h.Repo.GetTableByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table não encontrada"})
		return
	}
	c.JSON(http.StatusOK, table)
}

func (h *TablesHandler) UpdateTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input CreateTableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err = h.Repo.UpdateTable(c.Request.Context(), id, input.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Table atualizada com sucesso"})
}

func (h *TablesHandler) DeleteTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.Repo.DeleteTable(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Table removida com sucesso"})
}

func (h *TablesHandler) SearchTablesHandler(c *gin.Context) {
	queryName := c.Query("number")

	if queryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro 'number' é obrigatório para a busca"})
		return
	}

	tables, err := h.Repo.FilterTablesByName(c.Request.Context(), queryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar a busca no banco"})
		return
	}

	c.JSON(http.StatusOK, tables)
}
