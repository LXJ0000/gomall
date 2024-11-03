package notify

import (
	"fmt"

	"github.com/LXJ0000/gomall/app/email/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct {
}

func NewNooEmail() *NoopEmail {
	return &NoopEmail{}
}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("%v\n", req)
	fmt.Println("========================================")
	fmt.Println("NoopEmail")
	fmt.Println("========================================")
	return nil
}
