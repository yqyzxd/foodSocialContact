package times

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {

	cases := []struct {
		layout string
		time   time.Time
		want   string
	}{
		{
			layout: YYYY_MM_DD_LAYOUT,
			time:   time.Now(),
			want:   "2022-07-19",
		},
		{
			layout: "20060102",
			time:   time.Now(),
			want:   "20220719",
		},
		{
			layout: YYYYMM_LAYOUT,
			time:   time.Now(),
			want:   "202207",
		},
		{
			layout: YYYY_MM_LAYOUT,
			time:   time.Now(),
			want:   "2022-07",
		},
		{
			layout: YYYY_LAYOUT,
			time:   time.Now(),
			want:   "2022",
		},
		{
			layout: MM_LAYOUT,
			time:   time.Now(),
			want:   "07",
		},
	}

	for _, cc := range cases {

		result := Format(&cc.time, cc.layout)
		if result != cc.want {
			t.Errorf("want %s got:%s", cc.want, result)
		}

		fmt.Println(result)

	}

}
