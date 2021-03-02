package todo

import (
	"fmt"
	"todos/dao"
)

// Todo doc
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateATodo 创建一个 todo 事件
func CreateATodo(todo *Todo) (err error) {
	fmt.Println("todo:", todo)
	err = dao.DB.Create(&todo).Error
	return
}

// GetAllTodo 获取所有的 todo 事件
func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo 获取一个 todo，参数为 id
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateATodo 更新 todo 事件
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return err
}

// DeleteATodo 删除一个 todo，参数为 id
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
