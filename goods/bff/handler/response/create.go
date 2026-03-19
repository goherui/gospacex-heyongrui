package response

type Create struct {
	Code int64  `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg" binding:"required"`
}
type Alipay struct {
	Code    int64  `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg     string `form:"msg" json:"msg" xml:"msg" binding:"required"`
	OrderSn string `form:"orderSn" json:"orderSn" xml:"orderSn"  binding:"required"`
	Url     string `form:"url" json:"url" xml:"url" binding:"required"`
}
