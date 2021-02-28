package admin

//AdmModel is a blueprint or struct for admin
type AdmModel struct {
	AdminID  int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
