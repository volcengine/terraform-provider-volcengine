package kms_cancel_key_deletion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func kmsCancelKeyDeletionImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var (
		keyId       string
		keyName     string
		keyringName string
	)

	parts := strings.Split(d.Id(), ":")
	switch len(parts) {
	case 1:
		keyId = parts[0]
	case 2:
		keyName = parts[0]
		keyringName = parts[1]
	default:
		return []*schema.ResourceData{d}, fmt.Errorf("format of kms cancel key deletion id is invalid,%s", d.Id())
	}

	if keyId != "" {
		_ = d.Set("key_id", keyId)
	} else {
		_ = d.Set("key_name", keyName)
		_ = d.Set("keyring_name", keyringName)
	}

	return []*schema.ResourceData{d}, nil
}
