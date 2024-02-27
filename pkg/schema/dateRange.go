package schema

import (
	"emperror.dev/errors"
	"encoding/json"
	"strings"
	"time"
)

type DateFromTo struct {
	time.Time
	Granularity TimeGranularityType
}
type DateFrom DateFromTo
type DateTo DateFromTo

type TimeGranularityType int

const (
	SecondGranularity TimeGranularityType = iota
	MinuteGranularity
	HourGranularity
	DayGranularity
	MonthGranularity
	YearGranularity
)

var timeFormats = []string{
	time.RFC3339,
	"2006-01-02 15:04:05",
}
var dayFormats = []string{
	"2006-01-02",
	"20060102",
}
var monthFormats = []string{
	"2006-01",
	"200601",
}
var yearFormats = []string{
	"2006",
}

func UnmarshalTime(data []byte) (time.Time, TimeGranularityType, error) {
	var dataStr string
	dataStr = strings.Trim(string(data), `"`)
	if err := json.Unmarshal(data, &dataStr); err != nil {
	}
	for _, format := range timeFormats {
		if t, err := time.Parse(format, dataStr); err == nil {
			return t, SecondGranularity, nil
		}
	}
	for _, format := range dayFormats {
		if t, err := time.Parse(format, dataStr); err == nil {
			return t, DayGranularity, nil
		}
	}
	for _, format := range monthFormats {
		if t, err := time.Parse(format, dataStr); err == nil {
			return t, MonthGranularity, nil
		}
	}
	for _, format := range yearFormats {
		if t, err := time.Parse(format, dataStr); err == nil {
			return t, YearGranularity, nil
		}
	}
	return time.Time{}, 0, errors.Errorf("cannot parse date %s", string(data))
}
func (d *DateFrom) UnmarshalJSON(data []byte) error {
	t, granularity, err := UnmarshalTime(data)
	if err != nil {
		return errors.WithStack(err)
	}
	(*d).Time = t
	(*d).Granularity = granularity
	return nil
}

func (d *DateTo) UnmarshalJSON(data []byte) error {
	t, granularity, err := UnmarshalTime(data)
	if err != nil {
		return errors.WithStack(err)
	}
	switch granularity {
	case YearGranularity:
		t = t.AddDate(1, 0, -1)
		t = t.Add(time.Hour*24 - time.Second)
	case MonthGranularity:
		t = t.AddDate(0, 1, -1)
		t = t.Add(time.Hour*24 - time.Second)
	case DayGranularity:
		t = t.Add(time.Hour*24 - time.Second)
	case HourGranularity:
		t = t.Add(time.Hour - time.Second)
	case MinuteGranularity:
		t = t.Add(time.Minute - time.Second)
	case SecondGranularity:
		// nothing to do
	}
	(*d).Time = t
	(*d).Granularity = granularity
	return nil
}

type DateRange struct {
	From DateFrom `json:"gte,omitempty"`
	To   DateTo   `json:"lte,omitempty"`
}
