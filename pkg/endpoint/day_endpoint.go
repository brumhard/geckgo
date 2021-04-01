package endpoint

import (
	"context"
	"github.com/brumhard/geckgo/pkg/service"
	"time"

	"github.com/go-kit/kit/endpoint"
)

//AddDay(ctx context.Context, listID int, date time.Time, moments []Moment) (Day, error)
type AddDayRequest struct {
	ListID  int `validate:"gte=0"`
	Date    time.Time
	Moments []service.Moment
}

type AddDayResponse struct {
	Day *service.Day `json:"day"`
	Err error        `json:"-"`
}

func (r AddDayResponse) error() error { return r.Err }

func makeAddDayEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDayRequest)
		day, err := s.AddDay(ctx, req.ListID, req.Date, req.Moments)

		return AddDayResponse{
			Day: day,
			Err: err,
		}, nil
	}
}

//GetDays(ctx context.Context, listID int, opts ...ListDaysOption) ([]Day, error)
type GetDaysRequest struct {
	ListID int `validate:"gte=0"`
}

type GetDaysResponse struct {
	Days []service.Day `json:"days"`
	Err  error         `json:"-"`
}

func (r GetDaysResponse) error() error { return r.Err }

func makeGetDaysEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDaysRequest)
		// TODO: add filtering options
		days, err := s.GetDays(ctx, req.ListID)

		return GetDaysResponse{
			Days: days,
			Err:  err,
		}, nil
	}
}

//GetDay(ctx context.Context, listID int, date time.Time) (Day, error)
type GetDayRequest struct {
	ListID int `validate:"gte=0"`
	Date   time.Time
}

type GetDayResponse struct {
	Days *service.Day `json:"days"`
	Err  error        `json:"-"`
}

func (r GetDayResponse) error() error { return r.Err }

func makeGetDayEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDayRequest)
		day, err := s.GetDay(ctx, req.ListID, req.Date)

		return GetDayResponse{
			Days: day,
			Err:  err,
		}, nil
	}
}

//UpdateDay(ctx context.Context, listID int, date time.Time, moments []Moment) (Day, error)
type UpdateDayRequest struct {
	ListID  int `validate:"gte=0"`
	Date    time.Time
	Moments []service.Moment
}

type UpdateDayResponse struct {
	Days *service.Day `json:"days"`
	Err  error        `json:"-"`
}

func (r UpdateDayResponse) error() error { return r.Err }

func makeUpdateDayEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateDayRequest)
		day, err := s.UpdateDay(ctx, req.ListID, req.Date, req.Moments)

		return UpdateDayResponse{
			Days: day,
			Err:  err,
		}, nil
	}
}

//DeleteDay(ctx context.Context, listID int, date time.Time) error
type DeleteDayRequest struct {
	ListID int `validate:"gte=0"`
	Date   time.Time
}

type DeleteDayResponse struct {
	Days service.Day `json:"days"`
	Err  error       `json:"-"`
}

func (r DeleteDayResponse) error() error { return r.Err }

func makeDeleteDayEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteDayRequest)
		err := s.DeleteDay(ctx, req.ListID, req.Date)

		return DeleteDayResponse{
			Err: err,
		}, nil
	}
}
