package service

import "fmt"

type NotifService struct {
}

func NewNotifService() *NotifService {
	return &NotifService{}
}

func (n *NotifService) SendViaEmail(email string) (string, error) {
	resp := fmt.Sprintf("message has been sent via email: %s", email)
	return resp, nil
}
func (n *NotifService) SendViaSMS(phone string) (string, error) {
	resp := fmt.Sprintf("message has been sent via sms: %s", phone)
	return resp, nil
}
