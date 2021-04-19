package timeendpoint

import (
	"context"

	"github.com/brumhard/geckgo/pkg/timeservice"
	"github.com/go-kit/kit/endpoint"
)

// lists
//AddList(ctx context.Context, name string, settings ListSettings) (List, error)
type AddListRequest struct {
	Name     string `validate:"required"`
	Settings *timeservice.ListSettings
}

type AddListResponse struct {
	List *timeservice.List `json:"list"`
	Err  error             `json:"-"`
}

func (r AddListResponse) Failed() error { return r.Err }

func makeAddListEndpoint(s timeservice.Service) endpoint.Endpoint {
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
	List []timeservice.List `json:"lists"`
	Err  error              `json:"-"`
}

func (r GetListsResponse) Failed() error { return r.Err }

func makeGetListsEndpoint(s timeservice.Service) endpoint.Endpoint {
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
	List *timeservice.List `json:"list"`
	Err  error             `json:"-"`
}

func (r GetListResponse) Failed() error { return r.Err }

func makeGetListEndpoint(s timeservice.Service) endpoint.Endpoint {
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
	Settings *timeservice.ListSettings
}

type UpdateListResponse struct {
	List *timeservice.List `json:"list"`
	Err  error             `json:"-"`
}

func (r UpdateListResponse) Failed() error { return r.Err }

func makeUpdateListEndpoint(s timeservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateListRequest)
		list, err := s.UpdateList(ctx, req.ListID, req.Name, req.Settings)

		resp := UpdateListResponse{
			List: list,
			Err:  err,
		}

		return resp, nil
	}
}

//DeleteList(ctx context.Context, listID int) error
type DeleteListRequest struct {
	ListID int `validate:"gte=0"`
}

type DeleteListResponse struct {
	Err error `json:"-"`
}

func (r DeleteListResponse) Failed() error { return r.Err }

func makeDeleteListEndpoint(s timeservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteListRequest)
		err := s.DeleteList(ctx, req.ListID)

		return DeleteListResponse{
			Err: err,
		}, nil
	}
}
