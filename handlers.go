package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// getTasks godoc
// @Summary      Lista todas as tarefas
// @Description  Retorna todas as tarefas cadastradas
// @Tags         tasks
// @Produce      json
// @Success      200 {object} APIResponse{data=[]Task}
// @Failure      500 {object} APIResponse
// @Router       /tasks [get]
func getTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logWarning("Método não permitido em /tasks")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, title, description, completed, created_at, updated_at FROM tasks ORDER BY id ASC")
	if err != nil {
		logError("Falha ao buscar tarefas: " + err.Error())
		respondWithError(w, http.StatusInternalServerError, "Falha ao buscar tarefas")
		return
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt); err != nil {
			logError("Falha ao ler tarefa: " + err.Error())
			respondWithError(w, http.StatusInternalServerError, "Falha ao ler tarefa")
			return
		}
		tasks = append(tasks, task)
	}

	if tasks == nil {
		tasks = make([]Task, 0)
	}
	logSuccess("Tarefas listadas com sucesso")
	respondWithJSON(w, http.StatusOK, APIResponse{Success: true, Data: tasks})
}

// taskHandler manipula operações CRUD para uma única tarefa
func taskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID da tarefa inválido")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTask(w, r, id)
	case http.MethodPut:
		updateTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// getTask godoc
// @Summary      Busca uma tarefa por ID
// @Description  Retorna uma tarefa específica pelo ID
// @Tags         tasks
// @Produce      json
// @Param        id path int true "ID da tarefa"
// @Success      200 {object} APIResponse{data=Task}
// @Failure      404 {object} APIResponse
// @Failure      500 {object} APIResponse
// @Router       /tasks/{id} [get]
func getTask(w http.ResponseWriter, r *http.Request, id int) {
	var task Task
	err := db.QueryRow("SELECT id, title, description, completed, created_at, updated_at FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		logError("Tarefa não encontrada: ID " + strconv.Itoa(id))
		respondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
		return
	}

	logSuccess("Tarefa recuperada com sucesso: " + task.Title)
	respondWithJSON(w, http.StatusOK, APIResponse{Success: true, Data: task})
}

// createTask godoc
// @Summary      Cria uma nova tarefa
// @Description  Cria uma nova tarefa com título e descrição. O campo 'title' é obrigatório e único.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task body CreateTaskRequest true "Dados da tarefa"
// @Success      201 {object} APIResponse{data=Task}
// @Failure      400 {object} APIResponse
// @Failure      409 {object} APIResponse
// @Failure      500 {object} APIResponse
// @Router       /tasks [post]
func createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Corpo da solicitação inválido")
		return
	}

	// Validação do campo obrigatório
	if strings.TrimSpace(req.Title) == "" {
		respondWithError(w, http.StatusBadRequest, "O campo 'title' é obrigatório")
		return
	}

	// Verificar duplicidade de título
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE title = $1)", req.Title).Scan(&exists)
	if err != nil {
		logError("Erro ao verificar duplicidade de título: " + err.Error())
		respondWithError(w, http.StatusInternalServerError, "Erro ao verificar duplicidade de título")
		return
	}
	if exists {
		logWarning("Tentativa de criar tarefa duplicada: " + req.Title)
		respondWithError(w, http.StatusConflict, "Já existe uma tarefa com esse título")
		return
	}

	var task Task
	err = db.QueryRow(
		"INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id, title, description, completed, created_at, updated_at",
		req.Title,
		req.Description,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			logWarning("Tentativa de criar tarefa duplicada (banco): " + req.Title)
			respondWithError(w, http.StatusConflict, "Já existe uma tarefa com esse título")
			return
		}
		logError("Falha ao criar tarefa: " + err.Error())
		respondWithError(w, http.StatusInternalServerError, "Falha ao criar tarefa")
		return
	}
	logSuccess("Tarefa criada com sucesso: " + task.Title)
	respondWithJSON(w, http.StatusCreated, APIResponse{Success: true, Data: task})
}

// updateTask godoc
// @Summary      Atualiza uma tarefa
// @Description  Atualiza os dados de uma tarefa existente
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da tarefa"
// @Param        task body UpdateTaskRequest true "Dados para atualização"
// @Success      200 {object} APIResponse{data=Task}
// @Failure      400 {object} APIResponse
// @Failure      404 {object} APIResponse
// @Failure      500 {object} APIResponse
// @Router       /tasks/{id} [put]
func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logError("Corpo da solicitação inválido ao atualizar tarefa: " + err.Error())
		respondWithError(w, http.StatusBadRequest, "Corpo da solicitação inválido")
		return
	}

	var task Task
	err := db.QueryRow("SELECT id, title, description, completed, created_at, updated_at FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		logError("Tarefa não encontrada para atualização: ID " + strconv.Itoa(id))
		respondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
		return
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}
	task.UpdatedAt = time.Now()

	_, err = db.Exec(
		"UPDATE tasks SET title = $1, description = $2, completed = $3, updated_at = $4 WHERE id = $5",
		task.Title,
		task.Description,
		task.Completed,
		task.UpdatedAt,
		id,
	)
	if err != nil {
		logError("Falha ao atualizar tarefa: " + err.Error())
		respondWithError(w, http.StatusInternalServerError, "Falha ao atualizar tarefa")
		return
	}

	logSuccess("Tarefa atualizada com sucesso: " + task.Title)
	respondWithJSON(w, http.StatusOK, APIResponse{Success: true, Data: task})
}

// deleteTask godoc
// @Summary      Remove uma tarefa
// @Description  Deleta uma tarefa pelo ID
// @Tags         tasks
// @Produce      json
// @Param        id path int true "ID da tarefa"
// @Success      200 {object} APIResponse
// @Failure      404 {object} APIResponse
// @Failure      500 {object} APIResponse
// @Router       /tasks/{id} [delete]
func deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	result, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		logError("Falha ao deletar tarefa: " + err.Error())
		respondWithError(w, http.StatusInternalServerError, "Falha ao deletar tarefa")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logError("Falha ao obter linhas afetadas ao deletar tarefa: " + err.Error())
		respondWithError(w, http.StatusInternalServerError, "Falha ao obter linhas afetadas")
		return
	}

	if rowsAffected == 0 {
		logWarning("Tentativa de deletar tarefa inexistente: ID " + strconv.Itoa(id))
		respondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
		return
	}

	logSuccess("Tarefa deletada com sucesso: ID " + strconv.Itoa(id))
	respondWithJSON(w, http.StatusOK, APIResponse{Success: true, Message: "Tarefa deletada com sucesso"})
}

// respondWithError envia uma resposta de erro JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, APIResponse{Success: false, Message: message})
}

// respondWithJSON envia uma resposta JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// tasksHandler manipula solicitações GET e POST para /tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
