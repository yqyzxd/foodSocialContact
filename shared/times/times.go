package times

import "time"

const (
	YYYY_MM_DD_HH_MM_SS_LAYOUT = "2006-01-02 15:04:05"
	YYYY_MM_DD_LAYOUT          = "2006-01-02"
	YYYY_MM_LAYOUT             = "2006-01"
	YYYYMM_LAYOUT              = "200601"
	YYYY_LAYOUT                = "2006"
	MM_LAYOUT                  = "01"
)

func Format(t *time.Time, layout string) string {
	if layout == "" {
		layout = YYYY_MM_DD_HH_MM_SS_LAYOUT
	}
	return t.Format(layout)
}

func Parse(timeStr string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)

}

// LengthOfMonth 某年某月有多少天
func LengthOfMonth(year int, month int) int {
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}

	_, ok := day31[month]
	if ok {
		return 31
	}
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	_, ok = day30[month]
	if ok {
		return 30
	}

	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}
