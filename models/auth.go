package models

type InputSignUp struct {
	Login    string
	Password string
	Name     string
	Email    string
}

type OutputSignUp struct {
	Session string `json:"session"`
}

type InputSignIn struct {
	Login    string
	Password string
}

type OutputSignIn struct {
	Session string  `json:"session"`
	Account Account `json:"account"`
}
