package models

type UsersGet struct {
	RoutesWanted string `json:"routes"`
}

type UserInformation struct {
	RoutesWanted string `json:"routes"`
	Id           int    `json:"id"`
	Mail         string `json:"mail"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
}
