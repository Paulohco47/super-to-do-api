# API de Lista de Tarefas (To-Do List API) em Go

Uma API RESTful completa para gerenciamento de tarefas, desenvolvida em Go com PostgreSQL como banco de dados e configuração Docker para facilitar a implantação.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Características](#características)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Pré-requisitos](#pré-requisitos)
- [Instalação e Execução](#instalação-e-execução)
- [Documentação da API](#documentação-da-api)
- [Exemplos de Uso](#exemplos-de-uso)
- [Estrutura do Banco de Dados](#estrutura-do-banco-de-dados)
- [Configuração](#configuração)
- [Testes](#testes)
- [Contribuição](#contribuição)
- [Licença](#licença)

## 🎯 Visão Geral

Esta API de Lista de Tarefas foi desenvolvida como um projeto educacional para demonstrar os conceitos fundamentais do desenvolvimento de APIs em Go, incluindo:

- Uso do `http.Server` nativo do Go
- Manipulação de JSON com `encoding/json`
- Tratamento de rotas básicas
- Integração com banco de dados PostgreSQL
- Containerização com Docker
- Operações CRUD completas

## ✨ Características

- **CRUD Completo**: Criar, ler, atualizar e deletar tarefas
- **API RESTful**: Endpoints bem definidos seguindo padrões REST
- **Banco de Dados**: Persistência com PostgreSQL
- **Containerização**: Configuração Docker e Docker Compose
- **Validação**: Validação de dados de entrada
- **Tratamento de Erros**: Respostas de erro padronizadas
- **Timestamps**: Controle automático de criação e atualização
- **Health Check**: Verificação de saúde do banco de dados

## 🛠 Tecnologias Utilizadas

- **Go 1.18+**: Linguagem de programação principal
- **PostgreSQL 13**: Banco de dados relacional
- **Docker & Docker Compose**: Containerização e orquestração
- **github.com/lib/pq**: Driver PostgreSQL para Go

## 📁 Estrutura do Projeto

```
super-to-do-api/
├── main.go              # Ponto de entrada da aplicação
├── models.go            # Estruturas de dados e modelos
├── handlers.go          # Handlers HTTP e lógica de negócio
├── database.go          # Configuração e conexão com o banco
├── init.sql             # Script de inicialização do banco
├── Dockerfile           # Configuração do container da API
├── docker-compose.yaml  # Orquestração dos serviços
├── .dockerignore        # Arquivos ignorados no build Docker
├── go.mod               # Dependências do módulo Go
├── go.sum               # Checksums das dependências
├── docs/                # Documentação Swagger/OpenAPI
│   ├── docs.go          # Código de geração da documentação
│   ├── swagger.json     # Especificação Swagger em JSON
│   └── swagger.yaml     # Especificação Swagger em YAML
└── README.md            # Documentação do projeto
```

## 📋 Pré-requisitos

- Docker 20.10+
- Docker Compose 1.29+

**OU** para execução local:

- Go 1.18+
- PostgreSQL 17+

## 🚀 Instalação e Execução

### Usando Docker Compose (Recomendado)

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd todo-api-go
```

2. Execute com Docker Compose:
```bash
docker-compose up --build
```

3. A API estará disponível em `http://localhost:8080`

### Execução Local

1. Instale as dependências:
```bash
go mod download
```

2. Configure as variáveis de ambiente:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=todoapp
```

3. Execute a aplicação:
```bash
go run .
```

## 📚 Documentação da API

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. Listar todas as tarefas
- **GET** `/tasks`
- **Descrição**: Retorna todas as tarefas cadastradas
- **Resposta de Sucesso**: 200 OK

#### 2. Obter uma tarefa específica
- **GET** `/tasks/{id}`
- **Descrição**: Retorna uma tarefa específica pelo ID
- **Parâmetros**: `id` (integer) - ID da tarefa
- **Resposta de Sucesso**: 200 OK
- **Resposta de Erro**: 404 Not Found

#### 3. Criar uma nova tarefa
- **POST** `/tasks`
- **Descrição**: Cria uma nova tarefa
- **Body**: JSON com `title` e `description`
- **Resposta de Sucesso**: 201 Created

#### 4. Atualizar uma tarefa
- **PUT** `/tasks/{id}`
- **Descrição**: Atualiza uma tarefa existente
- **Parâmetros**: `id` (integer) - ID da tarefa
- **Body**: JSON com campos opcionais (`title`, `description`, `completed`)
- **Resposta de Sucesso**: 200 OK
- **Resposta de Erro**: 404 Not Found

#### 5. Deletar uma tarefa
- **DELETE** `/tasks/{id}`
- **Descrição**: Remove uma tarefa
- **Parâmetros**: `id` (integer) - ID da tarefa
- **Resposta de Sucesso**: 200 OK
- **Resposta de Erro**: 404 Not Found

### Estrutura de Dados

#### Task (Tarefa)
```json
{
  "id": 1,
  "title": "Título da tarefa",
  "description": "Descrição detalhada da tarefa",
  "completed": false,
  "created_at": "2023-07-25T10:30:00Z",
  "updated_at": "2023-07-25T10:30:00Z"
}
```

#### Resposta Padrão da API
```json
{
  "success": true,
  "message": "Mensagem opcional",
  "data": {}
}
```

## 🔧 Exemplos de Uso

### 1. Criar uma nova tarefa
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Estudar Go",
    "description": "Aprender os conceitos básicos da linguagem Go"
  }'
```

### 2. Listar todas as tarefas
```bash
curl -X GET http://localhost:8080/tasks
```

### 3. Obter uma tarefa específica
```bash
curl -X GET http://localhost:8080/tasks/1
```

### 4. Atualizar uma tarefa
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true
  }'
```

### 5. Deletar uma tarefa
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## 🗄 Estrutura do Banco de Dados

### Tabela: tasks

| Campo | Tipo | Descrição |
|-------|------|-----------|
| id | SERIAL PRIMARY KEY | Identificador único da tarefa |
| title | VARCHAR(255) NOT NULL | Título da tarefa |
| description | TEXT | Descrição detalhada da tarefa |
| completed | BOOLEAN DEFAULT FALSE | Status de conclusão |
| created_at | TIMESTAMP DEFAULT CURRENT_TIMESTAMP | Data de criação |
| updated_at | TIMESTAMP DEFAULT CURRENT_TIMESTAMP | Data da última atualização |

### Índices
- `idx_tasks_completed`: Índice no campo `completed` para consultas por status
- `idx_tasks_created_at`: Índice no campo `created_at` para ordenação temporal

## ⚙️ Configuração

### Variáveis de Ambiente

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| DB_HOST | Host do banco PostgreSQL | localhost |
| DB_PORT | Porta do banco PostgreSQL | 5432 |
| DB_USER | Usuário do banco | postgres |
| DB_PASSWORD | Senha do banco | password |
| DB_NAME | Nome do banco de dados | todoapp |

### Docker Compose

O arquivo `docker-compose.yaml` configura:
- **postgres**: Container PostgreSQL com dados persistentes
- **todo-api**: Container da aplicação Go
- **Rede**: Rede bridge para comunicação entre containers
- **Volumes**: Persistência de dados do PostgreSQL
- **Health Check**: Verificação de saúde do banco antes de iniciar a API

## 🧪 Testes

Para testar a API, você pode usar:

1. **curl** (exemplos acima)
2. **Postman** ou **Insomnia**
3. **HTTPie**:
```bash
# Criar tarefa
http POST localhost:8080/tasks title="Nova tarefa" description="Descrição"

# Listar tarefas
http GET localhost:8080/tasks

# Atualizar tarefa
http PUT localhost:8080/tasks/1 completed:=true
```

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

**Desenvolvido com ❤️ em Go**

