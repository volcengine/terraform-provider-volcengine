package common

import (
	"strings"

	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
)

func ResourceNotFoundError(err error) bool {
	if e, ok := err.(volcstackerr.RequestFailure); ok && e.StatusCode() == 404 {
		return true
	}
	errMessage := strings.ToLower(err.Error())
	if strings.Contains(errMessage, "notfound") ||
		strings.Contains(errMessage, "not found") ||
		strings.Contains(errMessage, "not exist") ||
		strings.Contains(errMessage, "not associate") ||
		strings.Contains(errMessage, "invalid") ||
		strings.Contains(errMessage, "not_found") {
		return true
	}
	return false
}
