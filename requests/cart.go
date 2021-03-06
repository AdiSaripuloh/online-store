package requests

type CreateCart struct {
	Items []Items `form:"items" json:"items" binding:"required"`
}
