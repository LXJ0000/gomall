package service

import (
	"context"
	"errors"

	"github.com/LXJ0000/gomall/app/user/biz/dal/mysql"
	user "github.com/LXJ0000/gomall/app/user/kitex_gen/user"
	"github.com/LXJ0000/gomall/app/user/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	// Check if the email has been registered
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		return nil, errors.New("email, password and confirm password are required")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and confirm password are not the same")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	if err := model.CreateUser(mysql.DB, &newUser); err != nil {
		return nil, err
	}
	return &user.RegisterResp{
		UserId: int32(newUser.ID),
	}, nil
}
