package contracts

type GetUserRes struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Role     string `json:"role" bson:"role" binding:"required"` // TODO: Use enum instead of string
	Username string `json:"username" bson:"_id" binding:"required"`
}
