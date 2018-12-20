package models

import (
	"github.com/jinzhu/gorm"
)

// Todo struct
type Todo struct {
	gorm.Model
	Value     string `json:"value"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
}

// Validate that all values are present
func (todo *Todo) Validate() {

}

// CreateTodo and save to db
func (todo *Todo) CreateTodo() {

}

// EditTodo and save to db
func (todo *Todo) EditTodo() {

}

// RemoveTodo from db
func (todo *Todo) RemoveTodo() {

}

// GetTodos - returns all for user
func GetTodos(user uint) {

}

// GetTodo by id
func GetTodo(id uint) {

}
