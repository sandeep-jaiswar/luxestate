package models

type User struct {
	UID      int    `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt string `json:"salt"`
}
