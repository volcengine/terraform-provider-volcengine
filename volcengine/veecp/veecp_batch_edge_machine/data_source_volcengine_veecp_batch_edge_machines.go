package veecp_batch_edge_machine

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVeecpBatchEdgeMachines() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVeecpBatchEdgeMachinesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Status of NodePool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.",
						},
						"edge_node_status_condition_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the status condition of the node pool in the active state. The value can be `Progressing` or `Ok` or `VersionPartlyUpgraded` or `StockOut` or `LimitedByQuota` or `Balance` or `Degraded` or `ClusterVersionUpgrading` or `Cluster` or `ResourceCleanupFailed` or `Unknown` or `ClusterNotRunning` or `SetByProvider`.",
						},
					},
				},
			},
			"create_client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ClientToken when successfully created.",
			},
			"ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IPs.",
			},
			"need_bootstrap_script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whether it is necessary to query the node management script.",
			},
			"cluster_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ClusterIds of NodePool IDs.",
			},
			"zone_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The Zone Ids.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of NodePool.",
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
			"machines": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of NodePool.",
						},
						"create_client_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ClientToken when successfully created.",
						},
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ClusterId of NodePool.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of NodePool.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CreateTime of NodePool.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UpdateTime time of NodePool.",
						},
						"status": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The Status of NodePool.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Condition state.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Indicates the status condition of the node pool in the active state.",
												},
											},
										},
									},
								},
							},
						},
						"bootstrap_script": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The bootstrap script.",
						},
						"profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Edge: Edge node pool. If the return value is empty, it is the central node pool.",
						},
						"edge_node_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Edge node type.",
						},
						"ttl_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The TTL time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVeecpBatchEdgeMachinesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVeecpBatchEdgeMachineService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVeecpBatchEdgeMachines())
}
