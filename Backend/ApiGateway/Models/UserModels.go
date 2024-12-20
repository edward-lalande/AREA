package models

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type SignUp struct {
	Mail         string `json:"mail"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
}
