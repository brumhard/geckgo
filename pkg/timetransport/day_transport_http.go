package timetransport

import (
	"context"
	"encoding/json"
	"github.com/brumhard/geckgo/pkg/timeendpoint"
	"github.com/brumhard/geckgo/pkg/timeservice"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
)

const dateLayout = "2006-02-01"

var ErrNoBody = errors.New("no body")

//AddDay(ctx context.Context, listID int, date time.Time, moments []Moment) (Day, error)
func decodeAddDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idString, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	var body struct {
		Moments []timeservice.Moment `json:"moments"`
	}

	if r.Body == http.NoBody {
		return nil, ErrNoBody
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return timeendpoint.AddDayRequest{
		ListID:  id,
		Date:    date,
		Moments: body.Moments,
	}, nil
}

//GetDays(ctx context.Context, listID int, opts ...ListDaysOption) ([]Day, error)
func decodeGetDaysRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idString, ok := mux.Vars(r)["listId"]
	if !ok {
		return nil, errBadRoute
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	return timeendpoint.GetDaysRequest{
		ListID: id,
	}, nil
}

//GetDay(ctx context.Context, listID int, date time.Time) (Day, error)
func decodeGetDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idString, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	return timeendpoint.GetDayRequest{
		ListID: id,
		Date:   date,
	}, nil
}

//UpdateDay(ctx context.Context, listID int, date time.Time, moments []Moment) (Day, error)
func decodeUpdateDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	idString, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	var body struct {
		Moments []timeservice.Moment `json:"moments"`
	}

	if r.Body == http.NoBody {
		return nil, ErrNoBody
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return timeendpoint.UpdateDayRequest{
		ListID:  id,
		Date:    date,
		Moments: body.Moments,
	}, nil
}

//DeleteDay(ctx context.Context, listID int, date time.Time) error
func decodeDeleteDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idString, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	return timeendpoint.DeleteDayRequest{
		ListID: id,
		Date:   date,
	}, nil
}
