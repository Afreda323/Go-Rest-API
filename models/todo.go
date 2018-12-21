package models

import (
	"fmt"
	"todo/utils"

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
func (todo *Todo) Validate() (map[string]interface{}, bool) {
	if todo.Value == "" {
		return utils.Message(false, "Invalid TODO value"), false
	}

	if todo.UserID <= 0 {
		return utils.Message(false, "Invalid User"), false
	}

	return utils.Message(true, "Valid"), true
}

// CreateTodo and save to db
func (todo *Todo) CreateTodo() map[string]interface{} {
	if msg, ok := todo.Validate(); !ok {
		return msg
	}

	todo.Completed = false

	GetDB().Create(todo)

	resp := utils.Message(true, "Todo Saved")
	resp["data"] = todo

	return resp
}

// EditTodo and save to db
func (todo *Todo) EditTodo() map[string]interface{} {
	if msg, ok := todo.Validate(); !ok {
		return msg
	}
	og := &Todo{}
	err := GetDB().Table("todos").Where("id = ?", todo.ID).First(og).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		resp := utils.Message(false, "Todo not found")
		return resp
	}

	if todo.Completed != og.Completed {
		og.Completed = todo.Completed
	}

	if todo.Value != og.Value {
		og.Value = todo.Value
	}

	GetDB().Update(og)

	resp := utils.Message(true, "Todo Updated")
	resp["data"] = og

	return resp
}

// RemoveTodo from db
func RemoveTodo(id uint, user uint) map[string]interface{} {
	todo := &Todo{}
	err := GetDB().Table("todos").Where("user_id = ? AND id = ?", user, id).First(todo).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return utils.Message(false, "Todo does not exist")
	}

	GetDB().Delete(todo)

	resp := utils.Message(true, "Todo Deleted")
	return resp
}

// GetTodos - returns all for user
func GetTodos(user uint) []*Todo {
	todos := make([]*Todo, 0)
	err := GetDB().Table("todos").Where("user_id = ?", user).Find(&todos).Error
	if err != nil {
		fmt.Println("ERROR", err)
		return nil
	}
	return todos
}

// GetTodo by id
func GetTodo(id uint, user uint) map[string]interface{} {
	todo := &Todo{}
	err := GetDB().Table("todos").Where("user_id = ? AND id = ?", user, id).Find(todo).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return utils.Message(false, "Todo does not exist")
	}

	resp := utils.Message(true, "Todo Found")
	resp["data"] = todo
	return resp
}
