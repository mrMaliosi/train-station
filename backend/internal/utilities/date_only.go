package utilities

import (
	"encoding/json"
	"errors"
	"time"
)

type DateOnly struct {
	time.Time
}

const layoutISO = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		return errors.New("empty date string")
	}
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(layoutISO))
}
