package auth

import (
	"auth/internal/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type HandleAuth struct {
	serviceAuth *service.AuthService
}

func NewHandleAuth(serviceAuth *service.AuthService) *HandleAuth {
	return &HandleAuth{serviceAuth: serviceAuth}
}

func (h *HandleAuth) Register(w http.ResponseWriter, r *http.Request) {
	email := "qwer"
	password := "asdf"
	hashepassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("err generate hashedpassword")
	}

	mess, err := h.serviceAuth.Register(email, string(hashepassword))
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(mess))

}

func (h *HandleAuth) Login(w http.ResponseWriter, r *http.Request) {
	email := "qwer"
	password := "asdf"

	token, err := h.serviceAuth.Login(email, password)
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
