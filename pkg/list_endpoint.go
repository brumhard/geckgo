package pkg

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// lists
//AddList(ctx context.Context, name string, settings ListSettings) (List, error)
type addListRequest struct {
	Name     string `validate:"required"`
	Settings *ListSettings
}

type addListResponse struct {
	List *List `json:"list"`
	Err  error `json:"-"`
}

func (r addListResponse) error() error { return r.Err }

func makeAddListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addListRequest)
		list, err := s.AddList(ctx, req.Name, req.Settings)

		return addListResponse{
			List: list,
			Err:  err,
		}, nil
	}
}

//GetLists(ctx context.Context) ([]List, error)
type getListsRequest struct{}

type getListsResponse struct {
	List []List `json:"lists"`
	Err  error  `json:"-"`
}

func (r getListsResponse) error() error { return r.Err }

func makeGetListsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getListsRequest)
		lists, err := s.GetLists(ctx)

		return getListsResponse{
			List: lists,
			Err:  err,
		}, nil
	}
}

//GetList(ctx context.Context, listID int) (List, error)
type getListRequest struct {
	ListID int `validate:"gte=0"`
}

type getListResponse struct {
	List *List `json:"list"`
	Err  error `json:"-"`
}

func (r getListResponse) error() error { return r.Err }

func makeGetListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getListRequest)
		list, err := s.GetList(ctx, req.ListID)

		return getListResponse{
			List: list,
			Err:  err,
		}, nil
	}
}

//UpdateList(ctx context.Context, listID int, settings ListSettings) (List, error)
type updateListRequest struct {
	ListID   int `validate:"gte=0"`
	Name     string
	Settings *ListSettings
}

type updateListResponse struct {
	List List  `json:"list"`
	Err  error `json:"-"`
}

func (r updateListResponse) error() error { return r.Err }

func makeUpdateListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateListRequest)
		list, err := s.UpdateList(ctx, req.ListID, "arr", req.Settings)

		return updateListResponse{
			List: *list,
			Err:  err,
		}, nil
	}
}

//DeleteList(ctx context.Context, listID int) error
type deleteListRequest struct {
	ListID int `validate:"gte=0"`
}

type deleteListResponse struct {
	Err error `json:"-"`
}

func (r deleteListResponse) error() error { return r.Err }

func makeDeleteListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteListRequest)
		err := s.DeleteList(ctx, req.ListID)

		return deleteListResponse{
			Err: err,
		}, nil
	}
}
