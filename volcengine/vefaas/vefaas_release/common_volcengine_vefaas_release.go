package vefaas_release

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

func vefaasReleaseImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'FunctionId:ReleaseRecordId'")
	}
	functionId := items[0]
	releaseRecordId := items[1]

	_ = d.Set("function_id", functionId)
	_ = d.Set("release_record_id", releaseRecordId)

	return []*schema.ResourceData{d}, nil
}
