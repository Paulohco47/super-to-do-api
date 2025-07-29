package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB inicializa a conexão com o banco de dados
func InitDB() {
	var err error

	// Obter parâmetros de conexão do banco de dados a partir de variáveis de ambiente
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "todoapp")

	// Criar string de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abrir conexão com o banco de dados
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		logCritical("Falha ao abrir conexão com o banco de dados: " + err.Error())
		os.Exit(1)
	}

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		logCritical("Falha ao testar conexão com o banco de dados: " + err.Error())
		os.Exit(1)
	}

	logSuccess("Conexão com o banco de dados estabelecida com sucesso")

	// Criar tabelas se não existirem
	createTables()
}

// createTables cria as tabelas necessárias no banco de dados
func createTables() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		logCritical("Falha ao criar tabelas: " + err.Error())
		os.Exit(1)
	}

	logSuccess("Tabelas do banco de dados criadas com sucesso")
}

// getEnv obtém uma variável de ambiente com um valor padrão
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
