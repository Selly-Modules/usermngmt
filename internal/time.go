package internal

import "time"

//
// NOTE: due to unique timezone in server's code, all using time will be convert to HCM timezone (UTC +7)
// All functions generate time, must be call util functions here
// WARNING: don't accept call time.Now() directly
//

// getHCMLocation ...
func getHCMLocation() *time.Location {
	l, _ := time.LoadLocation(timezoneHCM)
	return l
}

// Now ...
func Now() time.Time {
	return time.Now().In(getHCMLocation())
}

