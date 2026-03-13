package nlb_listener

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbListeners() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbListenersRead,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the NLB instance.",
			},
			"listener_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the listener.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The protocol of the listener. Valid values: `TCP`, `UDP`, `TLS`.",
			},
			"listener_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of listener IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": ve.TagsSchema(),
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"listeners": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of listener query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"listener_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the listener.",
						},
						"listener_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the listener.",
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the NLB instance.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the listener.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the listener. Valid values: `Creating`, `Active`, `Deleting`, `Disabled`.\n`Creating`: The listener is being created.\n`Active`: The listener is running.\n`Deleting`: The listener is being deleted.\n`Disabled`: The listener is disabled.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of the listener. Valid values: `TCP`, `UDP`, `TLS`.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The port used by the listener. 0 indicates that full port listening is enabled.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the server group associated with the listener.",
						},
						"connection_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The connection timeout of the listener.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the listener.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the listener.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the listener. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"start_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start port of the full port listening.",
						},
						"end_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end port of the full port listening.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of listener query.",
			},
		},
	}
}

func dataSourceVolcengineNlbListenersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbListenerService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbListeners())
}
