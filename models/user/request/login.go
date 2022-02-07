package request

type Userlogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
