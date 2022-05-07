package models

type InputRecoverAccessLink struct {
	Account string `json:"account" binding:"required"`
}

type InputRecoverAccessRegister struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OutputRecoverAccessRegister struct {
	Message string `json:"message"`
}
