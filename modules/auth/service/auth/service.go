package auth

import "example/modules/auth/repository"

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(username, password string) (string, error) {
	token, err := s.repo.Authenticate(username, password)
	if err != nil {
		return "", err
	}

	// // Логирование действия через auditService
	// s.auditService.Log("login", map[string]interface{}{
	//     "username": username,
	// })

	return token, nil
}
