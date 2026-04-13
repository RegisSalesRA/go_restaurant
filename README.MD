# Go intro fast

**start project**
- go mod init "project name"

**instalando dependências**

- go get -u golang.org/x/crypto/bcrypt   (Implementação da função de hashing de senhas adaptativa, projetada para mitigar ataques de força bruta e tabelas de arco-íris via salting.)

- go get -u github.com/gin-gonic/gin  (Framework HTTP de alto desempenho com roteamento baseado em Radix Tree e baixo overhead de memória.)
- go get -u github.com/jackc/pgx
- go get -u github.com/jackc/pgx/v5/pgxpool    (Gerenciador de pool de conexões, essencial para concorrência e reaproveitamento de sockets TCP no banco de dados.)
- go get -u github.com/golang-jwt/jwt/v5    (Biblioteca para geração e validação de JSON Web Tokens, permitindo a transmissão segura de claims entre partes de forma stateless.)
- go get -u github.com/joho/godotenv     (Carregador de variáveis de ambiente a partir de arquivos .env para conformidade com a metodologia Twelve-Factor App.)

- go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest  (Ferramenta de versionamento de esquema SQL, garantindo integridade e idempotência entre ambientes de desenvolvimento e produção.)
- go install github.com/air-verse/air@latest    (Utilitário de live-reloading para binários Go, automatizando o processo de compilação e execução durante o desenvolvimento.)

# Migrations

**manual**
1) migrate create -ext sql -dir migrations -seq "nome_tabela_migracao" ( sem aspas duplas, para voce documentar a tabela que ira fazer as mudanças ou criar)
2) migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/restaurante_go?sslmode=disable" up  (Subir no banco )

**script auto**
./scripts/migrate.sh create create_table_example
./scripts/migrate.sh up


**tests**
go test -coverprofile=c.out ./...