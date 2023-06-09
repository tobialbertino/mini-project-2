package domain

type Token struct {
	Token string
}

type ResToken struct {
	AccessToken string `json:"access_token,omitempty"`
}
