package service

import (
	"net/http"

	"example.com/dynamicWordpressBuilding/internal/model"
	"example.com/dynamicWordpressBuilding/utils"
)

type IUser interface {
	CreateUser(req *model.UserRequest) (*model.UserResponse, int, error)
	GetAllUser() ([]model.UserResponse, int, error)
	GetUser(uid int) (*model.UserResponse, int, error)
	DeleteUser(uid int) (int, error)
}

func (s Service) CreateUser(req *model.UserRequest) (*model.UserResponse, int, error) {
	user := model.NewUser(req)
	//hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	user.Password = hashedPassword
	result, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	response := result.UserRes()
	return response, http.StatusCreated, nil
}

func (s Service) GetAllUser() ([]model.UserResponse, int, error) {
	result, err := s.repo.GetAllUser()
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	responses := []model.UserResponse{}
	for _, user := range result {
		responses = append(responses, *user.UserRes())
	}
	return responses, http.StatusOK, nil
}

func (s Service) GetUser(uid int) (*model.UserResponse, int, error) {
	result, err := s.repo.GetUser(uid)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	response := result.UserRes()
	return response, http.StatusOK, nil
}

func (s Service) DeleteUser(uid int) (int, error) {
	err := s.repo.DeleteUser(uid)
	if err != nil {
		return http.StatusNotFound, err
	}
	return http.StatusOK, nil
}
