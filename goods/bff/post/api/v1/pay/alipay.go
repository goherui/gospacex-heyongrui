package pay

import (
	"gospacex/goods/bff/basic/config"
	"gospacex/goods/bff/handler/request"
	"gospacex/goods/bff/handler/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	payment "gospacex/proto/payment"
)

func Alipay(c *gin.Context) {
	var form request.Alipay
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := config.PaymentClient.Payment(c, &payment.PaymentReq{
		UserId:    int64(form.UserId),
		ProductId: int64(form.ProductId),
		Quantity:  form.Quantity,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.Alipay{
		Code:    r.Code,
		Msg:     r.Msg,
		OrderSn: r.OrderSn,
		Url:     r.Url,
	})
}
