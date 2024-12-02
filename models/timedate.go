package models

import (
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
