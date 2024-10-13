package user

type UserService interface {
	User(id int64) (string, error)
}

type userService struct{}

func NewService() *userService {
	return &userService{}
}

func (s *userService) User(id int64) (string, error) {
	return "отдаём какого-нибудь юзера", nil
}
