package rocketmq_access_key

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRocketmqAccessKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRocketmqAccessKeysRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of rocketmq instance.",
			},
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The access key id of the rocketmq key.",
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

			"access_keys": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The access key id of the rocketmq key.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of rocketmq instance.",
						},
						"acl_config_json": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The acl config of the rocketmq key.",
						},
						"actived": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The active status of the rocketmq key.",
						},
						"all_authority": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The default authority of the rocketmq key.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rocketmq key.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the rocketmq key.",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The secret key of the rocketmq key.",
						},
						"topic_permissions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The custom authority of the rocketmq key.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the rocketmq topic.",
									},
									"permission": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The custom authority for the topic.",
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

func dataSourceVolcengineRocketmqAccessKeysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRocketmqAccessKeyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRocketmqAccessKeys())
}
