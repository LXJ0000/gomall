package service

import (
	"context"

	"github.com/LXJ0000/gomall/app/product/biz/dal/mysql"
	product "github.com/LXJ0000/gomall/app/product/kitex_gen/product"
	"github.com/LXJ0000/gomall/app/product/model"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(200401, "product id is required")
	}
	query := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := query.GetByID(int(req.Id))
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(200402, "product not found")
	}
	resp = &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		},
	}
	return
}
