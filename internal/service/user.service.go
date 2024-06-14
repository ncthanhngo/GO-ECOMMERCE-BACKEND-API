package service

import "GO-ECOMMERCE-BACKEND-API/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// User repo -> ur
func (ur *UserService) GetInfoUserService() string {
	return ur.userRepo.GetInfoUser()
}
