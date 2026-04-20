package middleware

import (
	"net/http"
	"restaurante/internal/auth"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtém o header de autorização
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header é necessário"})
			c.Abort()
			return
		}

		// 2. Valida o formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido (use 'Bearer TOKEN')"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Valida o JWT (usando a função ValidateJWT que exportamos no pacote auth)
		token, err := auth.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		// 4. Extrai as Claims e injeta o userID no contexto do Gin
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falha ao processar as claims do token"})
			c.Abort()
			return
		}

		// Recupera o userID (no seu CreateJWT ele foi salvo como String)
		userID, ok := claims["userID"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário não encontrado no token"})
			c.Abort()
			return
		}

		// Define o userID no contexto para que os Handlers (como o /me) possam usá-lo
		c.Set("userID", userID)
		c.Next()
	}
}