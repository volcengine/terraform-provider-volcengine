package host

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsHosts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsHostsRead,
		Schema: map[string]*schema.Schema{
			"host_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of host group.",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ip address.",
			},
			"heartbeat_status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The the heartbeat status.",
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
			"host_infos": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip address.",
						},
						"log_collector_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of log collector.",
						},
						"host_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of host group.",
						},
						"heartbeat_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The the heartbeat status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsHostsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsHosts())
}
