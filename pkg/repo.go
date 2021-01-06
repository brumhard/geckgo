package pkg

import "github.com/go-kit/kit/log"

type Repository interface {
	// which functions to add?
}

type repo struct {
	// sql.DB
	logger log.Logger
}

func NewRepository(logger log.Logger) Repository {
	return &repo{
		// db: https://youtu.be/sjd2ePF3CuQ?t=589
		logger: logger,
	}
}
