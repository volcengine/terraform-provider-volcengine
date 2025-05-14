package vefaas_timer

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

func vefaasTimerImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'FunctionId:ReleaseRecordId'")
	}
	functionId := items[0]
	timerId := items[1]

	_ = d.Set("function_id", functionId)
	_ = d.Set("id", timerId)

	return []*schema.ResourceData{d}, nil
}
