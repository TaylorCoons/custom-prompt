package time

import "time"

func (t *Time) Initialize() {
	t.Time = time.Now().Format("15:04:05")
}
