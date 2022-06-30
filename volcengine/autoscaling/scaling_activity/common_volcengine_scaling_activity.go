package scaling_activity

import (
	"fmt"
	"time"
)

func timeValidation(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	if _, err := time.Parse("2006-01-02T15:04Z", v); err != nil {
		errors = append(errors, fmt.Errorf("expected %q to be a valid RFC3339 date, got %q: %+v", k, i, err))
	}

	return warnings, errors
}
