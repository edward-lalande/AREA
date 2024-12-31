package models

type DatabaseActions struct {
	Id         int
	UserToken  string
	AreaId     string
	ActionType int
	NbMessage  int
	NbEvents   int
}
