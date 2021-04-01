package pkg

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/go-kit/kit/log"
)

var (
	ErrAlreadyStarted = errors.New("already started")
	ErrNotStartedYet  = errors.New("not started yet")
	ErrNotEndedYet    = errors.New("not ended yet")
)

type Service interface {
	// lists
	AddList(ctx context.Context, name string, settings *ListSettings) (*List, error)
	GetLists(ctx context.Context) ([]List, error)
	GetList(ctx context.Context, listID int) (*List, error)
	UpdateList(ctx context.Context, listID int, name string, settings *ListSettings) (*List, error)
	DeleteList(ctx context.Context, listID int) error

	// days
	AddDay(ctx context.Context, listID int, date time.Time, moments []Moment) (*Day, error)
	GetDays(ctx context.Context, listID int, opts ...ListDaysOption) ([]Day, error)
	GetDay(ctx context.Context, listID int, date time.Time) (*Day, error)
	UpdateDay(ctx context.Context, listID int, date time.Time, moments []Moment) (*Day, error)
	DeleteDay(ctx context.Context, listID int, date time.Time) error

	// dynamic utility functions
	StartDay(ctx context.Context, listID int, timeStamp time.Time) error
	StartBreak(ctx context.Context, listID int, timeStamp time.Time) error
	EndBreak(ctx context.Context, listID int, timeStamp time.Time) error
	EndDay(ctx context.Context, listID int, timeStamp time.Time) error
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
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
	Settings *ListSettings `json:"settings,omitempty"`
}

type ListSettings struct {
	DailyTime *Duration `json:"daily_time,omitempty"`
}

type ListDaysOption func(*ListDayOptions)

type ListDayOptions struct {
}

func (s service) AddList(ctx context.Context, name string, settings *ListSettings) (*List, error) {
	newList := List{
		Name:     name,
		Settings: settings,
	}

	id, err := s.repo.AddList(ctx, newList)
	if err != nil {
		return nil, err
	}

	newList.ID = id

	return &newList, nil
}

func (s service) GetLists(ctx context.Context) ([]List, error) {
	return s.repo.GetLists(ctx)
}

func (s service) GetList(ctx context.Context, listID int) (*List, error) {
	return s.repo.GetListByID(ctx, listID)
}

func (s service) UpdateList(ctx context.Context, listID int, name string, settings *ListSettings) (*List, error) {
	toUpdate, err := s.repo.GetListByID(ctx, listID)
	if err != nil {
		return nil, err
	}

	if name != "" {
		toUpdate.Name = name
	}

	if toUpdate.Settings != nil {
		toUpdate.Settings = settings
	}

	err = s.repo.UpdateList(ctx, *toUpdate)
	if err != nil {
		return nil, err
	}

	return toUpdate, nil
}

func (s service) DeleteList(ctx context.Context, listID int) error {
	return s.repo.DeleteListByID(ctx, listID)
}

func (s service) AddDay(ctx context.Context, listID int, date time.Time, moments []Moment) (*Day, error) {
	newDay := Day{
		Date:    date,
		Moments: moments,
	}

	err := s.repo.AddDay(ctx, listID, newDay)
	if err != nil {
		return nil, err
	}

	return &newDay, err
}

func (s service) GetDays(ctx context.Context, listID int, opts ...ListDaysOption) ([]Day, error) {
	return s.repo.GetDays(ctx, listID)
}

func (s service) GetDay(ctx context.Context, listID int, date time.Time) (*Day, error) {
	if _, err := s.repo.GetListByID(ctx, listID); err != nil {
		return nil, err
	}

	return s.repo.GetDayByDate(ctx, listID, date)
}

func (s service) UpdateDay(ctx context.Context, listID int, date time.Time, moments []Moment) (*Day, error) {
	toUpdate, err := s.repo.GetDayByDate(ctx, listID, date)
	if err != nil {
		return nil, err
	}

	toUpdate.Moments = moments

	err = s.repo.UpdateDay(ctx, listID, *toUpdate)
	if err != nil {
		return nil, err
	}

	return toUpdate, nil
}

func (s service) DeleteDay(ctx context.Context, listID int, date time.Time) error {
	return s.repo.DeleteDayByDate(ctx, listID, date)
}

func (s service) StartDay(ctx context.Context, listID int, timeStamp time.Time) error {
	day, err := s.repo.GetDayByDate(ctx, listID, timeStamp)
	startMoment := Moment{Type: MomentTypeStart, Time: timeStamp}

	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}

		// check for ErrNotFound, if so AddDay with StartDay
		return s.repo.AddDay(ctx, listID, Day{
			Date:    timeStamp,
			Moments: []Moment{startMoment},
		})
	}

	// otherwise check if start is already there
	if startMoments := momentsOfType(day.Moments, MomentTypeStart); len(startMoments) > 0 {
		return errors.Wrap(ErrAlreadyStarted, "day")
	}

	day.Moments = append(day.Moments, startMoment)

	return s.repo.UpdateDay(ctx, listID, *day)
}

func (s service) StartBreak(ctx context.Context, listID int, timeStamp time.Time) error {
	day, err := s.repo.GetDayByDate(ctx, listID, timeStamp)
	if err != nil {
		return err
	}

	startBreakMoments := momentsOfType(day.Moments, MomentTypeStartBreak)
	endBreakMoments := momentsOfType(day.Moments, MomentTypeEndBreak)

	if len(startBreakMoments) > len(endBreakMoments) {
		return errors.Wrap(ErrAlreadyStarted, "break")
	}

	day.Moments = append(day.Moments, Moment{Type: MomentTypeStartBreak, Time: timeStamp})

	return s.repo.UpdateDay(ctx, listID, *day)
}

func (s service) EndBreak(ctx context.Context, listID int, timeStamp time.Time) error {
	day, err := s.repo.GetDayByDate(ctx, listID, timeStamp)
	if err != nil {
		return err
	}

	startBreakMoments := momentsOfType(day.Moments, MomentTypeStartBreak)
	endBreakMoments := momentsOfType(day.Moments, MomentTypeEndBreak)

	if len(startBreakMoments) == len(endBreakMoments) {
		return errors.Wrap(ErrNotStartedYet, "break")
	}

	day.Moments = append(day.Moments, Moment{Type: MomentTypeEndBreak, Time: timeStamp})

	return s.repo.UpdateDay(ctx, listID, *day)
}

func (s service) EndDay(ctx context.Context, listID int, timeStamp time.Time) error {
	day, err := s.repo.GetDayByDate(ctx, listID, timeStamp)
	if err != nil {
		return err
	}

	startMoments := momentsOfType(day.Moments, MomentTypeStart)
	if len(startMoments) == 0 {
		return errors.Wrap(ErrNotStartedYet, "day")
	}

	startBreakMoments := momentsOfType(day.Moments, MomentTypeStartBreak)
	endBreakMoments := momentsOfType(day.Moments, MomentTypeEndBreak)

	if len(startBreakMoments) > len(endBreakMoments) {
		return errors.Wrap(ErrNotEndedYet, "break")
	}

	day.Moments = append(day.Moments, Moment{Type: MomentTypeEndBreak, Time: timeStamp})

	return s.repo.UpdateDay(ctx, listID, *day)
}

func momentsOfType(moments []Moment, momentType MomentType) []Moment {
	var momentsOfType []Moment

	for _, moment := range moments {
		if moment.Type == momentType {
			momentsOfType = append(momentsOfType, moment)
		}
	}

	return momentsOfType
}
