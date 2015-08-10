// Author: zheng-ji.info

package monitor

import (
	"time"
)

func Since(t time.Time) int {
	return int(time.Since(t).Nanoseconds() / 1e6)
}
