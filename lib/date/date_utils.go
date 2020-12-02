package date

import (
	"time"
)

const LAYOUT_YYYYMMDD = "20060102"

func ToYYYYMMDD() string {
	t := time.Now()
	return t.Format(LAYOUT_YYYYMMDD)
}
