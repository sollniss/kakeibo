package templates

import "strings"

var (
	japaneseDays = strings.NewReplacer("Mon", "月", "Tue", "火", "Wed", "水", "Thu", "木", "Fri", "金", "Sat", "土", "Sun", "日")
)
