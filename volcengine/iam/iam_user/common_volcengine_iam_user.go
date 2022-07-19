package iam_user

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

var defaultConvert = func(data *schema.ResourceData, i interface{}) interface{} {
	return i
}

var phoneDiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
	return len(new) == 0
}
