package request

type Create struct {
	Title       string  `form:"title" json:"title" xml:"title"  binding:"required"`
	Price       float32 `form:"price" json:"price" xml:"price" binding:"required"`
	InventoryId int64   `form:"inventoryId" json:"inventoryId" xml:"inventoryId"  binding:"required"`
	TypeId      int64   `form:"typeId" json:"typeId" xml:"typeId"  binding:"required"`
	Img         string  `form:"img" json:"img" xml:"img"  binding:"required"`
}
type Alipay struct {
	UserId    int   `form:"userId" json:"userId" xml:"userId"  binding:"required"`
	ProductId int   `form:"productId" json:"productId" xml:"productId" binding:"required"`
	Quantity  int64 `form:"quantity" json:"quantity" xml:"quantity"  binding:"required"`
}
