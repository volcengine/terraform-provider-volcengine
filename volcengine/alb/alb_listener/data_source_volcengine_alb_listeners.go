package alb_listener

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineListeners() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbListenersRead,
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
				Description: "The id of the Alb.",
			},
			"listener_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Listener.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the listener.",
			},
			"listeners": {
				Description: "The collection of Listener query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Listener.",
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
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the backend server group which is associated with the Listener.",
						},
						"server_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of server groups with associated listeners.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of server group.",
									},
									"server_group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of server group.",
									},
								},
							},
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The load balancer ID that the listener belongs to.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of listener.",
						},
						"certificate_source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the certificate.",
						},
						"cert_center_certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The certificate id associated with the listener. Source is `cert_center`.",
						},
						"certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The certificate ID associated with the HTTPS listener.",
						},
						"ca_certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CA certificate ID associated with HTTPS listener.",
						},
						"enable_http2": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP2 feature switch,valid value is on or off.",
						},
						"enable_quic": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The QUIC feature switch,valid value is on or off.",
						},
						"acl_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to enable the access control function,valid value is on or off.",
						},
						"acl_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The access control type.",
						},
						"acl_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ID of the access control policy group bound to the listener, only returned when the AclStatus parameter is on.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"customized_cfg_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The customized configuration ID, the value is empty string when not bound.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the listener.",
						},
						"domain_extensions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The HTTPS listener association list of extension domains for.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_extension_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The extension domain ID.",
									},
									"certificate_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The server certificate ID that domain used.",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain.",
									},
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The listener ID that domain belongs to.",
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

func dataSourceVolcengineAlbListenersRead(d *schema.ResourceData, meta interface{}) error {
	listenerService := NewAlbListenerService(meta.(*ve.SdkClient))
	return listenerService.Dispatcher.Data(listenerService, d, DataSourceVolcengineListeners())
}
