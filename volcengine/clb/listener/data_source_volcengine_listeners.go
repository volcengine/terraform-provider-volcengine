package listener

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineListeners() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineListenersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Listener IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Listener.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Listener query.",
			},
			"load_balancer_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the Clb.",
			},
			"listener_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Listener.",
			},
			"listeners": {
				Description: "The collection of Listener query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the Listener.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Listener.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Listener.",
						},
						"listener_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Listener.",
						},
						"listener_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Listener.",
						},
						"acl_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The acl status of the Listener.",
						},
						"acl_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The acl type of the Listener.",
						},
						"acl_ids": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Computed:    true,
							Description: "The acl ID list to which the Listener is bound.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of the Listener.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The port receiving request of the Listener.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Listener.",
						},
						"enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enable status of the Listener.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the backend server group which is associated with the Listener.",
						},
						"certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the certificate which is associated with the Listener.",
						},
						"health_check_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enable status of health check function.",
						},
						"health_check_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The interval executing health check.",
						},
						"health_check_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The response timeout of health check.",
						},
						"health_check_healthy_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The healthy threshold of health check.",
						},
						"health_check_un_healthy_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The unhealthy threshold of health check.",
						},
						"health_check_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method of health check.",
						},
						"health_check_uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uri of health check.",
						},
						"health_check_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain of health check.",
						},
						"health_check_http_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The normal http status code of health check.",
						},
						"health_check_udp_request": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A request string to perform a health check.",
						},
						"health_check_udp_expect": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected response string for the health check.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineListenersRead(d *schema.ResourceData, meta interface{}) error {
	listenerService := NewListenerService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(listenerService, d, DataSourceVolcengineListeners())
}
