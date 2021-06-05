package service

import (
	"context"

	"github.com/brumhard/geckgo/pkg/pb/geckgo/v1"
	"go.uber.org/zap"
)

type Server struct {
	geckgov1.UnimplementedGeckgoServiceServer
	service Service
}

func NewServer(repo Repository, logger *zap.Logger) *Server {
	return &Server{service: NewService(repo, logger)}
}

func (s *Server) AddList(ctx context.Context, req *geckgov1.AddListRequest) (*geckgov1.AddListResponse, error) {
	panic("implement me")
}

func (s *Server) GetLists(ctx context.Context, _ *geckgov1.GetListsRequest) (*geckgov1.GetListsResponse, error) {
	panic("implement me")
}

func (s *Server) GetList(ctx context.Context, req *geckgov1.GetListRequest) (*geckgov1.GetListResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateList(ctx context.Context, req *geckgov1.UpdateListRequest) (*geckgov1.UpdateListResponse, error) {
	panic("implement me")
}

func (s *Server) DeleteList(ctx context.Context, req *geckgov1.DeleteListRequest) (*geckgov1.DeleteListResponse, error) {
	panic("implement me")
}

func (s *Server) AddDay(ctx context.Context, req *geckgov1.AddDayRequest) (*geckgov1.AddDayResponse, error) {
	panic("implement me")
}

func (s *Server) GetDays(ctx context.Context, req *geckgov1.GetDaysRequest) (*geckgov1.GetDaysResponse, error) {
	panic("implement me")
}

func (s *Server) GetDay(ctx context.Context, req *geckgov1.GetDayRequest) (*geckgov1.GetDayResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateDay(ctx context.Context, req *geckgov1.UpdateDayRequest) (*geckgov1.UpdateDayResponse, error) {
	panic("implement me")
}

func (s *Server) DeleteDay(ctx context.Context, req *geckgov1.DeleteDayRequest) (*geckgov1.DeleteDayResponse, error) {
	panic("implement me")
}

func (s *Server) StartDay(ctx context.Context, req *geckgov1.StartDayRequest) (*geckgov1.StartDayResponse, error) {
	panic("implement me")
}

func (s *Server) StartBreak(ctx context.Context, req *geckgov1.StartBreakRequest) (*geckgov1.StartBreakResponse, error) {
	panic("implement me")
}

func (s *Server) EndBreak(ctx context.Context, req *geckgov1.EndBreakRequest) (*geckgov1.EndBreakResponse, error) {
	panic("implement me")
}

func (s *Server) EndDay(ctx context.Context, req *geckgov1.EndDayRequest) (*geckgov1.EndDayResponse, error) {
	panic("implement me")
}
