package product

import (
	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {
	//var form request.Create
	//// This will infer what binder to use depending on the content-type header.
	//if err := c.ShouldBind(&form); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//r, err := config.ProductClient.ProductCreate(c, &__.ProductCreateReq{
	//	Title:       form.Title,
	//	Price:       form.Price,
	//	InventoryId: form.InventoryId,
	//	TypeId:      form.TypeId,
	//	Img:         form.Img,
	//})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//c.JSON(http.StatusOK,response.Create{
	//	Code: r.Code,
	//	Msg:  r.Msg,
	//})
}
