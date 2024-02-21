package models

type User struct {
	Id            int32  `json:"id"`
	Email         string `json:"email"`
	HashePassword string `json:"hashePassword"`
}
type UserDTO struct {
	Id    int32  `json:"id"`
	Email string `json:"email"`
}
