package model
type UserLogin struct{
	UserName 	string `json:"username"`
	Password	string `json:"password"`
}
type UserRegister struct {
	UserName 	string `json:"username"`
	Password1	string `json:"password1"`
	Password2	string `json:"password2"`
}