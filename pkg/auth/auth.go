package auth

import (
	"errors"

	"github.com/VENI-VIDIVICI/plus/app/models/user"
)

func Attempt(email string, passward string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(passward) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

func LoginPhone(phone string) (user.User, error) {
	userModel := user.GetPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}
