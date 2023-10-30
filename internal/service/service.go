package service

import (
	"example.com/dynamicWordpressBuilding/internal/repository"
	"example.com/dynamicWordpressBuilding/utils"
)

type Service struct {
	repo       repository.RepoInterface
	tokenMaker utils.Maker
}

type ServiceInterface interface {
	IUser
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	svc := &Service{}
	svc.repo = repo
	svc.repo = repository.NewRepo()
	svc.tokenMaker = utils.NewTokenMaker()
	return svc
}
