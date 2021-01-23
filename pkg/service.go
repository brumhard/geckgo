package pkg

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/log"
)

type Service interface {
	// lists
	AddList(ctx context.Context, name string, settings ListSettings) (List, error)
	GetLists(ctx context.Context) ([]List, error)
	GetList(ctx context.Context, listID string) (List, error)
	// TODO listsettings as optional?
	UpdateList(ctx context.Context, listID string, settings ListSettings) (List, error)
	DeleteList(ctx context.Context, listID string) error

	// days
	AddDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error)
	GetDays(ctx context.Context, listID string, opts ...ListDaysOption) ([]Day, error)
	GetDay(ctx context.Context, listID string, date time.Time) (Day, error)
	UpdateDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error)
	DeleteDay(ctx context.Context, listID string, date time.Time) error

	// dynamic utility functions
	StartDay(ctx context.Context, listID string, timeStamp time.Time) error
	StartBreak(ctx context.Context, listID string, timeStamp time.Time) error
	EndBreak(ctx context.Context, listID string, timeStamp time.Time) error
	EndDay(ctx context.Context, listID string, timeStamp time.Time) error
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repository Repository, logger log.Logger, db *sql.DB) Service {
	return &service{
		repo:   repository,
		logger: logger,
	}
}

type Day struct {
	Date    time.Time `json:"date"`
	Moments []Moment  `json:"moments"`
}

type Moment struct {
	Type MomentType `json:"type"`
	Time time.Time  `json:"time"`
}

type List struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Settings *ListSettings `json:"settings"`
}

type ListSettings struct {
	DailyTime *Duration `json:"omitempty,daily_time"`
}

type ListDaysOption func(*ListDayOptions)

type ListDayOptions struct {
}

func (s service) AddList(ctx context.Context, name string, settings ListSettings) (List, error) {
	panic("implement me")
}

func (s service) GetLists(ctx context.Context) ([]List, error) {
	panic("implement me")
}

func (s service) GetList(ctx context.Context, listID string) (List, error) {
	panic("implement me")
}

func (s service) UpdateList(ctx context.Context, listID string, settings ListSettings) (List, error) {
	panic("implement me")
}

func (s service) DeleteList(ctx context.Context, listID string) error {
	panic("implement me")
}

func (s service) AddDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error) {
	panic("implement me")
}

func (s service) GetDays(ctx context.Context, listID string, opts ...ListDaysOption) ([]Day, error) {
	panic("implement me")
}

func (s service) GetDay(ctx context.Context, listID string, date time.Time) (Day, error) {
	panic("implement me")
}

func (s service) UpdateDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error) {
	panic("implement me")
}

func (s service) DeleteDay(ctx context.Context, listID string, date time.Time) error {
	panic("implement me")
}

func (s service) StartDay(ctx context.Context, listID string, timeStamp time.Time) error {
	panic("implement me")
}

func (s service) StartBreak(ctx context.Context, listID string, timeStamp time.Time) error {
	panic("implement me")
}

func (s service) EndBreak(ctx context.Context, listID string, timeStamp time.Time) error {
	panic("implement me")
}

func (s service) EndDay(ctx context.Context, listID string, timeStamp time.Time) error {
	panic("implement me")
}
