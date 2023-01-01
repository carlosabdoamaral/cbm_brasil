package common

import (
	"fmt"
	"time"
)

func TimeToString(t *time.Time) string {
	var h, m, s int

	if t != nil {
		h = t.Hour()
		m = t.Minute()
		s = t.Second()
	} else {
		now := time.Now()
		h = now.Hour()
		m = now.Minute()
		s = now.Second()
	}

	return fmt.Sprintf("%d:%d:%d", h, m, s)
}
