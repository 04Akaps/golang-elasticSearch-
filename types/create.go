package types

type CreateUserReq struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Age     int64  `json:"age" binding:"required"`
}
