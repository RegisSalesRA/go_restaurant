package handler

import (
	"net/http"
	"restaurante/internal/auth"
	"restaurante/internal/models"
	"restaurante/internal/repository"
	"strconv"
	"os"
	"github.com/gin-gonic/gin"
)

type RegisterUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type UsersHandler struct {
	Repo *repository.UsersRepository
}

func NewUsersHandler(repo *repository.UsersRepository) *UsersHandler {
	return &UsersHandler{Repo: repo}
}

func (h *UsersHandler) RegisterUser(c *gin.Context) {
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: verifique e-mail e senha (mínimo 6 caracteres)"})
		return
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar senha"})
		return
	}

	user := models.Users{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
	}

	err = h.Repo.CreateUser(c.Request.Context(), user)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário registrado com sucesso"})
}

func (h *UsersHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := h.Repo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func (h *UsersHandler) UsersList(c *gin.Context) {
	users, err := h.Repo.UsersList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UsersHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "E-mail e senha são obrigatórios"})
		return
	}

	user, err := h.Repo.GetUserByEmail(c.Request.Context(), input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(input.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	token, err := auth.CreateJWT([]byte(os.Getenv("JWT_SECRET")), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
