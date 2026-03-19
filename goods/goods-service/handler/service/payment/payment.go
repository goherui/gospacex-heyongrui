package payment

import (
	"context"
	"errors"
	"gospacex/goods/goods-service/basic/config"
	"gospacex/goods/model"
	"gospacex/pkg"
	__ "gospacex/proto/payment"
	"net/http"
)

type Saver struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Saver) Payment(_ context.Context, in *__.PaymentReq) (*__.PaymentResp, error) {
	orderSn := pkg.OrderSn()
	total := 0.0
	tx := config.DB.Begin()
	var user model.User
	err := user.FindUserById(tx, in.UserId)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("用户不存在")
	}
	var product model.Product
	err = product.FindProductById(tx, in.ProductId)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("商品不存在")
	}
	var inventory model.Inventory
	err = inventory.FindProductByStock(tx, in.ProductId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if int64(inventory.Stock) < in.Quantity {
		tx.Rollback()
		return nil, errors.New("库存不足")
	}
	total = product.Price * float64(in.Quantity)
	order := model.Order{
		OrderSn:   orderSn,
		ProductId: int(in.ProductId),
		Quantity:  int(in.Quantity),
		Total:     total,
		Status:    "未支付",
	}
	err = order.CreateOrder(tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("订单创建失败")
	}
	tx.Commit()
	url := pkg.AliPay(orderSn, total)
	return &__.PaymentResp{
		Code:    http.StatusOK,
		Msg:     "订单创建成功",
		OrderSn: orderSn,
		Url:     url,
	}, nil
}

func (s *Saver) HandlePaymentNotify(_ context.Context, in *__.PaymentNotifyRequest) (*__.PaymentNotifyResponse, error) {
	if in.OrderSn == "" {
		return &__.PaymentNotifyResponse{Success: false, Message: "订单号为空"}, nil
	}
	tx := config.DB.Begin()
	var order model.Order
	if err := tx.Where("order_sn = ?", in.OrderSn).First(&order).Error; err != nil {
		return nil, errors.New("订单不存在")
	}
	if order.Status != "待支付" {
		return nil, nil
	}
	var inventory model.Inventory

	if err := tx.Model(&inventory).Update("stock", inventory.Stock-order.Quantity).Error; err != nil {
		return nil, errors.New("扣减库存失败")
	}

	if err := tx.Model(&order).Update("status", "已支付").Error; err != nil {
		return nil, errors.New("更新订单状态失败")
	}
	return nil, nil
}

//func (s *Server) sendStockMsg(tx *gorm.DB, order model.PaymentOrder, items []*model.OrderItem) {
//	for _, item := range items {
//		var goods model.Goods
//		if tx.First(&goods, item.GoodsID).Error != nil {
//			continue
//		}
//		msg := fmt.Sprintf("订单:%s,商品%d(%s),扣减%d,剩余%d",
//			order.OrderSn, item.GoodsID, item.Title, item.Quantity, goods.Stock)
//		rmq := pkg.NewRabbitMQSimple("stock_deduct_log")
//		rmq.PublishSimple(msg)
//		rmq.Destory()
//	}
//}
//
//func StartStockConsumer() {
//	logFile, _ := os.OpenFile("stock_deduct.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	defer logFile.Close()
//	logger := log.New(logFile, "[库存扣减]", log.LstdFlags)
//	rmq := pkg.NewRabbitMQSimple("stock_deduct_log")
//	defer rmq.Destory()
//	log.Println("启动库存扣减消费者...")
//	rmq.ConsumeSimpleCallback(func(body []byte) {
//		logMsg := fmt.Sprintf("%s - %s", time.Now().Format("2006-01-02 15:04:05"), string(body))
//		logger.Println(logMsg)
//		log.Printf("处理库存消息: %s", body)
//	})
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//
//	log.Println("停止库存扣减消费者...")
//}
