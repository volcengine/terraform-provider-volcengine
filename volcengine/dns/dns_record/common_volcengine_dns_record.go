package dns_record

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dnsRecordImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ZID:BackupID'")
	}
	zid := items[0]
	recordId := items[1]

	zidInt, err := strconv.Atoi(zid)
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf(" ZID cannot convert to int ")
	}

	_ = d.Set("zid", zidInt)
	_ = d.Set("record_id", recordId)

	return []*schema.ResourceData{d}, nil
}
