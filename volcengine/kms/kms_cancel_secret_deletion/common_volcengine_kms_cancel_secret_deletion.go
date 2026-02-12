package kms_cancel_secret_deletion

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func kmsCancelSecretDeletionImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	_ = d.Set("secret_name", d.Id())
	return []*schema.ResourceData{d}, nil
}
