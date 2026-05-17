package tasks_repository

import "database/sql"

type TasksRepository struct {
	DB *sql.DB
}

func NewTasksRepository(db *sql.DB) TasksRepository {
	return TasksRepository{DB: db}
}
