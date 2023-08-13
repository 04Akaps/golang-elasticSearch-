package types

type SearchByNameReq struct {
	Name string `uri:"name" binding:"required"`
	Sort
	Size
}

type SearchByAgeReq struct {
	Age int64 `uri:"age" binding:"required"`
	Sort
	Size
}

type SearchByAddressReq struct {
	Address string `uri:"address" binding:"required"`
	Sort
	Size
}
