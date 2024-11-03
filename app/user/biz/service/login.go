package service

import (
	"context"
	"errors"

	"github.com/LXJ0000/gomall/app/user/biz/dal/mysql"
	"github.com/LXJ0000/gomall/app/user/model"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user" // 这里使用远程库
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}
	// Check if the email has been registered
	row, err := model.GetUserByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password is incorrect")
	}
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}
	return
}
