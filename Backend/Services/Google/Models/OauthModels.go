package models

type OauthInformation struct {
	Code string `json:"code"`
}

type OauthInformationToken struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}
