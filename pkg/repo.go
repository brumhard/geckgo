package pkg

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/go-kit/kit/log"
)

type Repository interface {
	AddList(ctx context.Context, list List) (int, error)
	GetLists(ctx context.Context) ([]List, error)
	GetListByID(ctx context.Context, listID int) (*List, error)
	UpdateList(ctx context.Context, list List) error
	DeleteListByID(ctx context.Context, listID int) error

	AddDay(ctx context.Context, listID int, day Day) error
	GetDays(ctx context.Context, listID int) ([]Day, error)
	GetDayByDate(ctx context.Context, listID int, date time.Time) (*Day, error)
	UpdateDay(ctx context.Context, listID int, day Day) error
	DeleteDayByDate(ctx context.Context, listID int, date time.Time) error
}

type repo struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewRepository(db *sqlx.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

// AddList inserts a list into the db an returns the id of the newly created list.
func (r *repo) AddList(ctx context.Context, list List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	err = tx.QueryRowContext(ctx, `INSERT INTO lists (name) VALUES ($1) RETURNING id`, list.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	if list.Settings != nil {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO list_settings (list_id, daily_time) VALUES ($1, $2)",
			id, list.Settings.DailyTime,
		)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetLists(ctx context.Context) ([]List, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT l.id, l.name, ls.daily_time FROM lists l LEFT JOIN list_settings ls on l.id = ls.list_id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []List

	for rows.Next() {
		list := List{}
		listSettings := &ListSettings{}

		err := rows.Scan(&list.ID, &list.Name, &listSettings.DailyTime)
		if err != nil {
			return nil, err
		}

		if listSettings.DailyTime != nil {
			list.Settings = listSettings
		}

		lists = append(lists, list)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return lists, err
}

func (r *repo) GetListByID(ctx context.Context, listID int) (*List, error) {
	list := &List{}
	listSettings := &ListSettings{}

	err := r.db.QueryRowContext(ctx,
		`SELECT l.id, l.name, ls.daily_time 
				FROM lists l 
				LEFT JOIN list_settings ls on l.id = ls.list_id
				WHERE l.id=$1`,
		listID,
	).Scan(&list.ID, &list.Name, &listSettings.DailyTime)
	if err != nil {
		return nil, err
	}

	if listSettings.DailyTime != nil {
		list.Settings = listSettings
	}

	return list, nil
}

func (r *repo) UpdateList(ctx context.Context, list List) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE lists SET name=$1 WHERE id=$2`, list.Name, list.ID)
	if err != nil {
		return err
	}

	if list.Settings != nil {
		_, err := tx.ExecContext(ctx,
			"INSERT INTO list_settings (list_id, daily_time) VALUES ($1, $2) ON CONFLICT (list_id) DO UPDATE SET daily_time=$2",
			list.ID, list.Settings.DailyTime,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *repo) DeleteListByID(ctx context.Context, listID int) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM lists WHERE id=$1`, listID)
	return err
}

func (r *repo) AddDay(ctx context.Context, listID int, day Day) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	insertStmt, err := tx.Prepare(`INSERT INTO moments (date, time, type, list_id) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}

	for _, moment := range day.Moments {
		if _, err := insertStmt.ExecContext(ctx, insertStmt, moment.Time, moment.Time, moment.Type, listID); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *repo) GetDays(ctx context.Context, listID int) ([]Day, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT date, time, type FROM moments WHERE list_id = $1", listID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dateToMoments := map[time.Time][]Moment{}

	for rows.Next() {
		var date time.Time
		var moment Moment

		err := rows.Scan(&date, &moment.Time, &moment.Type)
		if err != nil {
			return nil, err
		}

		moments, ok := dateToMoments[date]
		if ok {
			moments := []Moment{moment}
			dateToMoments[date] = moments
		}

		moments = append(moments, moment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	var days []Day
	for date, moments := range dateToMoments {
		days = append(days, Day{
			Date:    date,
			Moments: moments,
		})
	}

	return days, nil
}

func (r *repo) GetDayByDate(ctx context.Context, listID int, date time.Time) (*Day, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT date, time, type FROM moments WHERE date = $1 AND list_id = $2", date, listID,
	)
	if err != nil {
		return nil, err
	}

	var moments []Moment
	for rows.Next() {
		var moment Moment

		err := rows.Scan(&date, &moment.Time, &moment.Type)
		if err != nil {
			return nil, err
		}

		moments = append(moments, moment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return &Day{Date: date, Moments: moments}, nil
}

func (r *repo) UpdateDay(ctx context.Context, listID int, day Day) error {
	if err := r.DeleteDayByDate(ctx, listID, day.Date); err != nil {
		return err
	}

	return r.AddDay(ctx, listID, day)
}

func (r *repo) DeleteDayByDate(ctx context.Context, listID int, date time.Time) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM moments WHERE date = $1 AND list_id = $2", date, listID)
	return err
}
