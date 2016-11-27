package util

import (
	"strconv"

	"github.com/Dunkelheit/feedbackapp/model"
)

// StringToID converts a string to a ID
func StringToID(s string) (userID model.ID) {
	parsedInt, _ := strconv.ParseInt(s, 10, 64)
	userID = model.ID(parsedInt)
	return
}
