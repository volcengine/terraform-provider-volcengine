package allow_list

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBAllowLists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMongoDBAllowListsRead,
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
			"allow_list_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The allow list IDs to query.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allow_lists": {
				Description: "The collection of mongodb allow list query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_list_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of allow list.",
						},
						"allow_list_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of allow list.",
						},
						"allow_list_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The allow list name.",
						},
						"allow_list_ip_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of allow list IPs.",
						},
						"allow_list_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address type in allow list.",
						},
						"associated_instance_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of instances bound under the allow list.",
						},
						"allow_list": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The list of IP address in allow list.",
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
										Description: "The instance id that bound to the allow list.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The instance name that bound to the allow list.",
									},
									"vpc": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The VPC ID.",
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

func dataSourceVolcengineMongoDBAllowListsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBAllowListService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineMongoDBAllowLists())
}