package rds_ip_list

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsIpLists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsIpListsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of RDS ip list IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of RDS ip list.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of RDS ip list query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the RDS instance.",
			},
			"rds_ip_lists": {
				Description: "The collection of RDS ip list account query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the RDS ip list.",
						},
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS ip list.",
						},
						"ip_list": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The list of IP address.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsIpListsRead(d *schema.ResourceData, meta interface{}) error {
	rdsIpListService := NewRdsIpListService(meta.(*volc.SdkClient))
	return rdsIpListService.Dispatcher.Data(rdsIpListService, d, DataSourceVolcengineRdsIpLists())
}
