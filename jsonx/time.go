package jsonx

import (
	"time"
)

type MilliTime struct {
	time.Time
}

func (mt MilliTime) MarshalJSON() ([]byte, error) {
	return json2.Marshal(mt.Milli())
}

func (mt *MilliTime) UnmarshalJSON(data []byte) error {
	var milli int64
	if err := json2.Unmarshal(data, &milli); err != nil {
		return err
	} else {
		mt.Time = time.Unix(0, milli*int64(time.Millisecond))
		return nil
	}
}

func (mt MilliTime) GetBSON() (interface{}, error) {
	return mt.Time, nil
}

func (mt MilliTime) Milli() int64 {
	return mt.UnixNano() / int64(time.Millisecond)
}
