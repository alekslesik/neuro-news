package service

import "github.com/alekslesik/neuro-news/internal/app/model"



type UserService interface {
    GetUserByID(id int) (*model.User, error)
}

type userService struct {
    UserRepository model.UserRepository
}

func NewUserService(userRepository model.UserRepository) UserService {
    return &userService{
        UserRepository: userRepository,
    }
}

func (us *userService) GetUserByID(id int) (*model.User, error) {
    return us.UserRepository.GetUserByID(id)
}
