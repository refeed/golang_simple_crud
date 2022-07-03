package contracts

type CreateUserReq struct {
	Username string `json:"username" bson:"_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}
