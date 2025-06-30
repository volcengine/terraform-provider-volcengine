package kafka_allow_list

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKafkaAllowLists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKafkaAllowListsRead,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The region ID.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance ID to query.",
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
			"allow_lists": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_list_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the allow list.",
						},
						"allow_list_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the allow list.",
						},
						"allow_list_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the allow list.",
						},
						"allow_list_ip_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of rules specified in the whitelist.",
						},
						"associated_instance_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of instances bound to the whitelist.",
						},
						"allow_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Whitelist rule list.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"associated_instances": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of associated instances.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the instance.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the instance.",
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

func dataSourceVolcengineKafkaAllowListsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKafkaAllowListService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKafkaAllowLists())
}
