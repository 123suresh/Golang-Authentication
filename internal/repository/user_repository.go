package repository

import (
	"fmt"

	"example.com/dynamicWordpressBuilding/internal/model"
)

type UserInterface interface {
	CreateUser(req *model.User) (*model.User, error)
	GetAllUser() ([]model.User, error)
	GetUser(uid int) (*model.User, error)
	DeleteUser(uid int) error
}

func (r *Repo) CreateUser(data *model.User) (*model.User, error) {
	err := r.db.Model(&model.User{}).Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repo) GetAllUser() ([]model.User, error) {
	users := []model.User{}
	err := r.db.Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repo) GetUser(uid int) (*model.User, error) {
	data := &model.User{}
	err := r.db.Model(&model.User{}).Where("id=?", uid).Take(&data).Error
	if err != nil {
		return nil, fmt.Errorf("user doesn't exists %v ", err)
	}
	return data, nil
}

func (r *Repo) DeleteUser(uid int) error {
	data := &model.User{}
	err := r.db.Model(&model.User{}).Where("id=?", uid).Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
