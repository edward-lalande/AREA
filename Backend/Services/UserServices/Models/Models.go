package models

type LoginId struct {
	Id int `json:"id"`
}

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type User struct {
	Id    int    `json:"id"`
	Login string `json:"name"`
}

type UserInformation struct {
	Id       int    `json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
