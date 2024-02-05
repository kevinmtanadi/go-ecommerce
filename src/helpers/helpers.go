package helpers

import (
	"go-backend-template/src/constants"
	"time"
)

func FormatTimeString(t time.Time) string {
	return t.Format(constants.DATE_FORMAT)
}
