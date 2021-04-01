package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

func validationMW(validator *validator.Validate) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err := validator.StructCtx(ctx, request); err != nil {
				return nil, err
			}

			return next(ctx, request)
		}
	}
}
