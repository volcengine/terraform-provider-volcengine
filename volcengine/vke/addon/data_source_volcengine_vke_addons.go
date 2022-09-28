package addon

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVkeAddons() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVkeAddonsRead,
		Schema: map[string]*schema.Schema{
			"cluster_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IDs of Cluster.",
			},
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The Names of addons.",
			},
			"deploy_modes": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The deploy model, the value is `Managed` or `Unmanaged`.",
			},
			"deploy_node_types": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Array of addon states to filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The status of addon. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.",
						},
						"conditions_type": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The state condition in the current main state of the addon, that is, the reason for entering the main state, there can be multiple reasons, " +
								"the value contains `Progressing`, `Ok`, `Degraded`," +
								"`Unknown`, `ClusterNotRunning`, `CrashLoopBackOff`, `SchedulingFailed`, `NameConflict`, `ResourceCleanupFailed`, `ClusterVersionUpgrading`.",
						},
					},
				},
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ClientToken when the addon is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
			},
			"update_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken when the last addon update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of addon.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of addon query.",
			},
			"addons": {
				Description: "The collection of addon query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Cluster.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cluster.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cluster.",
						},
						"config": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The config of addon.",
						},
						"create_client_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ClientToken on successful creation. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
						},
						"update_client_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ClientToken when the last update was successful. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Addon creation time. UTC+0 time in standard RFC3339 format.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last time a request was accepted by the addon and executed or completed. UTC+0 time in standard RFC3339 format.",
						},
						"deploy_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deploy mode.",
						},
						"deploy_node_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deploy node type.",
						},
						"status": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The status of the addon.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of addon. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "The state condition in the current main state of the addon, that is, the reason for entering the main state, there can be multiple reasons, " +
														"the value contains `Progressing`, `Ok`, `Degraded`," +
														"`Unknown`, `ClusterNotRunning`, `CrashLoopBackOff`, `SchedulingFailed`, `NameConflict`, `ResourceCleanupFailed`, `ClusterVersionUpgrading`.",
												},
											},
										},
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

func dataSourceVolcengineVkeAddonsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVkeAddonService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVkeAddons())
}
