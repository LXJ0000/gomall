package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LXJ0000/gomall/app/user/biz/dal/mysql"
	"github.com/LXJ0000/gomall/app/user/infra/mq"
	"github.com/LXJ0000/gomall/app/user/model"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/email"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
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

	// test send email
	data, _ := proto.Marshal(
		&email.EmailReq{
			To:          "xxx",
			From:        "xxx",
			ContentType: "xxx",
			Subject:     "xxx",
			Body:        "xxx",
		})
	msg := &nats.Msg{
		Subject: "email",
		Data:    data,
	}
	if err := mq.Nc.PublishMsg(msg); err != nil {
		fmt.Println(err)
	}
	resp = &user.RegisterResp{
		UserId: int32(newUser.ID),
	}
	return
}
