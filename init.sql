-- Criação do banco de dados todoapp
CREATE DATABASE IF NOT EXISTS todoapp;

-- Conectar ao banco todoapp
\c todoapp;

-- Criação da tabela tasks
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inserção de dados de exemplo
INSERT INTO tasks (title, description, completed) VALUES
('Configurar ambiente de desenvolvimento', 'Instalar Go, PostgreSQL e Docker', true),
('Implementar API CRUD', 'Criar endpoints para Create, Read, Update e Delete', false),
('Escrever documentação', 'Documentar todos os endpoints da API', false),
('Configurar Docker', 'Criar Dockerfile e docker-compose.yaml', false);

-- Criar índices para melhor performance
CREATE INDEX IF NOT EXISTS idx_tasks_completed ON tasks(completed);
CREATE INDEX IF NOT EXISTS idx_tasks_created_at ON tasks(created_at);
CREATE UNIQUE INDEX IF NOT EXISTS idx_tasks_title_unique ON tasks(title);

