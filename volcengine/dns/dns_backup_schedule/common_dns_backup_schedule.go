package dns_backup_schedule

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dnsBackupScheduleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	zid := d.Id()
	zidInt, err := strconv.Atoi(zid)
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf(" ZID cannot convert to int ")
	}

	_ = d.Set("zid", zidInt)

	return []*schema.ResourceData{d}, nil
}
