package tasks_repository

import (
	core_db "github.com/Fitray/go_final_project/internal/core/db"
)

type TasksRepository struct {
	DB core_db.Database
}

func NewTasksRepository(db core_db.Database) TasksRepository {
	return TasksRepository{DB: db}
}
