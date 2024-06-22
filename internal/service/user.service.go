package service

import (
	"GO-ECOMMERCE-BACKEND-API/internal/model"
	_ "GO-ECOMMERCE-BACKEND-API/internal/model"
	"GO-ECOMMERCE-BACKEND-API/internal/repo"
	"GO-ECOMMERCE-BACKEND-API/internal/utils"
	"errors"
)

type UserService struct {
	UserRepository *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) Register(email, password string) error {
	if _, err := s.UserRepository.FindByEmail(email); err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	return s.UserRepository.CreateUser(user)
}
