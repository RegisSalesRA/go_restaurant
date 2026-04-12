# Go intro fast

**start project**
- go mod init "project name"

**instalando dependências**  (  exemplo com o gin  )

- go get -u golang.org/x/crypto/bcrypt

- go get -u github.com/gin-gonic/gin
- go get -u github.com/jackc/pgx
- go get -u github.com/jackc/pgx/v5/pgxpool
- go get -u github.com/golang-jwt/jwt/v5
- go get -u github.com/joho/godotenv

- go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- go install github.com/air-verse/air@latest

# Migrations

**manual**
1) migrate create -ext sql -dir migrations -seq "nome_tabela_migracao" ( sem aspas duplas, para voce documentar a tabela que ira fazer as mudanças ou criar)
2) migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/restaurante_go?sslmode=disable" up  (Subir no banco )

**script auto**
./scripts/migrate.sh create create_table_example
./scripts/migrate.sh up