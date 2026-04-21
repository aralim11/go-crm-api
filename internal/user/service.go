package user

import "fmt"

type UserService interface {
	Create(req CreateUserRequest) (*User, error)
	List() ([]*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

/**
 * Create a new user
 */
func (s *userService) Create(req CreateUserRequest) (*User, error) {
	// check if email already exists
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// create user object
	user := &User{
		Name:    req.Name,
		Email:   req.Email,
		Mobile:  req.Mobile,
		Address: &req.Address,
		Status:  1,
	}

	user, err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

/**
 * List users
 */
func (s *userService) List() ([]*User, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	return users, nil
}
