package pkg

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/log"
)

type Repository interface {
	AddList(ctx context.Context, list List) (int, error)
	GetLists(ctx context.Context) ([]List, error)
	GetListByID(ctx context.Context, listID string) (*List, error)
	UpdateList(ctx context.Context, list List) error
	DeleteListByID(ctx context.Context, listID string) error

	AddDay(ctx context.Context, listID string, day Day) error
	GetDays(ctx context.Context, listID string) ([]Day, error)
	GetDayByDate(ctx context.Context, listID string, date time.Time) (*Day, error)
	UpdateDay(ctx context.Context, listID string, day Day) error
	DeleteDayByDate(ctx context.Context, listID string, date time.Time) error
}

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

// AddList inserts a list into the db an returns the id of the newly created list.
func (r *repo) AddList(ctx context.Context, list List) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, `INSERT INTO lists (name) VALUES ($1) RETURNING id`).Scan(&id)
	if err != nil {
		return 0, err
	}

	if list.Settings != nil {
		_, err := r.db.ExecContext(ctx,
			"INSERT INTO list_settings (list_id, daily_time) VALUES ($1, $2)",
			id, list.Settings.DailyTime,
		)
		if err != nil {
			return 0, err
		}
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
		list := List{Settings: &ListSettings{}}

		err := rows.Scan(&list.ID, &list.Name, &list.Settings.DailyTime)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return lists, err
}

func (r *repo) GetListByID(ctx context.Context, listID string) (*List, error) {
	list := &List{Settings: &ListSettings{}}
	err := r.db.QueryRowContext(ctx,
		`SELECT l.id, l.name, ls.daily_time 
				FROM lists l 
				LEFT JOIN list_settings ls on l.id = ls.list_id
				WHERE l.id=$1`,
		listID,
	).Scan(&list.ID, &list.Name, &list.Settings.DailyTime)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *repo) UpdateList(ctx context.Context, list List) error {
	_, err := r.db.ExecContext(ctx, `UPDATE lists SET name=$1 WHERE id=$2`, list.Name, list.ID)
	if err != nil {
		return err
	}

	if list.Settings != nil {
		_, err := r.db.ExecContext(ctx,
			"UPDATE list_settings SET daily_time=$1 WHERE list_id=$2",
			list.Settings.DailyTime, list.ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repo) DeleteListByID(ctx context.Context, listID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM lists WHERE id=$1`, listID)
	return err
}

func (r *repo) AddDay(ctx context.Context, listID string, day Day) error {
	panic("implement me")
}

func (r *repo) GetDays(ctx context.Context, listID string) ([]Day, error) {
	panic("implement me")
}

func (r *repo) GetDayByDate(ctx context.Context, listID string, date time.Time) (*Day, error) {
	panic("implement me")
}

func (r *repo) UpdateDay(ctx context.Context, listID string, day Day) error {
	panic("implement me")
}

func (r *repo) DeleteDayByDate(ctx context.Context, listID string, date time.Time) error {
	panic("implement me")
}
