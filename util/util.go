package util

import "strconv"

// StringToID converts a string to a ID
func StringToID(s string) (userID int64) {
	userID, _ = strconv.ParseInt(s, 10, 64)
	return
}
