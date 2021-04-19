package timeendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

func validationMW(validator *validator.Validate) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err := validator.StructCtx(ctx, request); err != nil {
				return nil, ErrValidation(err)
			}

			return next(ctx, request)
		}
	}
}

type ValidationErr struct {
	Err error
}

func ErrValidation(err error) error {
	return &ValidationErr{Err: err}
}

func (e *ValidationErr) Error() string {
	return e.Err.Error()
}

func (e *ValidationErr) Unwrap() error {
	return e.Err
}
