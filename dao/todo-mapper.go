package dao

import (
	"todo-list/models"
)



func GetOption(ID interface{}) (*models.Todo, error) {
	var todo models.Todo
	if err := DB.Where("id = ?", ID).Find(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func ListOptions() (*[]models.Todo, error) {
	var todos *[]models.Todo
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func AddOption(todo *models.Todo) error {
	if err := DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOption(ID interface{}) error {
	if err := DB.Model(&models.Todo{}).Delete("id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOptionState(ID interface{}, status interface{}) error {
	if err := DB.Model(&models.Todo{}).Where("id = ?", ID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}