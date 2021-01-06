package pkg

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
)

// request, response
// endpoints: https://github.com/go-kit/kit/blob/master/examples/shipping/booking/endpoint.go

// TODO: check list id for day methods
//AddDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error)
type addDayRequest struct {
	ListID  string
	Date    time.Time
	Moments []Moment
}

type addDayResponse struct {
	Day Day   `json:"day"`
	Err error `json:"err"`
}

func (r addDayResponse) error() error { return r.Err }

func makeAddDayEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addDayRequest)
		day, err := s.AddDay(ctx, req.ListID, req.Date, req.Moments)

		return addDayResponse{
			Day: day,
			Err: err,
		}, nil
	}
}

//GetDays(ctx context.Context, listID string, opts ...ListDaysOption) ([]Day, error)
type getDaysRequest struct {
	ListID string
}

type getDaysResponse struct {
	Days []Day `json:"days"`
	Err  error `json:"err"`
}

func (r getDaysResponse) error() error { return r.Err }

func makeGetDaysEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getDaysRequest)
		// TODO: add filtering options
		days, err := s.GetDays(ctx, req.ListID)

		return getDaysResponse{
			Days: days,
			Err:  err,
		}, nil
	}
}

//GetDay(ctx context.Context, listID string, date time.Time) (Day, error)
type getDayRequest struct {
	ListID string
	Date   time.Time
}

type getDayResponse struct {
	Days Day   `json:"days"`
	Err  error `json:"err"`
}

func (r getDayResponse) error() error { return r.Err }

func makeGetDayEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getDayRequest)
		day, err := s.GetDay(ctx, req.ListID, req.Date)
		return getDayResponse{
			Days: day,
			Err:  err,
		}, nil
	}
}

//UpdateDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error)
type updateDayRequest struct {
	ListID  string
	Date    time.Time
	Moments []Moment
}

type updateDayResponse struct {
	Days Day   `json:"days"`
	Err  error `json:"err"`
}

func (r updateDayResponse) error() error { return r.Err }

func makeUpdateDayEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateDayRequest)
		day, err := s.UpdateDay(ctx, req.ListID, req.Date, req.Moments)

		return updateDayResponse{
			Days: day,
			Err:  err,
		}, nil
	}
}

//DeleteDay(ctx context.Context, listID string, date time.Time) error
type deleteDayRequest struct {
	ListID  string
	Date    time.Time
	Moments []Moment
}

type deleteDayResponse struct {
	Days Day   `json:"days"`
	Err  error `json:"err"`
}

func (r deleteDayResponse) error() error { return r.Err }

func makeDeleteDayEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteDayRequest)
		err := s.DeleteDay(ctx, req.ListID, req.Date)

		return deleteDayResponse{
			Err: err,
		}, nil
	}
}
