package routes

import (
	"restaurante/internal/handler"
	"restaurante/internal/middleware"
	"restaurante/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(r *gin.Engine, pool *pgxpool.Pool) {
    // 1. Instancia os Repositories (Camada de Dados)
    cashRepo := repository.NewCashDrawerRepository(pool)
    categoriesRepo := repository.NewCategoriesRepository(pool)
    productsRepo := repository.NewProductRepository(pool)
    usersRepo := repository.NewUsersRepository(pool)
    tablesRepo := repository.NewTablesRepository(pool)

    // 2. Instancia os Handlers (Camada de Controle)
    // Ajustei de 'handlers' para 'handler' para bater com seu import
    cashH := &handler.CashDrawerHandler{Repo: cashRepo}
    categoriesHandle := &handler.CategoriesHandler{Repo: categoriesRepo}
    usersH := &handler.UsersHandler{Repo: usersRepo}
    productsHandle := &handler.ProductsHandler{Repo: productsRepo}
    tablesHandle := &handler.TablesHandler{Repo: tablesRepo}
 

    // 3. Define as Rotas
    api := r.Group("/api/v1")
    {
        // --- Rotas de Usuários (Públicas) ---
        api.POST("/register", usersH.RegisterUser)
        api.POST("/login", usersH.Login)
        
        // Rotas de Caixa
        api.POST("/caixa/abrir", cashH.AbrirCaixaHandler)
        
        // Rotas de Categorias
        api.POST("/categories", categoriesHandle.SaveCategory)
        api.GET("/categories", categoriesHandle.CategoriesList)
        api.GET("/categories/search", categoriesHandle.SearchCategoriesHandler)
        api.GET("/categories/:id", categoriesHandle.GetCategoryByID)
        api.PUT("/categories/:id", categoriesHandle.UpdateCategory)
        api.DELETE("/categories/:id", categoriesHandle.DeleteCategory)

        // Rotas de Products
        api.POST("/products", productsHandle.SaveProduct)
        api.GET("/products", productsHandle.ProductsList)
        api.GET("/products/search", productsHandle.SearchProductsHandler)
        api.GET("/products/:id", productsHandle.GetProductByID)
        api.PUT("/products/:id", productsHandle.UpdateProduct)
        api.DELETE("/products/:id", productsHandle.DeleteProduct)

        // Rotas de Tables
        api.POST("/tables", tablesHandle.SaveTable)
        api.GET("/tables", tablesHandle.TableList)
        api.GET("/tables/search", tablesHandle.SearchTablesHandler)
        api.GET("/tables/:id", tablesHandle.GetTableByID)
        api.PUT("/tables/:id", tablesHandle.UpdateTable)
        api.DELETE("/tables/:id", tablesHandle.DeleteTable)

        // --- Rotas de Usuário (Protegidas ou Administrativas) ---
        // Se você quiser listar todos os usuários ou ver um perfil específico
        
        api.GET("/users", usersH.UsersList)

        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            // Usuários protegidos
            protected.GET("/users_auth", usersH.UsersList)
            protected.GET("/me", usersH.Me)
    }
  }
}