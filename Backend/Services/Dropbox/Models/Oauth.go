package models

type OauthInformation struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}

type OauthInformationSignUp struct {
	UserToken   string `json:"user_token" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
}
