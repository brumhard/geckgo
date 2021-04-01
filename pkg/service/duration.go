package service

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

var ErrInvalidDuration = errors.New("invalid duration")

type Duration struct {
	time.Duration
}

func (d *Duration) Scan(src interface{}) error {
	mills, ok := src.(int64)
	if !ok {
		return ErrInvalidDuration
	}

	*d = Duration{time.Duration(mills) * time.Millisecond}

	return nil
}

func (d Duration) Value() (driver.Value, error) {
	return d.Milliseconds(), nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Milliseconds())
}

func (d *Duration) UnmarshalJSON(bytes []byte) error {
	var mills int64

	err := json.Unmarshal(bytes, &mills)
	if err != nil {
		return err
	}

	*d = Duration{time.Duration(mills) * time.Millisecond}

	return nil
}
