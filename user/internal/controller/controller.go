package user

import (
	"encoding/json"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"log"
	"net/http"
	"user/internal/service"
)

type HandleUser struct {
	serviceuser *service.UserService
}

func NewHandleUser(clUser *service.UserService) *HandleUser {
	return &HandleUser{clUser}
}

func (h *HandleUser) Profile(w http.ResponseWriter, r *http.Request) {
	email := "qwer"
	req := &userpr.ProfileRequest{Email: email}

	res, err := h.serviceuser.Profile(req.Email)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "err serv", http.StatusInternalServerError)
		return
	}
	jsData, err := json.Marshal(res)
	if err != nil {
		log.Println("err:", err)
	}
	w.Write(jsData)
}

func (h *HandleUser) List(w http.ResponseWriter, r *http.Request) {
	res, err := h.serviceuser.List()
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "err serv", http.StatusInternalServerError)
		return
	}
	jsData, err := json.Marshal(res)
	if err != nil {
		log.Println("err :", err)
	}
	w.Write(jsData)
}
