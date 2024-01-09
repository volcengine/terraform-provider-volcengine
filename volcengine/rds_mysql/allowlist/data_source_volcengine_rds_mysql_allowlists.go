package allowlist

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlAllowLists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlAllowListsRead,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The region of the allow lists.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Scaling Activity query.",
			},
			"allow_lists": {
				Description: "The list of allowed list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_list_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the allow list.",
						},
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
						"allow_list_ip_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of IP addresses (or address ranges) in the whitelist.",
						},
						"allow_list_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the allow list.",
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
							Description: "The list of instances.",
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

func dataSourceVolcengineRdsMysqlAllowListsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlAllowLists())
}
