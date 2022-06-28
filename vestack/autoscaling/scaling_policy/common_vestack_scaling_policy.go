package scaling_policy

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var scalingPolicyImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("scaling_policy_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("scaling_group_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

var policyDiffSuppressFunc = func(policyTypes ...string) func(k, old, new string, d *schema.ResourceData) bool {
	tyMap := make(map[string]bool)
	for _, policyType := range policyTypes {
		tyMap[policyType] = true
	}
	return func(k, old, new string, d *schema.ResourceData) bool {
		return !tyMap[d.Get("scaling_policy_type").(string)]
	}
}

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
