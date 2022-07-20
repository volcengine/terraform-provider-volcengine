package iam_user_policy_attachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var iamUserPolicyAttachmentImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id is invalid")
	}
	if err := data.Set("user_name", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("policy_name", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("policy_type", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
