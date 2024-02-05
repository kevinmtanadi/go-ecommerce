package dtos

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRes struct{}

type VerifyUserReq struct {
	UserID int `query:"user_id" json:"user_id"`
}

type VerifyUserRes struct{}
