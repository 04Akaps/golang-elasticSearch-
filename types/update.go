package types

type UpdateUserReq struct {
	Name string `uri:"name"`
	Age  int64  `uri:"age"`
}
