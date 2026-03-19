package product

import (
	"context"
	"gospacex/goods/goods-service/basic/config"
	"gospacex/goods/model"
	product "gospacex/proto/product"
	"net/http"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	product.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) ProductCreate(_ context.Context, in *product.ProductCreateReq) (*product.ProductCreateResp, error) {
	var products model.Product
	err := products.FindTitle(config.DB, in.Title)
	if err == nil {
		return &product.ProductCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "商品已存在",
		}, nil
	}
	products = model.Product{
		Title:       in.Title,
		Price:       float64(in.Price),
		InventoryId: int(in.InventoryId),
		TypeId:      int(in.TypeId),
		Img:         in.Img,
	}
	err = products.ProductCreate(config.DB)
	if err != nil {
		return &product.ProductCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "上架失败",
		}, nil
	}
	return &product.ProductCreateResp{
		Code: http.StatusOK,
		Msg:  "上架成功",
	}, nil
}
