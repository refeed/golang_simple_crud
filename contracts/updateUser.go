package contracts

type UpdateUserReq struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
