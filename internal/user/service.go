package user

import "fmt"

type UserService interface {
	Create(name, email string) (*User, error)
	List() error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(name, email string) (*User, error) {
	fmt.Println("User Create Service", name, email)
	user := &User{
		Name:  name,
		Email: email,
	}
	s.repo.Create(user)
	return nil, nil
}

func (s *userService) List() error {
	return nil
}
