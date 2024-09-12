package rds_postgresql_allowlist

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlAllowlists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlAllowlistsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the postgresql Instance.",
			},
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

			"postgresql_allow_lists": {
				Description: "The list of postgresql allowed list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the postgresql allow list.",
						},
						"allow_list_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the postgresql allow list.",
						},
						"allow_list_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the postgresql allow list.",
						},
						"allow_list_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the postgresql allow list.",
						},
						"allow_list_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the postgresql allow list.",
						},
						"allow_list_ip_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of IP addresses (or address ranges) in the whitelist.",
						},
						"associated_instance_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of instances bound under the whitelist.",
						},
						"allow_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The IP address or a range of IP addresses in CIDR format.",
						},
						"associated_instances": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of postgresql instances.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the postgresql instance.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the postgresql instance.",
									},
									"vpc": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the vpc.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlAllowlistsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlAllowlists())
}
