package form

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetInfosForm struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
