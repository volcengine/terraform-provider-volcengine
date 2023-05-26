package common

import (
	"strings"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

func ResourceNotFoundError(err error) bool {
	if e, ok := err.(volcengineerr.RequestFailure); ok && e.StatusCode() == 404 {
		return true
	}
	errMessage := strings.ToLower(err.Error())
	if strings.Contains(errMessage, "notfound") ||
		strings.Contains(errMessage, "not found") ||
		strings.Contains(errMessage, "not exist") ||
		strings.Contains(errMessage, "not associate") ||
		strings.Contains(errMessage, "invalid") ||
		strings.Contains(errMessage, "not_found") ||
		strings.Contains(errMessage, "notexist") {
		return true
	}
	return false
}
