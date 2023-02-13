package rpc

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/server"
	"github.com/bobgo0912/bob-armory/pkg/order"
	"github.com/bobgo0912/bob-order/interal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderRpcServer struct {
	order.UnimplementedOrderServer
}

func RegService(s *server.GrpcServer) {
	s.RegService(&order.Order_ServiceDesc, &OrderRpcServer{})
}

func (o *OrderRpcServer) GetOrdersByIds(ctx context.Context, req *order.OrderIdsReq) (*order.OrderResp, error) {
	if len(req.Ids) < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "ids is empty")
	}
	store, err := repo.GetOrderStore()
	if err != nil {
		log.Error("GetOrderStore fail")
		return nil, status.Errorf(codes.Internal, "get db fail")
	}
	list, err := store.QueryList(ctx, squirrel.Select("*").Where(squirrel.Eq{"id": req.Ids}))
	if err != nil {
		log.Error("QueryList fail err=", err)
		return nil, status.Errorf(codes.Internal, "query fail")
	}
	s := make([]*order.OrderS, 0, len(list))
	for _, o2 := range list {
		orderS := order.OrderS{
			Id:         o2.Id,
			Period:     o2.Period,
			CardNumber: int32(o2.CardNumber),
			PlayerId:   o2.PlayerId,
		}
		s = append(s, &orderS)
	}
	return &order.OrderResp{Data: s}, nil

}
func (o *OrderRpcServer) GetOrdersByPeriod(ctx context.Context, req *order.OrderPeriodReq) (*order.OrderResp, error) {
	if req.GetPeriod() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "period is empty")
	}
	store, err := repo.GetOrderStore()
	if err != nil {
		log.Error("GetOrderStore fail")
		return nil, status.Errorf(codes.Internal, "get db fail")
	}
	list, err := store.QueryList(ctx, squirrel.Select("*").Where(squirrel.Eq{"period": req.Period}))
	if err != nil {
		log.Error("QueryList fail err=", err)
		return nil, status.Errorf(codes.Internal, "query fail")
	}
	s := make([]*order.OrderS, 0, len(list))
	for _, o2 := range list {
		orderS := order.OrderS{
			Id:         o2.Id,
			Period:     o2.Period,
			CardNumber: int32(o2.CardNumber),
			PlayerId:   o2.PlayerId,
		}
		s = append(s, &orderS)
	}
	return &order.OrderResp{Data: s}, nil
}
func (o *OrderRpcServer) GetCardsByPeriod(ctx context.Context, req *order.CardPeriodReq) (*order.CardResp, error) {
	if req.GetPeriod() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "period is empty")
	}
	store, err := repo.GetCardStore()
	if err != nil {
		log.Error("GetCardStore fail")
		return nil, status.Errorf(codes.Internal, "get db fail")
	}
	list, err := store.QueryList(ctx, squirrel.Select("*").Where(squirrel.Eq{"period": req.Period}))
	if err != nil {
		log.Error("QueryList fail err=", err)
		return nil, status.Errorf(codes.Internal, "query fail")
	}
	s := make([]*order.Card, 0, len(list))
	for _, o2 := range list {
		orderS := order.Card{
			Id:       o2.Id,
			Period:   o2.Period,
			Numbers:  o2.Numbers,
			PlayerId: o2.PlayerId,
		}
		s = append(s, &orderS)
	}
	return &order.CardResp{Data: s}, nil
}
