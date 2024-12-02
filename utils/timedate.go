package utils

import (
	"errors"
	"strings"
	"time"
)

type CustomDate time.Time

const CustomDateFormat = "2006-01-02"

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	parsedTime, err := time.Parse(CustomDateFormat, str)
	if err != nil {
		return err
	}
	*cd = CustomDate(parsedTime)
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(cd).Format(CustomDateFormat) + `"`), nil
}

func (cd CustomDate) ToTime() time.Time {
	return time.Time(cd)
}

func TimeDateParse(dateStr string) (time.Time, error) {
	parsedTime, err := time.Parse(CustomDateFormat, dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	return parsedTime, nil
}
