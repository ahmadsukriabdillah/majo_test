package services

import (
	"errors"
	"majo_test/repository"
	"majo_test/utils"
)

type UserService interface {
	DoLogin(username string, password string) (string, error)
	IsAllowAccess(user_id uint, merchant_id string, outlet_id string) error
}

type userService struct {
	token  utils.Maker
	config utils.Config
	r      repository.UserRepository
}

func NewUserServicec(r repository.UserRepository, token utils.Maker, config utils.Config) *userService {
	return &userService{r: r, token: token, config: config}
}

func (u *userService) DoLogin(username string, password string) (string, error) {
	user, err := u.r.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if utils.IsNotEqualMd5(password, user.Password) {
		return "", errors.New("error: invalid username or password")
	}
	token, err := u.token.CreateToken(user.Id, user.Username, u.config.AccessTokenDuration)
	if err != nil {
		return "", errors.New("error: internal server error")
	}

	return token, nil
}

func (u *userService) IsAllowAccess(user_id uint, merchant_id string, outlet_id string) error {
	merchant_list, _ := u.r.GetAccessMerchant(user_id)
	if !utils.Contains(merchant_list, merchant_id) {
		return errors.New("error: user doesn't have access in this merchant")
	}

	outlet_list, _ := u.r.GetAccessOutlet(user_id)
	if !utils.Contains(outlet_list, outlet_id) {
		return errors.New("error: user doesn't have access in this Outlet")
	}
	return nil
}
