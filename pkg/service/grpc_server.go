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
	list, err := s.service.AddList(ctx, req.Name, UnmarshalListSettings(req.Settings))
	if err != nil {
		// TODO: error handling
		return nil, err
	}

	return &geckgov1.AddListResponse{List: MarshalList(list)}, nil
}

func (s *Server) GetLists(ctx context.Context, _ *geckgov1.GetListsRequest) (*geckgov1.GetListsResponse, error) {
	lists, err := s.service.GetLists(ctx)
	if err != nil {
		return nil, err
	}

	retLists := make([]*geckgov1.List, 0, len(lists))
	for _, l := range lists {
		retLists = append(retLists, MarshalList(&l))
	}

	return &geckgov1.GetListsResponse{Lists: retLists}, nil
}

func (s *Server) GetList(ctx context.Context, req *geckgov1.GetListRequest) (*geckgov1.GetListResponse, error) {
	list, err := s.service.GetList(ctx, int(req.ListId))
	if err != nil {
		return nil, err
	}

	return &geckgov1.GetListResponse{List: MarshalList(list)}, nil
}

func (s *Server) UpdateList(ctx context.Context, req *geckgov1.UpdateListRequest) (*geckgov1.UpdateListResponse, error) {
	panic("implement me")
}

func (s *Server) DeleteList(ctx context.Context, req *geckgov1.DeleteListRequest) (*geckgov1.DeleteListResponse, error) {
	err := s.service.DeleteList(ctx, int(req.ListId))
	if err != nil {
		return nil, err
	}

	return &geckgov1.DeleteListResponse{}, nil
}

func (s *Server) AddDay(ctx context.Context, req *geckgov1.AddDayRequest) (*geckgov1.AddDayResponse, error) {
	moments := make([]Moment, 0, len(req.Moments))
	for _, m := range req.Moments {
		moments = append(moments, *UnmarshalMoment(m))
	}

	day, err := s.service.AddDay(ctx, int(req.ListId), req.Date.AsTime(), moments)
	if err != nil {
		return nil, err
	}

	return &geckgov1.AddDayResponse{Day: MarshalDay(day)}, nil
}

func (s *Server) GetDays(ctx context.Context, req *geckgov1.GetDaysRequest) (*geckgov1.GetDaysResponse, error) {
	days, err := s.service.GetDays(ctx, int(req.ListId))
	if err != nil {
		return nil, err
	}

	retDays := make([]*geckgov1.Day, 0, len(days))
	for _, d := range days {
		retDays = append(retDays, MarshalDay(&d))
	}

	return &geckgov1.GetDaysResponse{Days: retDays}, nil
}

func (s *Server) GetDay(ctx context.Context, req *geckgov1.GetDayRequest) (*geckgov1.GetDayResponse, error) {
	day, err := s.service.GetDay(ctx, int(req.ListId), req.Date.AsTime())
	if err != nil {
		return nil, err
	}

	return &geckgov1.GetDayResponse{Day: MarshalDay(day)}, nil
}

func (s *Server) UpdateDay(ctx context.Context, req *geckgov1.UpdateDayRequest) (*geckgov1.UpdateDayResponse, error) {
	panic("implement me")
}

func (s *Server) DeleteDay(ctx context.Context, req *geckgov1.DeleteDayRequest) (*geckgov1.DeleteDayResponse, error) {
	err := s.service.DeleteDay(ctx, int(req.ListId), req.Date.AsTime())
	if err != nil {
		return nil, err
	}

	return &geckgov1.DeleteDayResponse{}, nil
}

func (s *Server) StartDay(ctx context.Context, req *geckgov1.StartDayRequest) (*geckgov1.StartDayResponse, error) {
	err := s.service.StartDay(ctx, int(req.ListId), req.Time.AsTime())
	if err != nil {
		return nil, err
	}

	return &geckgov1.StartDayResponse{}, nil
}

func (s *Server) StartBreak(ctx context.Context, req *geckgov1.StartBreakRequest) (*geckgov1.StartBreakResponse, error) {
	err := s.service.StartBreak(ctx, int(req.ListId), req.Time.AsTime())
	if err != nil {
		return nil, err
	}

	return &geckgov1.StartBreakResponse{}, nil
}

func (s *Server) EndBreak(ctx context.Context, req *geckgov1.EndBreakRequest) (*geckgov1.EndBreakResponse, error) {
	err := s.service.EndBreak(ctx, int(req.ListId), req.Time.AsTime())
	if err != nil {
		return nil, err
	}

	return &geckgov1.EndBreakResponse{}, nil
}

func (s *Server) EndDay(ctx context.Context, req *geckgov1.EndDayRequest) (*geckgov1.EndDayResponse, error) {
	err := s.service.EndDay(ctx, int(req.ListId), req.Time.AsTime())
	if err != nil {
		return nil, err
	}

	return &geckgov1.EndDayResponse{}, nil
}
