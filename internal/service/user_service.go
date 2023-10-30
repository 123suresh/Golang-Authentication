package service

import (
	"net/http"
	"os"
	"time"

	"example.com/dynamicWordpressBuilding/internal/model"
	"example.com/dynamicWordpressBuilding/utils"
)

type IUser interface {
	CreateUser(req *model.UserRequest) (*model.UserResponse, int, error)
	GetAllUser() ([]model.UserResponse, int, error)
	GetUser(uid int) (*model.UserResponse, int, error)
	DeleteUser(uid int) (int, error)
	LoginUser(loginReq *model.LoginRequest) (*model.LoginToken, int, error)
}

func (s Service) CreateUser(req *model.UserRequest) (*model.UserResponse, int, error) {
	user := model.NewUser(req)
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

func (s Service) LoginUser(loginReq *model.LoginRequest) (*model.LoginToken, int, error) {
	userData, err := s.repo.LoginUser(loginReq.Email)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	err = utils.CheckPassword(loginReq.Password, userData.Password)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	accessTokenDuration := os.Getenv("ACCESS_TOKEN_DURATION")
	duration, err := time.ParseDuration(accessTokenDuration)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	//then we can use CreateToken function
	accessToken, err := s.tokenMaker.CreateToken(userData.Email, duration)

	//for using struct make model and use like this
	// rrr := utils.JWTMaker{}
	// accessToken, err := rrr.CreateToken(userData.Email, duration)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := &model.LoginToken{
		AccessToken: accessToken,
		User:        userData.UserRes(),
	}
	return response, http.StatusOK, nil
}