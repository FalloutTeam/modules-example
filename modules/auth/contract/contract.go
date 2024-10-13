package contract

import "example/modules/auth/service/user"

type AuthContract interface {
	GetUser(id int64) string
}

type authContract struct {
	userService user.UserService
}

func NewAuthContract() *authContract {
	return new(authContract)
}

func (c *authContract) GetUser(id int64) string {
	user, err := c.userService.User(id)
	if err != nil {
		return "нет юзера"
	}
	return user
}
