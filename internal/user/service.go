package user

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

func (s *userService) Create(req CreateUserRequest) (*User, error) {

	user := &User{
		Name:    req.Name,
		Email:   req.Email,
		Mobile:  req.Mobile,
		Address: &req.Address,
		Status:  1,
	}

	user, err := s.repo.Create(user)
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
