package rest

type WebResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ReqAddActor struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
