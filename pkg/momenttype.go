package pkg

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

var ErrInvalidMomentType = errors.New("invalid moment type")

type MomentType int

const (
	MomentTypeStart MomentType = iota
	MomentTypeStartBreak
	MomentTypeEndBreak
	MomentTypeEnd
)

// TODO: add validation for MomentType enum
func (t MomentType) String() string {
	return []string{"start", "startBreak", "endBreak", "end"}[t]
}

func (t MomentType) StrErr() (string, error) {
	if t < MomentTypeStart || t > MomentTypeEnd {
		return "", ErrInvalidMomentType
	}
	return t.String(), nil
}

func (t *MomentType) ReadStr(str string) error {
	switch str {
	case "start":
		*t = MomentTypeStart
	case "startBreak":
		*t = MomentTypeStartBreak
	case "endBreak":
		*t = MomentTypeEndBreak
	case "end":
		*t = MomentTypeEnd
	default:
		return ErrInvalidMomentType
	}

	return nil
}

func (t *MomentType) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return ErrInvalidMomentType
	}

	return t.ReadStr(str)
}

func (t MomentType) Value() (driver.Value, error) {
	return t.StrErr()
}

func (t MomentType) MarshalJSON() ([]byte, error) {
	str, err := t.StrErr()
	return []byte(str), err
}

func (t *MomentType) UnmarshalJSON(bytes []byte) error {
	return t.ReadStr(string(bytes))
}
