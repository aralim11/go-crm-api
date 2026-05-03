package auth

type AuthService interface {
	LoginCheck(email string) (bool, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (r *authService) LoginCheck(email string) (bool, error) {
	return true, nil
}
