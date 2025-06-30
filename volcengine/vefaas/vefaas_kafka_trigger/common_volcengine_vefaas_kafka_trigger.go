package vefaas_kafka_trigger

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

func vefaasKafkaTriggerImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'FunctionId:ReleaseRecordId'")
	}
	functionId := items[0]
	kafkaTriggerId := items[1]

	_ = d.Set("function_id", functionId)
	_ = d.Set("id", kafkaTriggerId)

	return []*schema.ResourceData{d}, nil
}
