package storage

import (
	"errors"

	"github.com/rzashakh/todoro/models"
)

type MemoryStore struct {
	tasks  []models.Task
	nextID int
}

func InitMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (ms *MemoryStore) GetAll() []models.Task {
	return ms.tasks
}

func (ms *MemoryStore) Add(task models.Task) models.Task {
	task.ID = ms.nextID
	ms.nextID++
	ms.tasks = append(ms.tasks, task)
	return task
}

func (ms *MemoryStore) Update(id int, task models.Task) (models.Task, error) {
	for i, t := range ms.tasks {
		if t.ID == id {
			task.ID = id
			ms.tasks[i] = task
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (ms *MemoryStore) Delete(id int) error {
	for i, t := range ms.tasks {
		if t.ID == id {
			ms.tasks = append(ms.tasks[:i], ms.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
