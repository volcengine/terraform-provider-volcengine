package node

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVkeNodes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVkeNodesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Node IDs.",
			},
			"cluster_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Cluster IDs.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of Node.",
			},
			"node_pool_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The Node Pool IDs.",
			},
			"zone_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The Zone IDs.",
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Create Client Token.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Status of filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"Creating", "Running", "Updating", "Deleting", "Failed", "Starting", "Stopping", "Stopped"}, false),
							Description:  "The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Starting` or `Stopping` or `Stopped`.",
						},
						"conditions_type": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The Type of Node Condition, the value is `Progressing` or `Ok` or `Unschedulable` or `InitilizeFailed` or `Unknown`" +
								" or `NotReady` or `Security` or `Balance` or `ResourceCleanupFailed`.",
							ValidateFunc: validation.StringInSlice([]string{"Progressing", "Ok", "Unschedulable", "InitilizeFailed",
								"NotReady", "Security", "Balance", "ResourceCleanupFailed", "Unknown"}, false),
						},
					},
				},
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Node.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Node query.",
			},
			"nodes": {
				Description: "The collection of Node query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Node.",
						},
						"phase": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Phase of Node.",
						},
						"condition_types": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The Condition of Node.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of Node.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of Node.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of Node.",
						},
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cluster id of node.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance id of node.",
						},
						"node_pool_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node pool id.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id.",
						},
						"roles": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The roles of node.",
						},
						"create_client_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create client token of node.",
						},
						"is_virtual": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is virtual node.",
						},
						"additional_container_storage_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is Additional Container storage enables.",
						},
						"container_storage_path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Storage Path.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVkeNodesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVkeNodes())
}
