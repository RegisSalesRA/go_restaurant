#!/bin/bash

# Carrega as variáveis do arquivo .env
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo "❌ Erro: Arquivo .env não encontrado!"
    exit 1
fi

COMMAND=$1
NAME=$2

# Verifica se a DATABASE_URL foi carregada
# chmod +x scripts/migrate.sh

if [ -z "$DATABASE_URL" ]; then
    echo "Erro: DATABASE_URL não está definida no .env"
    exit 1
fi

case $COMMAND in
    "up")
        echo "Executando migrações UP..."
        migrate -path migrations -database "$DATABASE_URL" up
        ;;
    "down")
        COUNT=${NAME:-1}
        read -p "Reverter $COUNT migração(ões)? [y/N]: " confirm
        if [[ $confirm == [yY] ]]; then
            migrate -path migrations -database "$DATABASE_URL" down $COUNT
        fi
        ;;
    "create")
        if [ -z "$NAME" ]; then
            echo "Erro: Nome da migração é obrigatório!"
            exit 1
        fi
        migrate create -ext sql -dir migrations -seq "$NAME"
        ;;
    "force")
        if [ -z "$NAME" ]; then
            echo "Erro: Versão é obrigatória!"
            exit 1
        fi
        migrate -path migrations -database "$DATABASE_URL" force "$NAME"
        ;;
    "version")
        migrate -path migrations -database "$DATABASE_URL" version
        ;;
    *)
        echo "Uso: $0 {up|down|create 'nome'|force 'versão'|version}"
        exit 1
        ;;
esac