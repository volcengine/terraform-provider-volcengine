package dns_backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDnsBackups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDnsBackupsRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"zid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the domain for which you want to get the backup schedule.",
			},
			"backup_infos": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the backup.",
						},
						"backup_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the backup was created. The time zone is UTC + 8.",
						},
						"record_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of DNS records in the backup.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDnsBackupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDnsBackupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDnsBackups())
}
