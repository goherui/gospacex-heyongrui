package pay

import (
	"gospacex/goods/bff/basic/config"
	payment "gospacex/proto/payment"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaymentNotify(c *gin.Context) {
	orderSn := c.Query("out_trade_no")
	if orderSn == "" {
		orderSn = c.Query("orderSn")
	}
	if orderSn == "" {
		orderSn = c.PostForm("out_trade_no")
	}
	if orderSn == "" {
		orderSn = c.PostForm("orderSn")
	}
	if orderSn == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少订单号",
		})
		return
	}
	resp, err := config.PaymentClient.HandlePaymentNotify(c, &payment.PaymentNotifyRequest{
		OrderSn: orderSn,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
		return
	}
	if !resp.Success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
