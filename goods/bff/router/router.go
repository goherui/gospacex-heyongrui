package router

import (
	"gospacex/goods/bff/post/api/v1/pay"
	"gospacex/goods/bff/post/api/v1/product"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/product/create", product.ProductCreate)
	r.GET("/callback", pay.PaymentNotify)
	r.POST("/notify/pay", pay.PaymentNotify)
	r.POST("/alipay", pay.Alipay)
	return r
}
