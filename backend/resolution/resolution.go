package resolution

import "time"

var waitTimestamp int64

func WaitRes(seconds int64) {
	waitTimestamp = time.Now().UnixMilli() + seconds*1000
}

func ChangeRes(width uint32, height uint32) bool {
	if time.Now().UnixMilli() < waitTimestamp {
		return false
	}
	ChangeWindowsRes(width, height)
	return true
}
