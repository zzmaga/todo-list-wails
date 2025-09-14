package repo

import (
	"context"
	"database/sql"
	"fmt"

	"todo-list-wails/backend/models"

	_ "github.com/lib/pq"
)

// PostgresRepo реализует TaskRepo интерфейс для PostgreSQL базы данных
type PostgresRepo struct {
	db *sql.DB
}

// NewPostgresRepo создает новый экземпляр PostgreSQL репозитория
// Принимает строку подключения к базе данных
func NewPostgresRepo(connStr string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Проверяем подключение к базе данных
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	repo := &PostgresRepo{db: db}

	// Создаем таблицу если она не существует
	if err := repo.createTable(); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return repo, nil
}

// createTable создает таблицу tasks если она не существует
func (r *PostgresRepo) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id VARCHAR(36) PRIMARY KEY,
		title VARCHAR(500) NOT NULL,
		done BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		due_at TIMESTAMP WITH TIME ZONE,
		priority VARCHAR(10) NOT NULL DEFAULT 'medium'
	);`

	_, err := r.db.Exec(query)
	return err
}

// List возвращает все задачи из базы данных
func (r *PostgresRepo) List(ctx context.Context) ([]models.Task, error) {
	query := `SELECT id, title, done, created_at, due_at, priority FROM tasks ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		var dueAt sql.NullTime

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Done,
			&task.CreatedAt,
			&dueAt,
			&task.Priority,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		// Обрабатываем nullable поле due_at
		if dueAt.Valid {
			task.DueAt = &dueAt.Time
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return tasks, nil
}

// Get возвращает задачу по ID
func (r *PostgresRepo) Get(ctx context.Context, id string) (*models.Task, error) {
	query := `SELECT id, title, done, created_at, due_at, priority FROM tasks WHERE id = $1`

	var task models.Task
	var dueAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Done,
		&task.CreatedAt,
		&dueAt,
		&task.Priority,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	// Обрабатываем nullable поле due_at
	if dueAt.Valid {
		task.DueAt = &dueAt.Time
	}

	return &task, nil
}

// Add добавляет новую задачу в базу данных
func (r *PostgresRepo) Add(ctx context.Context, t models.Task) error {
	query := `
		INSERT INTO tasks (id, title, done, created_at, due_at, priority) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	var dueAt interface{}
	if t.DueAt != nil {
		dueAt = *t.DueAt
	}

	_, err := r.db.ExecContext(ctx, query,
		t.ID,
		t.Title,
		t.Done,
		t.CreatedAt,
		dueAt,
		t.Priority,
	)

	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	return nil
}

// Update обновляет существующую задачу в базе данных
func (r *PostgresRepo) Update(ctx context.Context, t models.Task) error {
	query := `
		UPDATE tasks 
		SET title = $2, done = $3, due_at = $4, priority = $5 
		WHERE id = $1`

	var dueAt interface{}
	if t.DueAt != nil {
		dueAt = *t.DueAt
	}

	result, err := r.db.ExecContext(ctx, query,
		t.ID,
		t.Title,
		t.Done,
		dueAt,
		t.Priority,
	)

	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with id %s not found", t.ID)
	}

	return nil
}

// Delete удаляет задачу из базы данных
func (r *PostgresRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with id %s not found", id)
	}

	return nil
}

func (r *PostgresRepo) Close() error {
	return r.db.Close()
}
