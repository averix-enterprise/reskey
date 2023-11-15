package resolution

import "time"

var waitTimestamp int64

func WaitRes(seconds int64) {
	waitTimestamp = time.Now().UnixMilli() + seconds*1000
}

func ChangeRes(width uint32, height uint32) {
	if time.Now().UnixMilli() < waitTimestamp {
		return
	}
	WaitRes(2)
	ChangeWindowsRes(width, height)
}
