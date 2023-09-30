package transfer

import "time"

type ByMonthRequest struct {
	Month time.Month
	Year  int
}
