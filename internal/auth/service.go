package auth

import (
	"fmt"

	"github.com/aralim11/go-crm-api/internal/utils/jwtToken"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LoginCheck(data LoginRequest) (*LoginResponse, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (r *authService) LoginCheck(req LoginRequest) (*LoginResponse, error) {
	// repository function
	user, err := r.repo.FindByEmail(req.Email)
	if err != nil || user == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// compare password (bcrypt)
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// 4. generate JWT token
	token, err := jwtToken.GenerateJWT(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &LoginResponse{
		AccessToken: token,
	}, nil
}
