package timeendpoint

import (
	"context"
	"time"

	"github.com/brumhard/geckgo/pkg/timeservice"

	"github.com/go-kit/kit/endpoint"
)

//AddDay(ctx context.Context, listID int, date time.Time, moments []Moment) (Day, error)
type AddDayRequest struct {
	ListID  int `validate:"gte=0"`
	Date    time.Time
	Moments []timeservice.Moment `validate:"gt=0"`
}

type AddDayResponse struct {
	Day *timeservice.Day `json:"day"`
	Err error            `json:"-"`
}

func (r AddDayResponse) Failed() error { return r.Err }

func makeAddDayEndpoint(s timeservice.Service) endpoint.Endpoint {
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
	Days []timeservice.Day `json:"days"`
	Err  error             `json:"-"`
}

func (r GetDaysResponse) Failed() error { return r.Err }

func makeGetDaysEndpoint(s timeservice.Service) endpoint.Endpoint {
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
	Day *timeservice.Day `json:"day"`
	Err error            `json:"-"`
}

func (r GetDayResponse) Failed() error { return r.Err }

func makeGetDayEndpoint(s timeservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDayRequest)
		day, err := s.GetDay(ctx, req.ListID, req.Date)

		return GetDayResponse{
			Day: day,
			Err: err,
		}, nil
	}
}

//UpdateDay(ctx context.Context, listID int, date time.Time, moments []Moment) (Day, error)
type UpdateDayRequest struct {
	ListID  int `validate:"gte=0"`
	Date    time.Time
	Moments []timeservice.Moment
}

type UpdateDayResponse struct {
	Days *timeservice.Day `json:"days"`
	Err  error            `json:"-"`
}

func (r UpdateDayResponse) Failed() error { return r.Err }

func makeUpdateDayEndpoint(s timeservice.Service) endpoint.Endpoint {
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
	Err error `json:"-"`
}

func (r DeleteDayResponse) Failed() error { return r.Err }

func makeDeleteDayEndpoint(s timeservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteDayRequest)
		err := s.DeleteDay(ctx, req.ListID, req.Date)

		return DeleteDayResponse{
			Err: err,
		}, nil
	}
}
