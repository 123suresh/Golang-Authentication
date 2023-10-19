package service

import "example.com/dynamicWordpressBuilding/internal/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	IUser
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	svc := &Service{}
	svc.repo = repo
	svc.repo = repository.NewRepo()
	return svc
}
