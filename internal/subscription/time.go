package subscription

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const monthYearLayout = "01-2006"

type MonthYear time.Time

func (m *MonthYear) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(monthYearLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date format (expected MM-YYYY): %w", err)
	}
	*m = MonthYear(t)
	return nil
}

func (m MonthYear) MarshalJSON() ([]byte, error) {
	t := time.Time(m)
	return []byte(fmt.Sprintf(`"%02d-%d"`, t.Month(), t.Year())), nil
}

func (m MonthYear) Value() (driver.Value, error) {
	return time.Time(m), nil
}

func (m *MonthYear) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*m = MonthYear(v)
		return nil
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*m = MonthYear(t)
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*m = MonthYear(t)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into MonthYear", value)
	}
}

func (m MonthYear) ToTime() time.Time {
	return time.Time(m)
}
