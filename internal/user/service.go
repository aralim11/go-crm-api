package user

import "fmt"

type UserService interface {
	Create(req CreateUserRequest) (*User, error)
	List() ([]*UserResponse, error)
	GetUserByID(id int64) (*UserResponse, error)
	UpdateUser(user *UpdateUserRequest, id int64) (*UserResponse, error)
	DeleteUser(id int64) error
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

	// check if mobile already exist
	existingMobile, err := s.repo.FindByMobile(req.Mobile)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing mobile: %w", err)
	}

	if existingMobile != nil {
		return nil, fmt.Errorf("mobile already exist")
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
func (s *userService) List() ([]*UserResponse, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	return users, nil
}

/**
 * Get user by ID
 */
func (s *userService) GetUserByID(id int64) (*UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

/**
 * Update user
 */
func (s *userService) UpdateUser(user *UpdateUserRequest, id int64) (*UserResponse, error) {
	var updateUserReq UpdateUserRequest
}

/**
 * Delete user
 */
func (s *userService) DeleteUser(id int64) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
