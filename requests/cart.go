package requests

type CreateCart struct {
	Items []Items `form:"items" json:"items" binding:"required"`
}

type Checkout struct {
	Items []Items `form:"items" json:"items" binding:"required"`
}
