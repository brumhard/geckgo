package endpoint

import (
	"context"
	"github.com/brumhard/geckgo/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

// lists
//AddList(ctx context.Context, name string, settings ListSettings) (List, error)
type AddListRequest struct {
	Name     string `validate:"required"`
	Settings *service.ListSettings
}

type AddListResponse struct {
	List *service.List `json:"list"`
	Err  error         `json:"-"`
}

func (r AddListResponse) error() error { return r.Err }

func makeAddListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddListRequest)
		list, err := s.AddList(ctx, req.Name, req.Settings)

		return AddListResponse{
			List: list,
			Err:  err,
		}, nil
	}
}

//GetLists(ctx context.Context) ([]List, error)
type GetListsRequest struct{}

type GetListsResponse struct {
	List []service.List `json:"lists"`
	Err  error          `json:"-"`
}

func (r GetListsResponse) error() error { return r.Err }

func makeGetListsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetListsRequest)
		lists, err := s.GetLists(ctx)

		return GetListsResponse{
			List: lists,
			Err:  err,
		}, nil
	}
}

//GetList(ctx context.Context, listID int) (List, error)
type GetListRequest struct {
	ListID int `validate:"gte=0"`
}

type GetListResponse struct {
	List *service.List `json:"list"`
	Err  error         `json:"-"`
}

func (r GetListResponse) error() error { return r.Err }

func makeGetListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetListRequest)
		list, err := s.GetList(ctx, req.ListID)

		return GetListResponse{
			List: list,
			Err:  err,
		}, nil
	}
}

//UpdateList(ctx context.Context, listID int, settings ListSettings) (List, error)
type UpdateListRequest struct {
	ListID   int `validate:"gte=0"`
	Name     string
	Settings *service.ListSettings
}

type UpdateListResponse struct {
	List service.List `json:"list"`
	Err  error        `json:"-"`
}

func (r UpdateListResponse) error() error { return r.Err }

func makeUpdateListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateListRequest)
		list, err := s.UpdateList(ctx, req.ListID, "arr", req.Settings)

		return UpdateListResponse{
			List: *list,
			Err:  err,
		}, nil
	}
}

//DeleteList(ctx context.Context, listID int) error
type DeleteListRequest struct {
	ListID int `validate:"gte=0"`
}

type DeleteListResponse struct {
	Err error `json:"-"`
}

func (r DeleteListResponse) error() error { return r.Err }

func makeDeleteListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteListRequest)
		err := s.DeleteList(ctx, req.ListID)

		return DeleteListResponse{
			Err: err,
		}, nil
	}
}
